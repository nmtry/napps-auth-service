package error

import (
	"encoding/json"
	"net/http"
	"time"
)

func WriteError(w http.ResponseWriter, status int, message string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:     message,
		Details:   details,
		Status:    status,
		Timestamp: time.Now(),
	})
}
