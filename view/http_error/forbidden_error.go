package http_error

type ForbiddenError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewForbiddenError(message string) ForbiddenError {
	return ForbiddenError{
		Message: message,
		Code:    403,
	}
}

func (b ForbiddenError) Error() string {
	return b.Message
}
