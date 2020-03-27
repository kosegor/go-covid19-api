package repository

import (
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type DynamoRepository interface {
	Insert(*model.Incident) *apierr.ApiError
	FindAll() ([]*model.Incident, *apierr.ApiError)
}
