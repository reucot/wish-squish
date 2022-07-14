package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ErrorResponse(c echo.Context, code int, err error) {
	c.JSON(code, err)
}
