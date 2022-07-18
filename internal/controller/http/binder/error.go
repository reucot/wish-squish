package binder

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ErrorField struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

const _defaultCapacity int = 4

type Error struct {
	Errors []ErrorField `json:"errors"`
}

func ValidationErrors(err error) error {

	errs := err.(validator.ValidationErrors)

	e := make([]ErrorField, 0, _defaultCapacity)

	for _, err := range errs {
		e = append(e, ErrorField{
			Field: err.Field(),
			Tag:   err.Tag(),
		})
	}

	return &Error{Errors: e}
}

func (e *Error) Error() string {
	s := ""
	for _, err := range e.Errors {
		s += fmt.Sprintf("Field \"%s\": on tag \"%s\"\n", err.Field, err.Tag)
	}
	return s
}
