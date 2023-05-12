package error_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/error"
)

func TestNotFoundError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	error.NotFoundError(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"NOTFOUNDERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"Vehicle not found"`)
}
