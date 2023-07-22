package error_test

import (
	"errors"
	"eve.vehicle.api.com/m/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidationErrorReturnsValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := errors.New("invalid input")
	error.ValidationError(c, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"VALIDATEERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"invalid input"`)
}
