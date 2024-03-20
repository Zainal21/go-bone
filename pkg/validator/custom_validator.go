package validator

import (
	vldtr "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type (
	CustomValidator struct {
		validator *vldtr.Validate
	}

	ErrorValidation struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

func NewCustomValidator() *CustomValidator {
	// Create a new validator for a Book model.
	validate := vldtr.New()

	_ = validate.RegisterValidation("uuid", validationUuid)
	_ = validate.RegisterValidation("dateformat", validationDateFormat)

	// add more custom validation if needed

	return &CustomValidator{validator: validate}
}
func (v *CustomValidator) Validate(data any) []ErrorValidation {
	var validationErrors []ErrorValidation

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(vldtr.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorValidation
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func validationUuid(fl vldtr.FieldLevel) bool {
	field := fl.Field().String()
	if _, err := uuid.Parse(field); err != nil {
		return true
	}
	return false
}

func validationDateFormat(fl vldtr.FieldLevel) bool {
	field := fl.Field().String()
	if _, err := time.Parse("2006-01-02", field); err != nil {
		return false
	}
	return true
}
