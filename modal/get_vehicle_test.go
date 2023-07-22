package modal_test

import (
	"eve.vehicle.api.com/m/v2/modal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetsVehicleWithoutError(t *testing.T) {
	t.Parallel()
	_, err := modal.GetVehicle("GB000000000")
	assert.NoError(t, err)
}
func TestGetsVehicleWithCorrectVin(t *testing.T) {
	t.Parallel()
	res, _ := modal.GetVehicle("GB000000000")
	assert.Equal(t, "GB000000000", *res.Item["vin"].S)
}
