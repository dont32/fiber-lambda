package errs

import (
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    fiber.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    fiber.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    fiber.StatusUnprocessableEntity,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    fiber.StatusUnauthorized,
	}

}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    fiber.StatusForbidden,
	}
}
