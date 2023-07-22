package error_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/error"
)

func TestReturnsStatusCode500AsDefaultResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := errors.New("database connection failed")
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"database connection failed"`)
}
func TestReturnsProvisionedThroughputExceededExceptionError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeProvisionedThroughputExceededException, "ProvisionedThroughputExceededException", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"ProvisionedThroughputExceededException"`)
}
func TestReturnsResourceNotFoundExceptionError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeResourceNotFoundException, "ResourceNotFoundException", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"ResourceNotFoundException"`)
}
func TestReturnsRequestLimitExceededError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeRequestLimitExceeded, "RequestLimitExceeded", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"RequestLimitExceeded"`)
}
func TestReturnsInternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeInternalServerError, "InternalServerError", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"InternalServerError"`)
}
func TestReturnsConditionalCheckFailedExceptionError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeConditionalCheckFailedException, "ConditionalCheckFailedException", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"ConditionalCheckFailedException"`)
}
func TestReturnsItemCollectionSizeLimitExceededExceptionError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, "ItemCollectionSizeLimitExceededException", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusRequestEntityTooLarge, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"ItemCollectionSizeLimitExceededException"`)
}
func TestReturnsTransactionConflictExceptionError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeTransactionConflictException, "TransactionConflictException", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"TransactionConflictException"`)
}
func TestReturnsInternalServerErrorForUnknownError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := awserr.New(dynamodb.ErrCodeReplicaNotFoundException, "", nil)
	error.DynamoDBError(c, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"DYNAMOERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"InternalServerError"`)
}
