package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) initOrphanagesRoutes(e *echo.Group) {
	orphanages := e.Group("/orphanages")
	{
		orphanages.GET("/", h.getOrphanages)
	}
}

func (h *Handler) getOrphanages(c echo.Context) error {
	orphanages, err := h.services.Orphanages.All()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, orphanages)
}
