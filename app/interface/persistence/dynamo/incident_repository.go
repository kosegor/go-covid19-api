package dynamo

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type incidentRepository struct {
	service *dynamodb.DynamoDB
}

func NewIncidentRepository() *incidentRepository {
	return &incidentRepository{
		service: NewDynamoService(),
	}
}

func (i *incidentRepository) Insert(incident *model.Incident) *apierr.ApiError {
	av, err := dynamodbattribute.MarshalMap(incident)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Incident"),
	}

	_, err = i.service.PutItem(input)
	if err != nil {
		logrus.Error(fmt.Errorf("ID %d: saving failed. Error: %s", incident.ID, err.Error()))
		return apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (i *incidentRepository) FindAll() ([]*model.Incident, *apierr.ApiError) {
	return nil, nil
}
