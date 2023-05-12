package error

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DynamoDBError(c *gin.Context, err error) {
	if awsError, ok := err.(awserr.Error); ok {
		switch awsError.Code() {
		case dynamodb.ErrCodeProvisionedThroughputExceededException:
			log.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, awsError.Error())
		case dynamodb.ErrCodeResourceNotFoundException:
			log.Println(dynamodb.ErrCodeResourceNotFoundException, awsError.Error())
		case dynamodb.ErrCodeRequestLimitExceeded:
			log.Println(dynamodb.ErrCodeRequestLimitExceeded, awsError.Error())
		case dynamodb.ErrCodeInternalServerError:
			log.Println(dynamodb.ErrCodeInternalServerError, awsError.Error())
		case dynamodb.ErrCodeConditionalCheckFailedException:
			log.Println(dynamodb.ErrCodeConditionalCheckFailedException, awsError.Error())
		case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
			log.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, awsError.Error())
		case dynamodb.ErrCodeTransactionConflictException:
			log.Println(dynamodb.ErrCodeTransactionConflictException, awsError.Error())
		default:
			log.Println(awsError.Error())
		}
	} else {
		log.Println(err.Error())
	}

	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"error": "DYNAMOERR-1", "message": err.Error()})
}
