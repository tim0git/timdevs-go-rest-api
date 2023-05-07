package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func DynamoDbClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := dynamodb.New(sess, &aws.Config{Endpoint: aws.String(os.Getenv("DYNAMODB_ENDPOINT"))})

	if client == nil {
		panic("Unable to create DynamoDB client")
	}

	return client
}
