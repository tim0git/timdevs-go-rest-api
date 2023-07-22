package modal_test

import (
	"eve.vehicle.api.com/m/v2/modal"
	"eve.vehicle.api.com/m/v2/vehicle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateVehicle(t *testing.T) {
	t.Parallel()

	mockVin := "GB000000000"

	mockVehicle := vehicle.Update{
		Manufacturer: "Jeep",
		Model:        "Grand Cherokee",
		Year:         2015,
		Color:        "Black",
		Capacity: vehicle.Capacity{
			Value: 14,
			Unit:  "kWh",
		},
	}

	_, err := modal.UpdateVehicle(mockVehicle, mockVin)
	assert.NoError(t, err)
}
