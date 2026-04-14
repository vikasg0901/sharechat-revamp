# LLD: GET /v1/health Endpoint — sharechat-revamp Go Backend

## Overview

Introduce a minimal Go HTTP backend to the sharechat-revamp repository, exposing a single liveness endpoint:

```
GET /v1/health → 200 OK  {"status": "ok", "uptime_seconds": N}
```

The endpoint is **stateless, unauthenticated, and publicly reachable** — its sole purpose is process liveness signalling (load-balancer health checks, uptime monitoring). No database access. No auth middleware.

---

## Package Structure

### Before

```
sharechat-revamp/
├── .git/
├── CLAUDE.md
├── CONTRIBUTING.md
├── LLD.md
├── LICENSE
├── feed.html
└── index.html
```

### After

```
sharechat-revamp/
├── .git/
├── CLAUDE.md
├── CONTRIBUTING.md
├── LLD.md
├── LICENSE
├── feed.html
├── index.html
└── backend/                  ← NEW Go module root
    ├── go.mod                ← module sharechat-revamp, go 1.22
    ├── go.sum                ← generated; chi v5 dependency
    ├── main.go               ← router wiring, server startup, startTime capture
    ├── health.go             ← HealthResponse struct, NewHealthHandler factory
    └── health_test.go        ← httptest unit tests for the health handler
```

All three `.go` files belong to `package main`. The module root is `backend/`; no sub-packages are introduced for a single handler.

> **Why `backend/` subdirectory and not repo root?**  
> The repo root contains HTML files; mixing Go module artefacts (`go.mod`, `go.sum`, compiled binaries) at the root would pollute the frontend tree and break `go build ./...` glob semantics. A `backend/` subdirectory gives a clean boundary without over-engineering into `cmd/server/` + `internal/health/` for a one-endpoint service.

---

## Data Models

### `HealthResponse` (Go struct)

Defined in `backend/health.go`.

```go
// HealthResponse is the JSON payload returned by GET /v1/health.
type HealthResponse struct {
    Status        string `json:"status"`          // always "ok"
    UptimeSeconds int64  `json:"uptime_seconds"`  // seconds since process start
}
```

| Field           | Go type | JSON key         | Constraints                          |
|-----------------|---------|------------------|--------------------------------------|
| `Status`        | `string`| `status`         | Always the literal string `"ok"`     |
| `UptimeSeconds` | `int64` | `uptime_seconds` | Non-negative; monotonically increasing |

No optional fields. No nullable fields. The JSON encoding is deterministic.

---

## DB Schema

**N/A.** The health endpoint does not read from or write to any database. It derives `uptime_seconds` entirely from an in-process `time.Time` value captured at startup — no persistence layer is involved.

---

## API Contract

### Request

```
GET /v1/health HTTP/1.1
Host: <server>
```

No query parameters. No request body. No authentication header required.

### Response — success (only possible response)

```
HTTP/1.1 200 OK
Content-Type: application/json

{"status":"ok","uptime_seconds":42}
```

| Status code | Condition                     |
|-------------|-------------------------------|
| `200 OK`    | Process is running (always)   |

The endpoint **never returns 4xx or 5xx** under normal operation. If the process cannot serve this endpoint, the connection itself will fail at the TCP/TLS layer — there is no application-level error path to design.

### Content-Type

`application/json` is set explicitly before writing the response body. The encoder writes a single-line JSON object followed by a newline (`\n`) — standard `json.NewEncoder` behaviour.

### Caching

No `Cache-Control` header is set. Load balancers and health-check clients must not cache this response; this is the expected default for `200 OK` responses with no explicit cache directive on dynamic endpoints.

---

## Sequence Diagram

### Normal request flow

```
Client (LB / curl)          chi Router           HealthHandler
        |                        |                      |
        |-- GET /v1/health ----->|                      |
        |                        |-- route matched      |
        |                        |-- call handler ----->|
        |                        |                      |-- time.Since(startTime)
        |                        |                      |-- marshal HealthResponse
        |                        |                      |-- w.WriteHeader(200)
        |                        |                      |-- json.Encoder.Encode(resp)
        |                        |<----- return --------|
        |<-- 200 {"status":"ok","uptime_seconds":N} -----|
```

