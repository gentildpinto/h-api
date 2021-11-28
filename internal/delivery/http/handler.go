package http

import (
	v1 "github.com/gentildpinto/h-api/internal/delivery/http/v1"
	"github.com/gentildpinto/h-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	handlerV1 := v1.NewHandler(h.services)
	api := e.Group("/api")
	{
		handlerV1.Init(api)
	}
}
