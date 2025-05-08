package go_validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var allowedPathCharacters = regexp.MustCompile("^[a-z0-9/-]*$")

func pathValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if v.Kind() != reflect.String {
		return false
	}

	value := v.String()

	if !strings.HasPrefix(value, "/") {
		return false
	}

	if strings.HasSuffix(value, "/") {
		return false
	}

	return allowedPathCharacters.Match([]byte(value))
}