### Server startup sequence

```
OS / container runtime        main()                  chi.Router
        |                        |                        |
        |-- exec backend ------->|                        |
        |                        |-- startTime = time.Now()
        |                        |-- r := chi.NewRouter() -->|
        |                        |                        |-- middleware.Logger
        |                        |                        |-- middleware.Recoverer
        |                        |-- r.Get("/v1/health", NewHealthHandler(startTime))
        |                        |-- http.ListenAndServe(":8080", r)
        |                        |   (blocks; accepts connections)
```

---

## Implementation Patterns

### `backend/go.mod`

```
module sharechat-revamp

go 1.22

require github.com/go-chi/chi/v5 v5.1.0
```

chi v5 is the only runtime dependency. No ORM, no logger library, no config library — the single-endpoint surface does not justify additional dependencies.

### `backend/health.go`

```go
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse is the JSON payload returned by GET /v1/health.
type HealthResponse struct {
	Status        string `json:"status"`
	UptimeSeconds int64  `json:"uptime_seconds"`
}

// NewHealthHandler returns an http.HandlerFunc that reports process liveness.
// startTime is captured at process start and injected here so tests can
// control the clock without patching globals.
func NewHealthHandler(startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := HealthResponse{
			Status:        "ok",
			UptimeSeconds: int64(time.Since(startTime).Seconds()),
		}
		json.NewEncoder(w).Encode(resp) //nolint:errcheck
	}
}
```

**Key design decisions:**

- `NewHealthHandler` is a **factory function** (not a method on a struct). This is the minimal form that allows `startTime` injection without a struct. If the handler later needs e.g. a DB pool or metrics client, migrate to a struct receiver — the call site in `main.go` does not change.
- `startTime` is **injected via closure**, not read from a package-level variable. This makes the handler directly unit-testable without any setup/teardown.
- `json.NewEncoder(w).Encode(resp)` is used instead of `json.Marshal` + `w.Write` because `Encoder` streams directly to the `ResponseWriter` and appends the required trailing newline without a separate allocation.
- `//nolint:errcheck` on the encode call: `w.Write` errors after a `WriteHeader(200)` are unrecoverable (the status code is already sent to the client); logging them would be noise. This is the standard Go pattern.

### `backend/main.go`

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	startTime := time.Now()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/v1/health", NewHealthHandler(startTime))

	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
```

**Key design decisions:**

- `startTime := time.Now()` is the **first statement in `main()`**, before any I/O or flag parsing. This maximises the accuracy of the reported uptime.
- The health route is registered **outside any authentication middleware group**. `middleware.Logger` and `middleware.Recoverer` are global middleware (request logging and panic recovery) that are acceptable on a public endpoint. No `jwtauth.Verifier` or equivalent is applied.
- Port `8080` is hardcoded. If environment-based config is added later, replace with `os.Getenv("PORT")` with a `"8080"` default — that is a separate concern outside this LLD's scope.

### `backend/health_test.go`

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler_StatusOK(t *testing.T) {
	fakeStart := time.Now().Add(-5 * time.Second)
	handler := NewHealthHandler(fakeStart)

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}
}

func TestHealthHandler_ResponseBody(t *testing.T) {
	fakeStart := time.Now().Add(-10 * time.Second)
	handler := NewHealthHandler(fakeStart)

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	var body HealthResponse
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("expected status %q, got %q", "ok", body.Status)
	}
	if body.UptimeSeconds < 10 {
		t.Errorf("expected uptime_seconds >= 10, got %d", body.UptimeSeconds)
	}
}

func TestHealthHandler_UptimeGrowsOverTime(t *testing.T) {
	fakeStart := time.Now().Add(-100 * time.Second)
	handler := NewHealthHandler(fakeStart)

	// First call
	req1 := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)
	var body1 HealthResponse
	json.NewDecoder(w1.Body).Decode(&body1) //nolint:errcheck

	// Second call — wall clock has advanced; uptime_seconds must be >= first
	req2 := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)
	var body2 HealthResponse
	json.NewDecoder(w2.Body).Decode(&body2) //nolint:errcheck

	if body2.UptimeSeconds < body1.UptimeSeconds {
		t.Errorf("uptime must be non-decreasing: first=%d second=%d",
			body1.UptimeSeconds, body2.UptimeSeconds)
	}
}
```

