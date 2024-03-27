package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	BadRequestError = errors.New("Bad Request")
	ForbiddenError  = errors.New("Forbidden")
	InternalError   = errors.New("Internal")
	NotFoundError   = errors.New("Not Found")
	Unauthorized    = errors.New("Unauthorized")
)

type validationError struct {
	Tag   string
	Value interface{}
}

type ValidationErrors struct {
	Errs error
}

func (ves ValidationErrors) Error() string {
	errStr := ""
	for tag, value := range *ves.Decode() {
		errStr += fmt.Sprintf("%s: %s ", tag, value)
	}
	return errStr
}

func (ves ValidationErrors) Decode() *map[string]validationError {

	validationErrors := map[string]validationError{}

	for _, err := range ves.Errs.(validator.ValidationErrors) {
		validationErrors[strings.ToLower(err.Field())] = validationError{
			Tag:   err.Tag(),
			Value: err.Value(),
		}
	}

	return &validationErrors
}

func (ves ValidationErrors) DecodeJSON() *[]byte {
	res, _ := json.Marshal(ves.Decode())
	return &res
}
