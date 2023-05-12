package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundError(c *gin.Context) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		gin.H{"error": "NOTFOUNDERR-1", "message": "Vehicle not found"})
}
