package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/internal/controller/http/dto"
)

func SuccessResponse(ctx echo.Context, s dto.Success) {
	ctx.JSON(http.StatusOK, s)
}
