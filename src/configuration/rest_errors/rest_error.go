package rest_error

import "net/http"

type RestError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Reasons []Reason `json:"reasons,omitempty"`
}

type Reason struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func New(message, err string, code int, reasons []Reason) *RestError {
	return &RestError{
		Message: message,
		Err:     err,
		Code:    code,
		Reasons: reasons,
	}
}

func NewBadRequestError(message string) *RestError {
	return New(message, "bad_request", http.StatusBadRequest, nil)
}

func NewBadRequestValidationError(message string, reasons []Reason) *RestError {
	return New(message, "bad_request", http.StatusBadRequest, reasons)
}

func NewInternalServerError(message string) *RestError {
	return New(message, "internal_server_error", http.StatusInternalServerError, nil)
}

func NewNotFoundError(message string) *RestError {
	return New(message, "not_found", http.StatusNotFound, nil)
}

func NewForbiddenError(message string) *RestError {
	return New(message, "forbidden", http.StatusForbidden, nil)
}
