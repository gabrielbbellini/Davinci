package http_error

type UnauthorizedError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		Message: message,
		Code:    401,
	}
}

func (b UnauthorizedError) Error() string {
	return b.Message
}
