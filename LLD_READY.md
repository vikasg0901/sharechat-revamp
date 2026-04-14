# LLD: Add GET /api/v7/ready endpoint to sharechat-revamp

## 1. Overview & Scope

Add a `GET /api/v7/ready` HTTP endpoint to the existing Go/chi backend at `backend/`. The endpoint is a static readiness probe that returns `{"ready": true}` with no dynamic computation or external dependency checks.

**In scope:**
- New `backend/ready.go` — handler function + response type
- New `backend/ready_test.go` — two unit tests
- One-line addition to `backend/main.go` to register the route

**Out of scope:**
- Authentication / authorisation
- Database or cache dependency checks
- Metrics, tracing, or structured logging
- Deployment configuration (Dockerfile, CI)

---

### Spec Clarification: Express / server.js

The ticket describes adding the endpoint "to the existing Express routes in server.js." **The backend is Go (chi router), not Node/Express; there is no `server.js`.** The intent (add a readiness probe to the existing server) is clear and is implemented in Go below. If a Node/Express server is genuinely required, that is a separate, larger task (introduce a second runtime, define coexistence strategy) and is outside the scope of this LLD.

---

### Path Inconsistency Note

The existing health endpoint is at `/v1/health`. This new endpoint is specified at `/api/v7/ready` — a different prefix scheme and a jump to version 7 in a single-endpoint backend. The LLD implements the path exactly as specified, but the inconsistency is flagged here for the tech lead to resolve before the implementation is merged. Options: align to `/api/v1/ready`, or document why the two endpoints use different versioning conventions.

---

## 2. Package Structure

### Before

```
backend/
├── go.mod
├── go.sum
├── main.go          # wires router, registers /v1/health
├── health.go        # HealthResponse type + NewHealthHandler constructor
└── health_test.go   # three unit tests for health handler
```

### After

```
backend/
├── go.mod
├── go.sum
├── main.go          # +1 line: r.Get("/api/v7/ready", ReadyHandler)
├── health.go
├── health_test.go
├── ready.go         ← NEW: ReadyResponse type + ReadyHandler func
└── ready_test.go    ← NEW: two parallel unit tests
```

All files remain in `package main`. No new dependencies; `encoding/json` and `net/http` are already imported in the existing codebase.

---

## 3. API Contract

| Property       | Value                        |
|----------------|------------------------------|
| Method         | `GET`                        |
| Path           | `/api/v7/ready`              |
| Success status | `200 OK`                     |
| Content-Type   | `application/json`           |
| Response body  | `{"ready":true}`             |
| Error path     | None — handler is stateless  |
| Auth required  | No                           |

`ready` is a JSON boolean (`true`), not a string. The Go encoder maps `bool` to the JSON literal `true` with no quotes.

---

## 4. Auth-Scope Call-Out

The ready route **must be registered directly on the root router**, not inside any `r.Group(...)` that attaches authentication middleware. Registering a readiness probe behind auth causes orchestrator health checks (Kubernetes, ECS, load balancers) to receive 401, which will cycle the pod or drain the target group even when the process is fully healthy.

```go
// main.go — CORRECT placement
r.Get("/api/v7/ready", ReadyHandler)   // unauthenticated, direct on root router

// WRONG — do not do this
r.Group(func(r chi.Router) {
    r.Use(authMiddleware)
    r.Get("/api/v7/ready", ReadyHandler)  // orchestrators would get 401
})
```

---

## 5. Data Models

`ReadyResponse` is the only type introduced. It has one field.

```go
// ready.go
type ReadyResponse struct {
    Ready bool `json:"ready"`
}
```

The `Ready` field is always set to `true`. It is modelled as a field (not a literal in `json.Marshal`) so that:
- The JSON key `"ready"` is part of the struct definition and visible to a schema-generation tool.
- Tests can decode into the struct and assert the field value by name rather than doing string-equality on raw JSON.

---

## 6. DB Schema

**N/A.** The endpoint reads no database and writes no database. It returns a constant value derived solely from the fact that the process started. Adding a database ping to a readiness probe is an anti-pattern: a transient DB blip would take the process out of rotation even though it can still serve cached or stateless requests. Dependency checks belong in a separate, explicitly-named endpoint (e.g., `/api/v7/healthz`) with a documented failure policy.

---

## 7. Sequence Diagram

```
Client (LB / orchestrator)          Go/chi server
          |                               |
          |-- GET /api/v7/ready --------->|
          |                         ReadyHandler called
          |                         encode ReadyResponse{Ready: true}
          |<-- 200 OK                     |
          |    Content-Type: application/json
          |    {"ready":true}             |
```

