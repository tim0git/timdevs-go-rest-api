package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidationError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{"error": "VALIDATEERR-1", "message": err.Error()})
}

func DynamoDBError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"error": "DYNAMOERR-1", "message": err.Error()})
}
