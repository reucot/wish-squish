package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/reucot/wish-squish/config"
	"github.com/reucot/wish-squish/internal/controller/http/binder"
	"github.com/reucot/wish-squish/internal/controller/http/handler"
	"github.com/reucot/wish-squish/internal/controller/http/validator"
)

func NewRouter(e *echo.Echo) {

	if config.Get().Mode == "develop" {
		e.Debug = true
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
	}

	e.Validator = validator.NewWishValidator(validator.RegisterValidatorForPassword("password"))
	e.Binder = binder.NewWishBinder(binder.SetBinder(&echo.DefaultBinder{}), binder.SetValidator(e.Validator))

	handler.NewPages(e)
	handler.NewAuthorization(e)
}
