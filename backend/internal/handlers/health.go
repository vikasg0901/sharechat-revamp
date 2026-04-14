package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// HealthHandler holds optional downstream dependency handles for health checks.
type HealthHandler struct {
	db  *pgxpool.Pool // nil if database not configured
	rdb *redis.Client // nil if Redis not configured
}

// NewHealthHandler creates a HealthHandler with optional DB and Redis clients.
// Pass nil for any dependency that is not used by the service.
func NewHealthHandler(db *pgxpool.Pool, rdb *redis.Client) *HealthHandler {
	return &HealthHandler{db: db, rdb: rdb}
}

type healthResponse struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks,omitempty"`
	Time   string            `json:"time"`
}

// Health handles GET /v1/health.
// Returns 200 OK when all configured dependencies are reachable, or
// 503 Service Unavailable when any dependency check fails.
// No authentication required.
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	resp := healthResponse{
		Status: "ok",
		Time:   time.Now().UTC().Format(time.RFC3339),
	}

	checks := make(map[string]string)
	degraded := false

	if h.db != nil {
		if err := h.db.QueryRow(ctx, "SELECT 1").Scan(new(int)); err != nil {
			checks["db"] = "error"
			degraded = true
		} else {
			checks["db"] = "ok"
		}
	}

	if h.rdb != nil {
		if err := h.rdb.Ping(ctx).Err(); err != nil {
			checks["redis"] = "error"
			degraded = true
		} else {
			checks["redis"] = "ok"
		}
	}

	if len(checks) > 0 {
		resp.Checks = checks
	}

	if degraded {
		resp.Status = "degraded"
	}

	status := http.StatusOK
	if degraded {
		status = http.StatusServiceUnavailable
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck
}
