package router

import (
	"net/http"

	"github.com/sharechat/revamp/backend/handler"
)

// New returns an http.Handler with all API routes registered.
func New() http.Handler {
	mux := http.NewServeMux()

	// v4 routes
	mux.HandleFunc("GET /api/v4/ready", handler.Ready)

	return mux
}
