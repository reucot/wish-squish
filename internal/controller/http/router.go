package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reucot/wish-squish/config"
)

func NewRouter(e *echo.Echo) {

	if config.Get().Mode == "develop" {
		e.Debug = true
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
	}

	NewPages(e)
	NewAuthorization(e)

}
