package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port          string
	ReadTimeout   int
	WriteTimeout  int
	HttpFramework *echo.Echo
}

func New(port string, readTimeout int, writeTimeout int, debug bool) *Server {
	e := echo.New()
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}\n"}),
		middleware.Recover(),
	)

	e.Debug = debug

	return &Server{
		Port:          port,
		ReadTimeout:   readTimeout,
		WriteTimeout:  writeTimeout,
		HttpFramework: e,
	}
}

func (s *Server) Run() {
	srvr := &http.Server{
		Addr:         ":" + s.Port,
		ReadTimeout:  time.Duration(s.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.WriteTimeout) * time.Second,
	}

	go func() {
		if err := s.HttpFramework.StartServer(srvr); err != nil && err != http.ErrServerClosed {
			s.HttpFramework.Logger.Fatal("shutting down the server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.HttpFramework.Shutdown(ctx); err != nil {
		s.HttpFramework.Logger.Fatal("server shutdown", err)
	}
}
