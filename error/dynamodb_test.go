package error_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/error"
)

func TestDynamoDBErrorReturnsDynamoDBError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := errors.New("database connection failed")
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"database connection failed"`)
}
