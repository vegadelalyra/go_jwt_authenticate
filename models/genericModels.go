package models

type Response struct {
	Data    interface{}   `json:"data,omitempty"`
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Error   []ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}

const (
	ErrorTypeError        = "error"
	ErrorTypeValidation   = "validation"
	ErrorTypeUnauthorized = "unauthorized"
)
