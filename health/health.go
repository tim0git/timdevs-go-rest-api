package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

// Health godoc
// @Summary health check endpoint
// @Schemes
// @Description returns the status of the service
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} Status
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, Status{Status: "OK"})
}
