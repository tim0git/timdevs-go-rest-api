package handler_update_vehicle

import (
	"eve.vehicle.api.com/m/v2/error"
	"eve.vehicle.api.com/m/v2/modal"
	"eve.vehicle.api.com/m/v2/vehicle"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateVehicle godoc
// @Summary update an existing vehicle
// @Schemes
// @Description update an existing vehicle
// @Tags vehicle
// @Accept json
// @Produce json
// @Param vin path string true "Vehicle identification number"
// @Param request body vehicle.Update true "Vehicle information"
// @Success 200
// @Router /vehicle/{vin} [patch]
func UpdateVehicle(c *gin.Context) {
	vin := c.Param("vin")

	vehicleUpdate := vehicle.Update{}

	validationError := c.ShouldBindJSON(&vehicleUpdate)
	if validationError != nil {
		error.ValidationError(c, validationError)
		return
	}

	_, updateItemError := modal.UpdateVehicle(vehicleUpdate, vin)
	if updateItemError != nil {
		error.DynamoDBError(c, updateItemError)
		return
	}

	c.Status(http.StatusOK)
}
