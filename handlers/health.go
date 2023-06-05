package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthStatus struct {
	Status string `json:"status"`
}

// Health godoc
// @Summary health check endpoint
// @Schemes
// @Description returns the status of the service
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthStatus
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthStatus{Status: "OK"})
}
