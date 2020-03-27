package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoService() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	sess := session.Must(session.NewSession(config))

	service := dynamodb.New(sess)

	return service
}
