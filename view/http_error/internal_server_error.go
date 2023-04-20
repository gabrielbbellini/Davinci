package http_error

type InternalServerError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		Message: message,
		Code:    400,
	}
}

func (b InternalServerError) Error() string {
	return b.Message
}
