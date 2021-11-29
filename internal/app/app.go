package app

import (
	"github.com/gentildpinto/h-api/internal/config"
	"github.com/gentildpinto/h-api/internal/repository"
	"github.com/gentildpinto/h-api/internal/server"
	"github.com/gentildpinto/h-api/internal/service"
	"github.com/gentildpinto/h-api/pkg/database/postgresql"
	"github.com/gentildpinto/h-api/pkg/logger"
)

func Run() (err error) {
	if err = logger.Init(); err != nil {
		return
	}

	cfg := config.New()

	if err != nil {
		return
	}

	db, err := postgresql.New(cfg.Postgresql)

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Dependencies{
		Repos: repos,
	})

	e := server.New(services)

	server.Run(e, &server.Server{
		Port:         cfg.AppPort,
		Debug:        cfg.AppDebug,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
	})

	return
}
