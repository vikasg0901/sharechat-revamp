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
