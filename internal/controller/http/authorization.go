package http

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

	e.POST("/signin", ah.Signin)
}

func (h AuthorizationHandler) Signin(c echo.Context) error {
	req := dto.Signin{}

	err := c.Bind(&req)
	if err != nil {
		log.Error("internal - controller - http - authorization - Signin() - c.Bind(req): %s", err.Error())
		ErrorResponse(c, http.StatusBadRequest, err)
	}

	log.Info("%v", req)

	return nil
}
