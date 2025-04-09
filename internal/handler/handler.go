package handler

import (
	"click-counter/internal/model"
	"click-counter/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	s *service.Service
}

func NewCounterHandler(c *service.Service) *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())

	h := &Handler{
		c,
	}

	app.GET("/counter/:bannerID", h.incrementClick)
	app.POST("/stats/:bannerID", h.getStats)

	return app

}

func (h *Handler) incrementClick(c echo.Context) error {
	bannerID := c.Param("bannerID")
	bannerIDInt, err := strconv.Atoi(bannerID)
	if err != nil {
		h.s.Logger.Errorf("failed convert bannerID to int: %v", err)
		return c.String(http.StatusBadRequest, "invalid banner ID")
	}

	err = h.s.IncrementClick(bannerIDInt)
	if err != nil {
		h.s.Logger.Errorf("failed incrementing click: %v", err)
		return c.String(http.StatusInternalServerError, "failed to increment click")
	}

	return c.String(http.StatusOK, fmt.Sprintf("click for banner %d incremented", bannerIDInt))
}

func (h *Handler) getStats(c echo.Context) error {
	bannerID := c.Param("bannerID")
	bannerIDInt, err := strconv.Atoi(bannerID)
	if err != nil {
		h.s.Logger.Errorf("failed convert bannerID to int: %v", err)
		return c.String(http.StatusBadRequest, "invalid banner ID")
	}

	params := model.Click{}
	err = c.Bind(&params)
	if err != nil {
		h.s.Logger.Errorf("failed bind params: %v", err)
		return c.String(http.StatusBadRequest, "invalid params")
	}

	stats, err := h.s.GetStats(bannerIDInt, params.TsFrom, params.TsTo)
	if err != nil {
		h.s.Logger.Errorf("failed getting stat: %v", err)
		return c.String(http.StatusInternalServerError, "failed to get stat")
	}

	var response model.StatsResponse
	for _, stat := range stats {
		response.Stats = append(response.Stats, model.StatResponse{
			Timestamp: stat.Timestamp.Format(time.RFC3339),
			Value:     stat.Count,
		})
	}

	return c.JSON(http.StatusOK, response)

}
