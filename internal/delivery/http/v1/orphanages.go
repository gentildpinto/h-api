package v1

import (
	"net/http"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/labstack/echo/v4"
)

type createOrphanagePayload struct {
	Name           string   `json:"name" validate:"required"`
	Latitude       float64  `json:"latitude" validate:"required"`
	Longitude      float64  `json:"longitude" validate:"required"`
	About          string   `json:"about" validate:"required"`
	Instructions   string   `json:"instructions" validate:"required"`
	OpenedHours    string   `json:"opened_hours" validate:"required"`
	OpenOnWeekends bool     `json:"open_on_weekends" validate:"required"`
	Images         []string `json:"images" validate:"required"`
}

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
		orphanageRequest := createOrphanagePayload{}

		if err := c.Bind(&orphanageRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		orphanage := domain.Orphanage{
			Name:           orphanageRequest.Name,
			Latitude:       orphanageRequest.Latitude,
			Longitude:      orphanageRequest.Longitude,
			About:          orphanageRequest.About,
			Instructions:   orphanageRequest.Instructions,
			OpenedHours:    orphanageRequest.OpenedHours,
			OpenOnWeekends: orphanageRequest.OpenOnWeekends,
			Images:         orphanageRequest.Images,
		}

		if err := h.services.Orphanages.Create(&orphanage); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, orphanage)
	}
}
