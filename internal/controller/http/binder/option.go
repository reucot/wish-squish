package binder

import "github.com/reucot/wish-squish/internal/controller/http/validator"

type Option func(*WishBinder)

func SetBinder(b Binder) Option {
	return func(wb *WishBinder) {
		wb.b = b
	}
}

func SetValidator(v validator.Validator) Option {
	return func(wb *WishBinder) {
		wb.v = v
	}
}
