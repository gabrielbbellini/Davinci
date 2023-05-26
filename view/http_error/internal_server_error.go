package http_error

const (
	InternalErrorMessage = "Ocorreu um erro inesperado."
)

type InternalServerError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		Message: message,
		Code:    500,
	}
}

func (b InternalServerError) Error() string {
	return b.Message
}
