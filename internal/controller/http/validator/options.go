package validator

import (
	"unicode"

	"github.com/go-playground/validator"
)

type Option func(*WishValidator)

func RegisterValidatorForPassword(tag string) Option {
	return func(wv *WishValidator) {
		wv.validator.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
			return isValidPwd(fl.Field().String())
		})
	}
}

func isValidPwd(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
