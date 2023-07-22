package error

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DynamoDBError(c *gin.Context, err error) {
	const dynamoDBError = "DYNAMOERR-1"
	if awsError, ok := err.(awserr.Error); ok {
		switch awsError.Code() {
		case dynamodb.ErrCodeProvisionedThroughputExceededException:
			log.Println(dynamodb.ErrCodeProvisionedThroughputExceededException)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": dynamoDBError, "message": "ProvisionedThroughputExceededException"})
			break
		case dynamodb.ErrCodeResourceNotFoundException:
			log.Println(dynamodb.ErrCodeResourceNotFoundException)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": dynamoDBError, "message": "ResourceNotFoundException"})
			break
		case dynamodb.ErrCodeRequestLimitExceeded:
			log.Println(dynamodb.ErrCodeRequestLimitExceeded)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": dynamoDBError, "message": "RequestLimitExceeded"})
			break
		case dynamodb.ErrCodeInternalServerError:
			log.Println(dynamodb.ErrCodeInternalServerError)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dynamoDBError, "message": "InternalServerError"})
			break
		case dynamodb.ErrCodeConditionalCheckFailedException:
			log.Println(dynamodb.ErrCodeConditionalCheckFailedException)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dynamoDBError, "message": "ConditionalCheckFailedException"})
			break
		case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
			log.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException)
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{"error": dynamoDBError, "message": "ItemCollectionSizeLimitExceededException"})
			break
		case dynamodb.ErrCodeTransactionConflictException:
			log.Println(dynamodb.ErrCodeTransactionConflictException)
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": dynamoDBError, "message": "TransactionConflictException"})
			break
		default:
			log.Println(awsError.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dynamoDBError, "message": "InternalServerError"})
			break
		}
		return
	}

	log.Println(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dynamoDBError, "message": err.Error()})
}
