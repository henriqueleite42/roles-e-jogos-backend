package go_validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

func xidValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if v.Kind() != reflect.String {
		return false
	}

	value := v.String()

	_, err := xid.FromString(value)

	return err == nil
}

func xidListValidator(fl validator.FieldLevel) bool {
	v := fl.Field()

	if v.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)

		if item.Kind() != reflect.String {
			return false
		}

		value := item.String()

		_, err := xid.FromString(value)

		if err != nil {
			return false
		}
	}

	return true
}
