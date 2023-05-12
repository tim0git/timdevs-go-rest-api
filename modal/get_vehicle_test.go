package modal_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"timdevs.rest.api.com/m/v2/modal"
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
