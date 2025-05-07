package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("custom", custom)
}

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}

// Custom regex validation for practice
func custom(fl validator.FieldLevel) bool {
	custom := fl.Field().String()
	customRegex := `^\d{9}$` // Only 9 digits
	match, _ := regexp.MatchString(customRegex, custom)
	return match
}
