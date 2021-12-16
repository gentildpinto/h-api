package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	delivery "github.com/gentildpinto/h-api/internal/delivery/http"
	"github.com/gentildpinto/h-api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port         string
	Debug        bool
	ReadTimeout  int
	WriteTimeout int
}

func New(services *service.Services) *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}\n"}),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
	)

	handlers := delivery.NewHandler(services)

	handlers.InitRoutes(e)

	return e
}

func Run(e *echo.Echo, s *Server) {
	srvr := &http.Server{
		Addr:         ":" + s.Port,
		ReadTimeout:  time.Duration(s.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.WriteTimeout) * time.Second,
	}

	e.Debug = s.Debug

	go func() {
		if err := e.StartServer(srvr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("server shutdown", err)
	}
}
