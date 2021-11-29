package app

import (
	"github.com/gentildpinto/h-api/internal/config"
	"github.com/gentildpinto/h-api/internal/server"
	"github.com/gentildpinto/h-api/pkg/logger"
)

func Run() (err error) {
	if err = logger.Init(); err != nil {
		return
	}

	cfg := config.New()

	e := server.New()

	server.Run(e, &server.Server{
		Port:         cfg.AppPort,
		Debug:        cfg.AppDebug,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
	})

	return
}
