package go_validator

import (
	"github.com/go-playground/validator/v10"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

type playgroundValidator struct {
	validator *validator.Validate
}

func (self *playgroundValidator) Validate(i interface{}) error {
	return self.validator.Struct(i)
}

func NewGoValidator() (adapters.Validator, error) {
	vald := validator.New(validator.WithRequiredStructEnabled())

	vald.RegisterValidation("handle", handleValidator)
	vald.RegisterValidation("id", idValidator)
	vald.RegisterValidation("id-list", idListValidator)
	vald.RegisterValidation("fullname", fullNameValidator)
	vald.RegisterValidation("xid", xidValidator)
	vald.RegisterValidation("xid-list", xidListValidator)
	vald.RegisterValidation("path", pathValidator)

	return &playgroundValidator{
		validator: vald,
	}, nil
}
