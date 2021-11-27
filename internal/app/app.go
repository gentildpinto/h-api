package app

import (
	"os"

	"github.com/gentildpinto/h-api/internal/server"
	"github.com/gentildpinto/h-api/pkg/logger"
)

var (
	appPort = "80"
)

func Run() (err error) {
	if err = logger.Init(); err != nil {
		return
	}

	if os.Getenv("PORT") != "" {
		appPort = os.Getenv("PORT")
	}

	debug := os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true"

	s := server.New(appPort, 60, 60, debug)

	s.Run()

	return
}
