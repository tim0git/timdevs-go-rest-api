package validators

import (
	"github.com/gin-gonic/gin"
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

func RegisterVehicleInput(c *gin.Context) (Vehicle, error) {
	body := Vehicle{}
	err := c.ShouldBindJSON(&body)

	if err != nil {
		return body, err
	}

	return body, nil
}
