package repository

import (
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type ElasticRepository interface {
	Insert(*model.Incident) (*model.Incident, *apierr.ApiError)
	//FindByCountry(string) ([]*model.Incident, *apierr.ApiError)
}
