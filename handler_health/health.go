package handler_health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

// Health godoc
// @Summary handler_health check endpoint
// @Schemes
// @Description returns the status of the service
// @Tags handler_health
// @Accept json
// @Produce json
// @Success 200 {object} Status
// @Router /handler_health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, Status{Status: "OK"})
}
