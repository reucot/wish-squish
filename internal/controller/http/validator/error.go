package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ErrorField struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
	Msg   string      `json:"error_msg"`
}

const _defaultCapacity int = 4

type Error struct {
	Errors []ErrorField `json:"errors"`
}

func (e *ErrorField) Error() string {
	return e.Msg
}

func GetErrors(err error) error {

	errs := err.(validator.ValidationErrors)

	e := make([]ErrorField, 0, _defaultCapacity)

	for _, err := range errs {
		e = append(e, ErrorField{
			Field: err.Field(),
			Value: err.Value(),
			Tag:   err.Tag(),
			Msg:   fmt.Sprintf(`Field '%s' failed: on tag '%s', current value: '%v'`, err.Field(), err.Tag(), err.Value()),
		})
	}

	return &Error{Errors: e}
}

func (e *Error) Error() string {
	return "Validation errors"
}
