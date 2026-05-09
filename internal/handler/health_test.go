package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		checkBody      bool
	}{
		{
			name:           "Valid GET request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			checkBody:      true,
		},
		{
			name:           "Invalid POST request",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			checkBody:      false,
		},
		{
			name:           "Invalid PUT request",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			checkBody:      false,
		},
		{
			name:           "Invalid DELETE request",
			method:         http.MethodDelete,
			expectedStatus: http.StatusMethodNotAllowed,
			checkBody:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/healthz", nil)
			w := httptest.NewRecorder()

			HealthCheck(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.checkBody {
				var response HealthResponse
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}

				if response.Status != "ok" {
					t.Errorf("Expected status 'ok', got '%s'", response.Status)
				}

				if response.Timestamp == "" {
					t.Error("Expected non-empty timestamp")
				}
			}
		})
	}
}

func TestHealthCheckResponseHeaders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	HealthCheck(w, req)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}
}

func TestHealthCheckConcurrent(t *testing.T) {
	concurrency := 100
	done := make(chan bool, concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
			w := httptest.NewRecorder()
			HealthCheck(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", w.Code)
			}
			done <- true
		}()
	}

	for i := 0; i < concurrency; i++ {
		<-done
	}
}
