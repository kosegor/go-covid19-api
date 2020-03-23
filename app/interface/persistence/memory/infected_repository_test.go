package memory

import (
	"testing"
	"time"

	"github.com/kosegor/go-covid19-api/app/domain/model"
)

func TestIncidentRepositoryInsertAndFindAll(t *testing.T) {
	incident := model.Incident{
		Name:               "Andres",
		Surname:            "Gomez",
		Latitude:           -34.583863,
		Longitude:          -58.3452797,
		Country:            "Argentina",
		CountryOfResidence: "Argentina",
		Date:               time.Now().Format(time.RFC3339),
	}

	repo := NewIncidentRepository()
	repo.Insert(&incident)

	incidents, _ := repo.FindAll()

	if len(incidents) != 1 {
		t.Error("Failed: the incident didn't save")
	}

	if incidents[0].Name != incident.Name {
		t.Errorf("Failed: expected name [%s] but got [%s]", incident.Name, incidents[0].Name)
	}
}
