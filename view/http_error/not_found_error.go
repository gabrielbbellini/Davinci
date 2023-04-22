package http_error

type NotFoundError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		Message: message,
		Code:    404,
	}
}

func (b NotFoundError) Error() string {
	return b.Message
}
