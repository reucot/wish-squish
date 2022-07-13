package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"

	"github.com/reucot/wish-squish/config"
	"github.com/reucot/wish-squish/pkg/httpserver"
	log "github.com/reucot/wish-squish/pkg/logger"
)

func Run() {

	err := config.Load()
	if err != nil {
		log.Error("internal - app - main.go - Run() - config.Load(): %w", err)
		return
	}

	// HTTP Server
	handler := echo.New()
	//v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(config.Get().Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("internal - app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error("internal - app - Run - httpServer.Notify: %w", err)
	}
}
