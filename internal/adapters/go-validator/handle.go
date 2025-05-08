package go_validator

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var allowedHandleCharacters = regexp.MustCompile("^[a-z0-9]*$")

func handleValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if v.Kind() != reflect.String {
		return false
	}

	length := v.Len()

	if length > 16 || length < 3 {
		return false
	}

	value := v.String()

	return allowedHandleCharacters.Match([]byte(value))
}
