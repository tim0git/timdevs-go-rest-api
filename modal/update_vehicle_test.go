package modal_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"timdevs.rest.api.com/m/v2/modal"
	"timdevs.rest.api.com/m/v2/vehicle"
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
