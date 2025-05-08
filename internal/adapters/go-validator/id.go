package go_validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func idValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if !v.CanInt() {
		return false
	}

	value := v.Int()

	return value >= 0
}

func idListValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if v.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)

		if !item.CanInt() {
			return false
		}

		value := v.Int()

		if value < 0 {
			return false
		}
	}

	return true
}
