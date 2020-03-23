package usecase

import (
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/domain/repository"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type IncidentUsecase interface {
	Post(incident model.Incident) (model.Incident, *apierr.ApiError)
	List() ([]*model.Incident, *apierr.ApiError)
}

type incidentUsecase struct {
	repo repository.IncidentRepository
}

func NewIncidentUsecase(repo repository.IncidentRepository) *incidentUsecase {
	return &incidentUsecase{
		repo: repo,
	}
}

func (i *incidentUsecase) Post(incident model.Incident) (model.Incident, *apierr.ApiError) {
	i.repo.Insert(&incident)
	return incident, nil
}

func (i *incidentUsecase) List() ([]*model.Incident, *apierr.ApiError) {
	return i.repo.FindAll()
}
