package app

import (
	"github.com/reucot/wish-squish/config"
	"github.com/reucot/wish-squish/internal/pkg/log"
)

func Run() {
	var err error

	if err = config.Load(); err != nil {
		log.Error("internal - app - main.go - Run() - config.Load(): %w", err)
		return
	}

	if err = log.OpenLog(); err != nil {
		log.Error("internal - app - main.go - Run() - logger.OpenLog(): %w", err)
		return
	}

}
