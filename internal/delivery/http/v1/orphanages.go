package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) initOrphanagesRoutes(e *echo.Group) {
	orphanages := e.Group("/orphanages")
	{
		orphanages.GET("", getOrphanages(h))
		orphanages.GET("/:id", getOrphanage(h))
	}
}

func getOrphanages(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		orphanages, err := h.services.Orphanages.All()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, orphanages)
	}
}

func getOrphanage(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		orphanage, err := h.services.Orphanages.FindByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, orphanage)
	}
}
