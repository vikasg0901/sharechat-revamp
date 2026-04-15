package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestV10HealthHandler_StatusOK verifies the handler returns HTTP 200.
func TestV10HealthHandler_StatusOK(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status: got %d want %d", rr.Code, http.StatusOK)
	}
}

// TestV10HealthHandler_ContentType verifies Content-Type is application/json.
func TestV10HealthHandler_ContentType(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)
	ct := rr.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Fatalf("content-type: got %q want application/json", ct)
	}
}

// TestV10HealthHandler_ResponseBody verifies the JSON payload matches the spec
// {"status":"ok","version":"1.0"} with both fields as JSON strings.
func TestV10HealthHandler_ResponseBody(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)

	var body V10HealthResponse
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Status != "ok" {
		t.Errorf("status: got %q want \"ok\"", body.Status)
	}
	if body.Version != "1.0" {
		t.Errorf("version: got %q want \"1.0\"", body.Version)
	}
}

// TestV10HealthHandler_SchemaExactKeys verifies the JSON body has exactly two
// keys — "status" and "version" — with no extra or missing fields.
func TestV10HealthHandler_SchemaExactKeys(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)

	var raw map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&raw); err != nil {
		t.Fatalf("decode: %v", err)
	}

	want := map[string]bool{"status": true, "version": true}
	for k := range raw {
		if !want[k] {
			t.Errorf("unexpected key %q in response body", k)
		}
		delete(want, k)
	}
	for k := range want {
		t.Errorf("missing required key %q in response body", k)
	}
}

// TestV10HealthHandler_VersionIsString verifies the "version" field is encoded
// as a JSON string, not a floating-point number. json.Marshal would encode a
// Go float64 1.0 as the number 1, not the string "1.0".
func TestV10HealthHandler_VersionIsString(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)

	// Decode into map[string]any so we get the raw JSON type.
	var raw map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&raw); err != nil {
		t.Fatalf("decode: %v", err)
	}
	v, ok := raw["version"]
	if !ok {
		t.Fatal("version field missing from response body")
	}
	if _, isStr := v.(string); !isStr {
		t.Errorf("version must be a JSON string, got %T (%v)", v, v)
	}
}

// TestV10HealthHandler_BodyEndsWithNewline verifies that the response body ends
// with '\n', which is the encoding contract of json.NewEncoder (used by the
// handler). Switching to json.Marshal without appending a newline would break
// this contract and fail stream parsers.
func TestV10HealthHandler_BodyEndsWithNewline(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)

	body := rr.Body.Bytes()
	if len(body) == 0 {
		t.Fatal("response body is empty")
	}
	if body[len(body)-1] != '\n' {
		t.Errorf("body must end with '\\n' (json.NewEncoder contract), last byte: %q", body[len(body)-1])
	}
}

// TestV10HealthHandler_Idempotent verifies the handler returns the same
// payload across repeated calls (it is stateless and has no mutable fields).
func TestV10HealthHandler_Idempotent(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()

	call := func() V10HealthResponse {
		req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
		rr := httptest.NewRecorder()
		h(rr, req)
		var b V10HealthResponse
		_ = json.NewDecoder(rr.Body).Decode(&b)
		return b
	}

	first, second, third := call(), call(), call()
	for i, b := range []V10HealthResponse{first, second, third} {
		if b.Status != "ok" {
			t.Errorf("call %d: status got %q want \"ok\"", i+1, b.Status)
		}
		if b.Version != "1.0" {
			t.Errorf("call %d: version got %q want \"1.0\"", i+1, b.Version)
		}
	}
}

// TestV10HealthHandler_NoServerHeader verifies that the response does not
// expose a Server header. Go's net/http does not add one by default; its
// presence would indicate a newly-added reverse proxy or middleware regression.
func TestV10HealthHandler_NoServerHeader(t *testing.T) {
	t.Parallel()
	h := NewV10HealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v10/health", nil)
	rr := httptest.NewRecorder()
	h(rr, req)

	if s := rr.Header().Get("Server"); s != "" {
		t.Errorf("unexpected Server header: %q", s)
	}
}
