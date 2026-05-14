package handler

import (
	"net/http"
	"strconv"

	"figureshelf-backend/internal/model"
	"figureshelf-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type FigureHandler struct {
	figureService *service.FigureService
}

func NewFigureHandler(figureService *service.FigureService) *FigureHandler {
	return &FigureHandler{
		figureService: figureService,
	}
}

func (h *FigureHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.CreateFigureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	figure, err := h.figureService.Create(c.Request.Context(), userID, req)
	if err != nil {
		if err.Error() == "invalid figure status" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid figure status",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create figure",
		})
		return
	}

	c.JSON(http.StatusCreated, figure)
}

func (h *FigureHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")

	figures, err := h.figureService.List(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list figures",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": figures,
		"total": len(figures),
	})
}

func (h *FigureHandler) GetByID(c *gin.Context) {
	userID := c.GetString("user_id")
	figureID := c.Param("id")

	figure, err := h.figureService.GetByID(c.Request.Context(), userID, figureID)
	if err != nil {
		if err.Error() == "figure not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "figure not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get figure",
		})
		return
	}

	c.JSON(http.StatusOK, figure)
}

func (h *FigureHandler) Update(c *gin.Context) {
	userID := c.GetString("user_id")
	figureID := c.Param("id")

	var req model.UpdateFigureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	figure, err := h.figureService.Update(c.Request.Context(), userID, figureID, req)
	if err != nil {
		if err.Error() == "invalid figure status" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid figure status",
			})
			return
		}

		if err.Error() == "figure not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "figure not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update figure",
		})
		return
	}

	c.JSON(http.StatusOK, figure)
}

func (h *FigureHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	figureID := c.Param("id")

	err := h.figureService.Delete(c.Request.Context(), userID, figureID)
	if err != nil {
		if err.Error() == "figure not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "figure not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete figure",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *FigureHandler) ListUpcomingPayments(c *gin.Context) {
	userID := c.GetString("user_id")

	days := 30
	if value := c.Query("days"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			days = parsed
		}
	}

	figures, err := h.figureService.ListUpcomingPayments(c.Request.Context(), userID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list upcoming payments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": figures,
		"total": len(figures),
		"days":  days,
	})
}

func (h *FigureHandler) ListUpcomingReleases(c *gin.Context) {
	userID := c.GetString("user_id")

	days := 60
	if value := c.Query("days"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			days = parsed
		}
	}

	figures, err := h.figureService.ListUpcomingReleases(c.Request.Context(), userID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list upcoming releases",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": figures,
		"total": len(figures),
		"days":  days,
	})
}

// func (h *FigureHandler) GetDashboardSummary(c *gin.Context) {
// 	userID := c.GetString("user_id")

// 	summary, err := h.figureService.GetDashboardSummary(c.Request.Context(), userID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "failed to get dashboard summary",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, summary)
// }