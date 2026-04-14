package handler

import (
	"encoding/json"
	"net/http"
)

// Ready handles GET /api/v4/ready.
// Returns 200 OK with {"ready": true} once the server is accepting requests.
func Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ready": true})
}
