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

type (
	Postgresql struct {
		User     string
		Host     string
		Port     string
		DBName   string
		SSLMode  string
		Password string
	}
	Config struct {
		AppPort            string
		AppDebug           bool
		ServerReadTimeout  int
		ServerWriteTimeout int
		Postgresql         Postgresql
	}
)

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
		Postgresql: Postgresql{
			User:     os.Getenv("DATABASE_USERNAME"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			DBName:   os.Getenv("DATABASE_NAME"),
			SSLMode:  os.Getenv("DATABASE_SSL_MODE"),
			Password: os.Getenv("DATABASE_PASSWORD"),
		},
	}
}
