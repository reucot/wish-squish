package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/internal/controller/http/dto"
)

func ErrorResponse(ctx echo.Context, c int, e dto.Error) {
	ctx.JSON(c, e)
}
