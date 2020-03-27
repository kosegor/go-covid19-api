package dynamo

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/kosegor/go-covid19-api/app/domain/model"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (c *mockDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoDBClient) Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return &dynamodb.ScanOutput{}, nil
}

func TestDynamoRepository_Insert(t *testing.T) {
	mockSvc := &mockDynamoDBClient{}
	incident := &model.Incident{
		Name:               "Andres",
		Surname:            "Gomez",
		Latitude:           -34.583863,
		Longitude:          -58.3452797,
		Country:            "Argentina",
		CountryOfResidence: "Argentina",
		Date:               time.Now().Format(time.RFC3339),
	}

	err := insert(mockSvc, incident)
	if err != nil && err.Error() != "" {
		t.Errorf("%s: ", err.Error())
	}
}

func TestDynamoRepository_FindAll(t *testing.T) {
	mockSvc := &mockDynamoDBClient{}

	_, err := findAll(mockSvc)
	if err != nil && err.Error() != "" {
		t.Errorf("%s: ", err.Error())
	}
}