**Test strategy:**

| Test | What it asserts | Why separate |
|------|-----------------|--------------|
| `TestHealthHandler_StatusOK` | HTTP 200, `Content-Type: application/json` | HTTP contract — independent of body semantics |
| `TestHealthHandler_ResponseBody` | `status == "ok"`, `uptime_seconds >= injected offset` | Body contract — verifies struct fields, JSON keys, and uptime floor |
| `TestHealthHandler_UptimeGrowsOverTime` | Second call's uptime ≥ first call's uptime | Monotonicity — guards against accidentally returning a constant |

All three tests use **`httptest.NewRecorder`** (no network, no port binding). The `startTime` injection pattern means no global variable mutation, no `t.Cleanup` required, and tests are safe to run in parallel with `t.Parallel()` if added later.

---

## Validation Checklist

| # | Check | Command | Expected |
|---|-------|---------|----------|
| 1 | Module resolves | `go mod tidy` in `backend/` | `go.sum` updated, no missing deps |
| 2 | Build succeeds | `go build ./...` in `backend/` | Binary produced, exit 0 |
| 3 | Tests pass | `go test ./...` in `backend/` | `PASS`, 3 tests |
| 4 | Route reachable | `go run . &; curl -s localhost:8080/v1/health` | `{"status":"ok","uptime_seconds":0}` |
| 5 | Unknown route 404 | `curl -s -o /dev/null -w "%{http_code}" localhost:8080/v1/unknown` | `404` |
| 6 | Content-Type header | `curl -sI localhost:8080/v1/health \| grep Content-Type` | `Content-Type: application/json` |

---

## Risks & Mitigations

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| `uptime_seconds` wraps or overflows | Very low (int64 max ~292 years) | Incorrect metric | `int64` is sufficient; no mitigation needed |
| Clock skew on container restart resets uptime | Medium | Monitoring alert noise | Documented behaviour: uptime resets on process restart by design; this is a liveness check, not a wall-clock uptime |
| Auth middleware accidentally applied to `/v1/health` | Low | Load balancer marks backend unhealthy | Health route is registered before any auth group in `main.go`; CI test confirms 200 with no auth header |
| chi not available at build time (no `go mod tidy`) | Low | Build failure | `go mod tidy` is step 1 in validation checklist |
| Port 8080 conflicts in CI | Low | Test startup failure | Tests use `httptest.NewRecorder` — they never bind a port; port conflict only affects manual validation |

---

## Alternatives Considered

| Option | Rationale for Rejection |
|--------|------------------------|
| **`net/http` without chi** | The task explicitly requires chi. Additionally, chi's `middleware.Recoverer` prevents panics in future handlers from crashing the process — worth the dependency |
| **Global `startTime` package var** | Testable only by mutating global state between tests; fragile under `t.Parallel()`. Factory-function injection is cleaner with identical runtime behaviour |
| **`status: "healthy"` instead of `"ok"`** | The spec is explicit: `{"status": "ok"}`. No deviation |
| **Separate `internal/health` package** | One handler, one struct — a sub-package adds import indirection with no encapsulation benefit at this scale |
| **Return `uptime_seconds` as `float64`** | Sub-second precision is meaningless for a liveness check; `int64` integer seconds matches the spec's `N` notation and avoids `1.234e+02` scientific notation in JSON |

---

## Open Questions

### Q1: TLS / HTTPS

**UNRESOLVED — OUT OF SCOPE.** `http.ListenAndServe` serves plaintext HTTP. TLS termination (nginx, load balancer, or `http.ListenAndServeTLS`) is an infra concern outside this LLD. The health endpoint contract is identical over HTTPS.

### Q2: Graceful shutdown

**UNRESOLVED — OUT OF SCOPE.** `http.ListenAndServe` blocks until the process exits. A production deployment would use `http.Server` with `Shutdown(ctx)` on SIGTERM. Adding graceful shutdown is a follow-on task; it does not affect the health endpoint contract.

### Q3: Metrics endpoint (`/v1/metrics`)

**RESOLVED — OUT OF SCOPE.** Prometheus metrics exposure is a separate concern. This LLD covers only the liveness check.
