package memory

import (
	"sync"

	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
)

type incidentRepository struct {
	mux       *sync.Mutex
	incidents map[int]*model.Incident
}

func NewIncidentRepository() *incidentRepository {
	return &incidentRepository{
		mux:       &sync.Mutex{},
		incidents: map[int]*model.Incident{},
	}
}

func (i *incidentRepository) Insert(incident *model.Incident) *apierr.ApiError {
	i.mux.Lock()
	defer i.mux.Unlock()

	if incident.ID == 0 {
		incident.ID = len(i.incidents) + 1
	}
	i.incidents[incident.ID] = incident

	return nil
}

func (i *incidentRepository) FindAll() ([]*model.Incident, *apierr.ApiError) {
	i.mux.Lock()
	defer i.mux.Unlock()

	incidents := make([]*model.Incident, len(i.incidents))
	for i, incident := range i.incidents {
		incidents[i-1] = incident
	}

	return incidents, nil
}
