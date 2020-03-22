package memory

import (
	"sync"

	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type infectedRepository struct {
	mux       *sync.Mutex
	infecteds map[int]*model.Infected
}

func NewInfectedRepository() *infectedRepository {
	return &infectedRepository{
		mux:       &sync.Mutex{},
		infecteds: map[int]*model.Infected{},
	}
}

func (i *infectedRepository) Insert(infected *model.Infected) *apierr.ApiError {
	i.mux.Lock()
	defer i.mux.Unlock()

	if infected.ID == 0 {
		infected.ID = len(i.infecteds) + 1
	}
	i.infecteds[infected.ID] = infected

	return nil
}

func (i *infectedRepository) FindAll() ([]*model.Infected, *apierr.ApiError) {
	i.mux.Lock()
	defer i.mux.Unlock()

	infecteds := make([]*model.Infected, len(i.infecteds))
	for i, infected := range i.infecteds {
		infecteds[i-1] = infected
	}

	return infecteds, nil
}
