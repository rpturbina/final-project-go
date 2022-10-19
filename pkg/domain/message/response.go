package message

import "time"

type Response struct {
	Code      int        `json:"code"`
	Message   string     `json:"message,omitempty"`
	Error     string     `json:"error,omitempty"`      // nullable
	Data      any        `json:"data,omitempty"`       // nullable
	StartTime *time.Time `json:"start_time,omitempty"` // nullable
}
