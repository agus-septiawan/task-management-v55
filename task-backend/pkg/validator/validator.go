package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/Mahathirrr/task-management-backend/internal/model"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct memvalidasi struct dan mengembalikan error details
func ValidateStruct(s interface{}) []model.ValidationError {
	var errors []model.ValidationError

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ValidationError
			element.Field = err.Field()
			element.Message = getErrorMessage(err)
			errors = append(errors, element)
		}
	}

	return errors
}

// getErrorMessage mengkonversi validation error ke pesan yang user-friendly
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	case "oneof":
		return fe.Field() + " must be one of: " + fe.Param()
	default:
		return fe.Field() + " is invalid"
	}
}