package helper

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) ValidateStruct(data interface{}) error {
	return v.validate.Struct(data)
}

func FormatValidationError(err error) string {
	var errMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		errMessages = append(errMessages, err.Field()+" is invalid: "+err.Tag())
	}
	
	// for _, err := range err.(validator.ValidationErrors) {
	// 	field := err.Field()
	// 	tag := err.Tag()

	// 	var message string
	// 	switch field + "_" + tag {
	// 	case "Status_required":
	// 		message = "Status is required and must be either 'ok' or 'cancel'"
	// 	case "Status_oneof":
	// 		message = "Status must be either 'ok' or 'cancel'"
	// 	case "Email_required":
	// 		message = "Email is required"
	// 	case "Email_email":
	// 		message = "Email format is invalid"
	// 	case "Phone_required":
	// 		message = "Phone number is required"
	// 	case "Phone_len":
	// 		message = "Phone number must be exactly 10 digits"
	// 	case "Phone_numeric":
	// 		message = "Phone number must be numeric"
	// 	case "EventID_required":
	// 		message = "Event ID is required"
	// 	default:
	// 		message = field + " is invalid: " + tag
	// 	}

	// 	errMessages = append(errMessages, message)
	// }

	return strings.Join(errMessages, ", ")
}
