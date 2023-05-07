package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Capacity struct {
	Value int    `value:"value" binding:"required"`
	Unit  string `unit:"unit" binding:"required"`
}

type Vehicle struct {
	Vin          string   `vin:"vin" binding:"required"`
	Manufacturer string   `make:"make" binding:"required"`
	Model        string   `model:"manufacturer" binding:"required"`
	Year         int      `year:"year" binding:"required"`
	Color        string   `color:"color" binding:"required"`
	Capacity     Capacity `capacity:"capacity" binding:"required"`
	LicensePlate string   `license_plate:"license_plate"`
}

func RegisterVehicle(c *gin.Context) {
	body := Vehicle{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "VALIDATEERR-1", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &body)
}
