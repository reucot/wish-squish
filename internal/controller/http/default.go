package http

import "github.com/labstack/echo/v4"

func ErrorResponse(c echo.Context, code int, err error) {
	c.JSON(code, err)
}
