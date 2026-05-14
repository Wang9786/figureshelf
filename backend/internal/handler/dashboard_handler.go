package handler

import (
	"net/http"

	"figureshelf-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService *service.DashboardService
}

func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	userID := c.GetString("user_id")

	summary, err := h.dashboardService.GetSummary(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get dashboard summary",
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}