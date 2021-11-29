package v1

import (
	"net/http"

	"github.com/gentildpinto/h-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(e *echo.Group) {
	v1 := e.Group("/v1")
	{
		v1.GET("/ping", func(ec echo.Context) error {
			return ec.String(http.StatusOK, "pong")
		})

		h.initOrphanagesRoutes(v1)
	}
}
