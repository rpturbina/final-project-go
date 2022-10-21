package message

type ErrorType string
type ErrorMessage struct {
	Error error     `json:"error_message,omitempty"`
	Type  ErrorType `json:"error_type"`
}
