package helpers

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/gofiber/fiber/v2"
)

type ValidationErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func NewValidationErrorResponse(errorMessage string, fieldErrors interface{}) ValidationErrorResponse {
	return ValidationErrorResponse{
		Status:  false,
		Message: errorMessage,
		Errors:  fieldErrors,
	}
}

func CreateErrorResponse(code int, message string, errors *interface{}) appctx.Response {
	if errors == nil {
		return *appctx.NewResponse().
			WithCode(code).
			WithMessage(message)
	}
	return *appctx.NewResponse().
		WithCode(code).
		WithMessage(message).
		WithError(&errors)
}

func CreateValidationResponse(field string, errorMessage string) appctx.Response {
	errorMap := make(map[string]interface{})
	errorMap[field] = []string{errorMessage}

	response := NewValidationErrorResponse("The given data was invalid", errorMap)
	return CreateErrorResponse(fiber.StatusUnprocessableEntity, response.Message, &response.Errors)
}
