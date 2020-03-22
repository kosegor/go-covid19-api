package repository

import (
	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type InfectedRepository interface {
	Insert(*model.Infected) *apierr.ApiError
	FindAll() ([]*model.Infected, *apierr.ApiError)
}
