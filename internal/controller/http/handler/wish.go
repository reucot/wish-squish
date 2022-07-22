package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/internal/service"
)

type WishHandler struct {
	services *service.Service
}

func NewWishHandler(e *echo.Echo) {
	wh := WishHandler{}

	e.POST("wish/create", wh.Create)
}

func (h *WishHandler) Create(c echo.Context) error {

	return nil
}
