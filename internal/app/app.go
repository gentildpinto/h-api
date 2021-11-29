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

	s := server.New(cfg)

	s.Run()

	return
}
