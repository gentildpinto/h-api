package v1

import (
	"net/http"
	"strconv"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) initOrphanagesRoutes(e *echo.Group) {
	orphanages := e.Group("/orphanages")
	{
		orphanages.GET("", getOrphanages(h))
		orphanages.GET("/:id", getOrphanage(h))
		orphanages.POST("", createOrphanage(h))
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

func createOrphanage(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		latitude, _ := strconv.ParseFloat(c.FormValue("latitude"), 64)
		longitude, _ := strconv.ParseFloat(c.FormValue("longitude"), 64)
		openOnWeekends, _ := strconv.ParseBool(c.FormValue("open_on_weekends"))

		orphanage := domain.Orphanage{
			Name:           c.FormValue("name"),
			Latitude:       latitude,
			Longitude:      longitude,
			About:          c.FormValue("about"),
			Instructions:   c.FormValue("instructions"),
			OpenedHours:    c.FormValue("opened_hours"),
			OpenOnWeekends: openOnWeekends,
		}

		form, err := c.MultipartForm()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		images := form.File["images"]

		if err := h.services.Orphanages.Create(&orphanage, images); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, orphanage)
	}
}
