package memory

import (
	"testing"
	"time"

	"github.com/kosegor/go-covid19-api/app/domain/model"
)

func TestInfectedRepositoryInsertAndFindAll(t *testing.T) {
	infected := model.Infected{
		Name:               "Andres",
		Surname:            "Gomez",
		Latitude:           -34.583863,
		Longitude:          -58.3452797,
		Country:            "Argentina",
		CountryOfResidence: "Argentina",
		Date:               time.Now().Format(time.RFC3339),
	}

	repo := NewInfectedRepository()
	repo.Insert(&infected)

	infecteds, _ := repo.FindAll()

	if len(infecteds) != 1 {
		t.Error("Failed: the infected didn't save")
	}

	if infecteds[0].Name != infected.Name {
		t.Errorf("Failed: expected name [%s] but got [%s]", infected.Name, infecteds[0].Name)
	}
}
