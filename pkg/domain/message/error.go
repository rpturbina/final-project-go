package message

type ErrorType string
type ErrorMessage struct {
	Error error     `json:"error_message"`
	Type  ErrorType `json:"error_type"`
}
