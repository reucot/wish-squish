package validator

import (
	"github.com/go-playground/validator"
)

type Validator interface {
	Validate(i interface{}) error
}

type WishValidator struct {
	validator *validator.Validate
}

func NewWishValidator(opts ...Option) *WishValidator {
	wv := &WishValidator{
		validator: validator.New(),
	}

	for _, opt := range opts {
		opt(wv)
	}

	return wv
}

func (cv *WishValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
