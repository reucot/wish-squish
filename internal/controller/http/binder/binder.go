package binder

import (
	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/internal/controller/http/validator"
)

type Binder interface {
	Bind(i interface{}, c echo.Context) error
}

type WishBinder struct {
	b Binder
	v validator.Validator
}

func NewWishBinder(opts ...Option) *WishBinder {

	wb := &WishBinder{}

	for _, opt := range opts {
		opt(wb)
	}

	return wb
}

func (wb *WishBinder) Bind(i interface{}, c echo.Context) error {
	var err error

	if err = wb.b.Bind(i, c); err != nil {
		return err
	}

	if err = wb.v.Validate(i); err != nil {
		return err
	}

	return nil
}
