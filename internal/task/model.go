package task

import (
	"time"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
	StatusCanceled  Status = "canceled"
)

type Task struct {
	ID           string        `json:"id"`
	Status       Status        `json:"status"`
	CreatedAt    time.Time     `json:"created_at"`
	Duration     time.Duration `json:"duration"`
	Result       string        `json:"result,omitempty"`
	ErrorMessage string        `json:"error_message,omitempty"`
}
