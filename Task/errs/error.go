package errs

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (a AppError) GetMessage() *AppError {
	return &AppError{
		Message: a.Message,
	}
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}