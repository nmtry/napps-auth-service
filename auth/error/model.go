package error

import (
	"time"
)

type ErrorResponse struct {
	Error     string    `json:"error"`
	Details   string    `json:"details,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
}
