package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/reucot/wish-squish/internal/controller/http/binder"
	"github.com/reucot/wish-squish/internal/controller/http/dto"
	"github.com/reucot/wish-squish/internal/service"
	log "github.com/reucot/wish-squish/pkg/logger"
)

type AuthorizationHandler struct {
	as service.AuthorizationService
}

func NewAuthorization(e *echo.Echo) {
	ah := AuthorizationHandler{}

	e.POST("/signin", ah.SignIn)
	e.POST("/signup", ah.SignUp)
}

func (h AuthorizationHandler) SignIn(c echo.Context) error {
	var req dto.SignIn
	var err error

	if err = c.Bind(&req); err != nil {
		log.Error("internal - controller - http - authorization - SignIn() - c.Bind(req): %s", err.Error())
		e := dto.NewError(binder.ValidationErrors(err))
		ErrorResponse(c, http.StatusBadRequest, e)
	}

	c.Redirect(http.StatusFound, "/")

	return nil
}

func (h AuthorizationHandler) SignUp(c echo.Context) error {
	var req dto.SignUp
	var err error

	if err = c.Bind(&req); err != nil {
		log.Error("internal - controller - http - authorization - SignUp() - c.Bind(req): %s", err.Error())
		//ErrorResponse(c, http.StatusBadRequest, dto.Error{Message: err.Error(), Error: binder.ValidationErrors(err)})
		return err
	}

	log.Info("%v", req)

	c.Redirect(http.StatusFound, "/")

	return nil
}