No downstream calls. End-to-end latency is bounded by network RTT + JSON encode (~microseconds).

---

## 8. Implementation Pattern

### `backend/ready.go`

```go
package main

import (
	"encoding/json"
	"net/http"
)

type ReadyResponse struct {
	Ready bool `json:"ready"`
}

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ReadyResponse{Ready: true})
}
```

`ReadyHandler` is a plain `http.HandlerFunc`-compatible function, not a constructor returning a closure (contrast with `NewHealthHandler`). A constructor is unnecessary here because there is no injected dependency (no `startTime`, no config). Using a plain function keeps the signature minimal.

### Route registration (`backend/main.go`, one-line addition)

```go
// Ready probe — registered outside auth group so orchestrators
// never receive 401. Do not move inside r.Group calls that attach
// auth middleware.
r.Get("/api/v7/ready", ReadyHandler)
```

Place this line immediately after the existing `/v1/health` registration so all unauthenticated probes are co-located and the invariant comment applies to both.

---

## 9. Test Plan

Both tests call `t.Parallel()`. No real server is started; `httptest.NewRecorder` is used throughout.

| Test name | Assertion |
|---|---|
| `TestReadyHandler_StatusOK` | HTTP status is 200; `Content-Type` header is `application/json` |
| `TestReadyHandler_ResponseBody` | JSON body decodes without error; `ready == true` |

### `backend/ready_test.go`

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadyHandler_StatusOK(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/api/v7/ready", nil)
	rr := httptest.NewRecorder()
	ReadyHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status: got %d want %d", rr.Code, http.StatusOK)
	}
	if got := rr.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("content-type: got %q want application/json", got)
	}
}

func TestReadyHandler_ResponseBody(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/api/v7/ready", nil)
	rr := httptest.NewRecorder()
	ReadyHandler(rr, req)
	var body ReadyResponse
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !body.Ready {
		t.Fatalf("ready field: got %v want true", body.Ready)
	}
}
```

Two tests is the minimum justified set for a stateless constant-response handler:

- **`TestReadyHandler_StatusOK`** — verifies the HTTP contract (status + Content-Type). These are observable at the transport level and cannot be folded into the body test without reading the recorder twice.
- **`TestReadyHandler_ResponseBody`** — verifies the JSON contract. Without this test, a handler that returns `{}` or `{"ready":false}` would pass the first test.

A third "value never changes" test (analogous to `TestHealthHandler_UptimeGrowsOverTime`) is not justified here: the handler has no time-varying state, so there is nothing to assert across two calls.

---

## 10. Reflection Block

### Tradeoffs

**(a) Plain function vs. closure constructor**

`NewHealthHandler` wraps a `startTime` parameter via closure because the handler has an injected dependency. `ReadyHandler` has no dependency and no closure is needed. Using a constructor here would add indirection for no benefit and diverge from the Go standard library pattern (`http.HandlerFunc`). Plain function is correct.

**(b) `/api/v7/ready` path vs. `/v1/ready`**

The path is implemented as specified. The mismatch between `/v1/health` (existing) and `/api/v7/ready` (new) is noted in §1 and left for the tech lead to resolve. The LLD does not unilaterally normalize the path, because changing a specified API contract is a product decision.

---

### Failure Modes

**(a) Route moved inside auth group during future refactor**

If a developer adds a blanket `r.Use(authMiddleware)` above the ready route in `main.go`, orchestrators begin receiving 401. Mitigation: the comment in `main.go` documents the invariant; a smoke test or integration test should issue an unauthenticated `GET /api/v7/ready` and assert 200.

**(b) `ready` field changed from `bool` to `string` (e.g., `"true"`)**

JSON encoding produces `"true"` (quoted string) instead of `true` (boolean). Clients that check `response.ready === true` (strict equality) would fail. Mitigation: `TestReadyHandler_ResponseBody` decodes into `ReadyResponse{Ready bool}` — a type change to `string` would break compilation and surface the contract drift before merge.

---

### Six-Month Concerns

1. **Path versioning alignment.** If the API grows and `/api/v7/` becomes the canonical prefix, the `/v1/health` endpoint will be misaligned. A router-level prefix group (`r.Route("/api/v7", ...)`) should be introduced when a second `/api/v7/` route is added.

2. **Readiness semantics may expand.** A common evolution: "ready" should mean "DB connection is live," not just "process is up." That requires a dependency check and changes the error contract (503 when DB is down). The current handler is a liveness probe in readiness clothing. If semantics expand, the response schema and test suite must be updated; the DB section (§6) records the rationale for keeping dependency checks out of this endpoint.
