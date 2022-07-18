package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"

	"github.com/reucot/wish-squish/config"
	v1 "github.com/reucot/wish-squish/internal/controller/http"
	"github.com/reucot/wish-squish/pkg/httpserver"
	log "github.com/reucot/wish-squish/pkg/logger"
)

func Run() {

	err := config.Load()
	if err != nil {
		log.Error("internal - app - main.go - Run() - config.Load(): %s", err.Error())
		return
	}

	// HTTP Server
	handler := echo.New()
	v1.NewRouter(handler)
	httpServer := httpserver.New(handler, httpserver.Port(config.Get().Port))

	log.Info("Server starting http://%s:%s", config.Get().Host, config.Get().Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("internal - app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error("internal - app - Run - httpServer.Notify(): %s", err.Error())
	}

	if err = httpServer.Shutdown(); err != nil {
		log.Error("internal - app - main.go - Run() - httpServer.Shutdown(): %s", err.Error())
		return
	}

}
