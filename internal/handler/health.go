package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status" example:"ok"`
	Timestamp string `json:"timestamp" example:"2024-01-01T00:00:00Z"`
}

// HealthCheck godoc
// @Summary      Show health status
// @Description  Get the health status of the API service
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  HealthResponse
// @Router       /healthz [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
