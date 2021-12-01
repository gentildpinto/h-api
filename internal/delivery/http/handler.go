package http

import (
	"net/http"

	v1 "github.com/gentildpinto/h-api/internal/delivery/http/v1"
	"github.com/gentildpinto/h-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	e.GET("/", func(ec echo.Context) error {
		return ec.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World from happy!",
		})
	})

	e.Static("/public/uploads", "public/uploads")

	handlerV1 := v1.NewHandler(h.services)
	api := e.Group("/api")
	{
		handlerV1.Init(api)
	}
}
