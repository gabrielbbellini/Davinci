package http_error

type BadRequestError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		Message: message,
		Code:    400,
	}
}

func (b BadRequestError) Error() string {
	return b.Message
}
