package http_error

const (
	UnauthorizedMessage = "Você não tem permissão para acessar este recurso."
)

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
