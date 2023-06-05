package utils_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"timdevs.rest.api.com/m/v2/utils"
)

func TestGetTableName(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("TABLE_NAME", "MockTableName")

	response := utils.GetTableName()

	assert.Equal(t, "MockTableName", response)
}
