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
