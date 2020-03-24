package dynamo

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

const (
	IncidentTable = "Incident"
	IdsTable      = "Ids"
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
	id, _ := i.getId()
	id.ID++
	incident.ID = id.ID

	err := insert(i.service, incident)

	i.updateId(id)

	return err
}

func (i *incidentRepository) getId() (*model.Id, *apierr.ApiError) {
	id := &model.Id{}
	filt := expression.Name("table").Equal(expression.Value(IncidentTable))
	proj := expression.NamesList(expression.Name("table"), expression.Name("id"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(IdsTable),
	}

	result, err := i.service.Scan(params)
	if err != nil {
		return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], id)

	if err != nil {
		return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	return id, nil
}

func (i *incidentRepository) updateId(id *model.Id) *apierr.ApiError {
	av, err := dynamodbattribute.MarshalMap(id)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(IdsTable),
	}

	_, err = i.service.PutItem(input)
	if err != nil {
		logrus.Error(fmt.Errorf("Table %s: saving id failed. Error: %s", id.Table, err.Error()))
		return apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func insert(svc dynamodbiface.DynamoDBAPI, incident *model.Incident) *apierr.ApiError {
	av, err := dynamodbattribute.MarshalMap(incident)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(IncidentTable),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		logrus.Error(fmt.Errorf("ID %d: saving incident failed. Error: %s", incident.ID, err.Error()))
		return apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (i *incidentRepository) FindAll() ([]*model.Incident, *apierr.ApiError) {
	return findAll(i.service)
}

func findAll(svc dynamodbiface.DynamoDBAPI) ([]*model.Incident, *apierr.ApiError) {
	proj := expression.NamesList(expression.Name("id"), expression.Name("name"), expression.Name("surname"), expression.Name("latitude"), expression.Name("longitude"), expression.Name("country"), expression.Name("country_of_residence"), expression.Name("date"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(IncidentTable),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	incidents := make([]*model.Incident, len(result.Items))

	for x, i := range result.Items {
		incident := &model.Incident{}

		err = dynamodbattribute.UnmarshalMap(i, incident)

		if err != nil {
			return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
		}

		incidents[x] = incident
	}

	return incidents, nil
}
