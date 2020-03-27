package usecase

import (
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/domain/repository"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type IncidentUsecase interface {
	Post(incident model.Incident) (*model.Incident, *apierr.ApiError)
	List() ([]*model.Incident, *apierr.ApiError)
}

type incidentUsecase struct {
	dynamoRepo  repository.DynamoRepository
	elasticRepo repository.ElasticRepository
}

func NewIncidentUsecase(dynamoRepo repository.DynamoRepository, elasticRepo repository.ElasticRepository) *incidentUsecase {
	return &incidentUsecase{
		dynamoRepo:  dynamoRepo,
		elasticRepo: elasticRepo,
	}
}

func (i *incidentUsecase) Post(incident model.Incident) (*model.Incident, *apierr.ApiError) {
	inc, err := i.elasticRepo.Insert(&incident)
	if err != nil {
		return nil, err
	}

	err = i.dynamoRepo.Insert(inc)
	if err != nil {
		return nil, err
	}

	return inc, nil
}

func (i *incidentUsecase) List() ([]*model.Incident, *apierr.ApiError) {
	return i.dynamoRepo.FindAll()
}
