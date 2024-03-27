package validation

import (
	"time"

	"github.com/go-playground/validator/v10"

	e "github.com/duartqx/ttgowebddv2/src/common/errors"
)

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	v.RegisterValidation("future", func(fl validator.FieldLevel) bool {
		t, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false
		}
		return t.After(time.Now())
	})

	return &Validator{Validate: v}
}

func (v Validator) ValidateStruct(s interface{}) error {
	if errs := v.Struct(s); errs != nil {
		return e.ValidationErrors{Errs: errs}
	}
	return nil
}
