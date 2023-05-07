package controllers

import (
	"net/http"
	"timdevs.rest.api.com/m/v2/database"
	"timdevs.rest.api.com/m/v2/validators"

	"github.com/gin-gonic/gin"
)

func RegisterVehicle(c *gin.Context) {

	// Validate inputs against struct tags
	body, err := validators.RegisterVehicleRequest(c)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "VALIDATEERR-1", "message": err.Error()})
		return
	}

	item, err := database.MarshallStruct(body)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "DYNAMOERR-1", "message": err.Error()})
		return
	}

	// Create a new DynamoDB client
	_, err = database.PutItem(item)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "DYNAMOERR-1", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &body)
}
