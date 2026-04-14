package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealth_NoDependencies verifies that the health endpoint returns 200 ok
// when no downstream dependencies are configured.
func TestHealth_NoDependencies_ReturnsOK(t *testing.T) {
	h := NewHealthHandler(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	w := httptest.NewRecorder()

	h.Health(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp healthResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got %q", resp.Status)
	}
	if resp.Time == "" {
		t.Error("expected non-empty time field")
	}
	if resp.Checks != nil {
		t.Errorf("expected nil checks when no dependencies configured, got %v", resp.Checks)
	}
}

// TestHealth_ContentType verifies the response has the correct Content-Type.
func TestHealth_ContentType_IsJSON(t *testing.T) {
	h := NewHealthHandler(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	w := httptest.NewRecorder()

	h.Health(w, req)

	ct := w.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}
}
