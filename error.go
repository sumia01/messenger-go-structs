package messenger

import "fmt"

type rawError struct {
	Error Error `json:"error"`
}

// Error ...
type Error struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      int    `json:"code"`
	ErrorData string `json:"error_data"`
	TraceID   string `json:"fbtrace_id"`
}

// Error ...
func (e Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
