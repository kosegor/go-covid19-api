package usecase

import (
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/domain/repository"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type InfectedUsecase interface {
	Post(infected model.Infected) (model.Infected, *apierr.ApiError)
	List() ([]*model.Infected, *apierr.ApiError)
}

type infectedUsecase struct {
	repo repository.InfectedRepository
}

func NewInfectedUsecase(repo repository.InfectedRepository) *infectedUsecase {
	return &infectedUsecase{
		repo: repo,
	}
}

func (i *infectedUsecase) Post(infected model.Infected) (model.Infected, *apierr.ApiError) {
	i.repo.Insert(&infected)
	return infected, nil
}

func (i *infectedUsecase) List() ([]*model.Infected, *apierr.ApiError) {
	return i.repo.FindAll()
}
