package go_validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var fullNameRegex = regexp.MustCompile(`^[A-Za-záàâãéèêíóôõúçÁÀÂÃÉÈÍÓÔÕÚÇ ]+$`)

func fullNameValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if v.Kind() != reflect.String {
		return false
	}

	value := v.String()

	if !strings.Contains(value, " ") {
		return false
	}

	if len(value) < 3 {
		return false
	}

	return fullNameRegex.MatchString(value)
}
