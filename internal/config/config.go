package config

import (
	"os"
	"strconv"
)

var (
	appPort            = "8000"
	debugApp           = false
	serverReadTimeout  = 60
	serverWriteTimeout = 60
)

type Config struct {
	AppPort            string
	AppDebug           bool
	ServerReadTimeout  int
	ServerWriteTimeout int
}

func New() *Config {
	if os.Getenv("PORT") != "" {
		appPort = os.Getenv("PORT")
	}

	if os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true" {
		debugApp = true
	}

	if os.Getenv("SERVER_READ_TIMEOUT") != "" {
		serverReadTimeout, _ = strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	}

	if os.Getenv("SERVER_WRITE_TIMEOUT") != "" {
		serverWriteTimeout, _ = strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	}

	return &Config{
		AppPort:            appPort,
		AppDebug:           debugApp,
		ServerReadTimeout:  serverReadTimeout,
		ServerWriteTimeout: serverWriteTimeout,
	}
}
