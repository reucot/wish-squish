package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/internal/controller/http/dto"
)

type Error struct {
	ctx echo.Context
	err dto.Error
}

func (e *Error) SetError(err dto.Error) *Error {
	e.err = err
	return e
}

// func (e *Error) SetErrors(err dto.Error) *Error {
// 	e.err = err.Err.()
// 	return e
// }

func (e *Error) SetContext(ctx echo.Context) *Error {
	e.ctx = ctx
	return e
}

func (e *Error) Response(c int) {
	e.ctx.JSON(c, e.err)
}
