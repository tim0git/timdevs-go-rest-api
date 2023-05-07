package controllers

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"

	"timdevs.rest.api.com/m/v2/clients"
)

type Capacity struct {
	Value int    `json:"value" binding:"required"`
	Unit  string `json:"unit" binding:"required"`
}

type Vehicle struct {
	Vin          string   `json:"vin" binding:"required"`
	Manufacturer string   `json:"make" binding:"required"`
	Model        string   `json:"model" binding:"required"`
	Year         int      `json:"year" binding:"required"`
	Color        string   `json:"color" binding:"required"`
	Capacity     Capacity `json:"capacity" binding:"required"`
	LicensePlate string   `json:"license_plate"`
}

func RegisterVehicle(c *gin.Context) {
	body := Vehicle{}

	// Validate inputs against struct tags
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "VALIDATEERR-1", "message": err.Error()})
		return
	}

	// Marshal the struct into a map
	item, err := dynamodbattribute.MarshalMap(body)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "DYNAMOERR-1", "message": err.Error()})
		return
	}

	// Create a new DynamoDB client
	client := clients.DynamoDbClient()
	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item:      item,
	})
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "DYNAMOERR-1", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &body)
}
