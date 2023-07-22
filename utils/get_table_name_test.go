package utils_test

import (
	"eve.vehicle.api.com/m/v2/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetTableName(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("TABLE_NAME", "MockTableName")

	response := utils.GetTableName()

	assert.Equal(t, "MockTableName", response)
}
