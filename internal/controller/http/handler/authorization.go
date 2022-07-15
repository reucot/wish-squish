package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reucot/wish-squish/internal/controller/http/dto"
	log "github.com/reucot/wish-squish/pkg/logger"
)

type AuthorizationHandler struct {
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
		ErrorResponse(c, http.StatusBadRequest, dto.Error{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		log.Error("internal - controller - http - authorization - SignIn() - c.Validate(req): %s", err.Error())
		ErrorResponse(c, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return err
	}

	log.Info("%v", req)

	c.Redirect(http.StatusFound, "/")

	return nil
}

func (h AuthorizationHandler) SignUp(c echo.Context) error {
	var req dto.SignUp
	var err error

	if err = c.Bind(&req); err != nil {
		log.Error("internal - controller - http - authorization - SignUp() - c.Bind(req): %s", err.Error())
		ErrorResponse(c, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return err
	}

	if err = c.Validate(req); err != nil {
		log.Error("internal - controller - http - authorization - SignUp() - c.Validate(req): %s", err.Error())
		ErrorResponse(c, http.StatusBadRequest, dto.Error{Message: err.Error()})
		return err
	}

	log.Info("%v", req)

	c.Redirect(http.StatusFound, "/")

	return nil
}
