package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"net/http"
	"timdevs.rest.api.com/m/v2/error"
	"timdevs.rest.api.com/m/v2/modal"
)

type VehicleCapacity struct {
	Value int    `json:"value" binding:"required"`
	Unit  string `json:"unit" binding:"required"`
}

type Vehicle struct {
	Vin          string          `json:"vin" binding:"required"`
	Manufacturer string          `json:"manufacturer" binding:"required"`
	Model        string          `json:"model" binding:"required"`
	Year         int             `json:"year" binding:"required"`
	Color        string          `json:"color" binding:"required"`
	Capacity     VehicleCapacity `json:"capacity" binding:"required"`
	LicensePlate string          `json:"license_plate"`
}

func RegisterVehicle(c *gin.Context) {
	body := Vehicle{}

	validationError := c.ShouldBindJSON(&body)
	if validationError != nil {
		error.ValidationError(c, validationError)
		return
	}

	item, marshalError := dynamodbattribute.MarshalMap(body)
	if marshalError != nil {
		error.DynamoDBError(c, marshalError)
		return
	}

	_, putItemError := modal.PutVehicle(item)
	if putItemError != nil {
		error.DynamoDBError(c, putItemError)
		return
	}

	c.JSON(http.StatusCreated, &body)
}
