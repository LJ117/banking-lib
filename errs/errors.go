package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (ae AppError) AsMessage() *AppError {
	return &AppError{
		Message: ae.Message,
	}
}

// NewNotFoundError  自定义数据未找到错误
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// NewUnexpectedError 自定义系统错误
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity, // CODE 400 ，过了之后，到业务层无法处理的错误代码
		Message: message,
	}
}

// NewAuthenticationError 未做权限校验
func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

// NewAuthorizationError 权限不足
func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}
