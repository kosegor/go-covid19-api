package model

import (
	"testing"
	"time"
)

func TestIncident_Validate(t *testing.T) {
	cases := []struct {
		name     string
		incident Incident
		errText  string
	}{
		{
			name: "FAIL/NAME_MISSING",
			incident: Incident{
				Surname:            "Gomez",
				Latitude:           -34.583863,
				Longitude:          -58.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "name: cannot be blank.",
		},
		{
			name: "FAIL/SURNAME_MISSING",
			incident: Incident{
				Name:               "Andres",
				Latitude:           -34.583863,
				Longitude:          -58.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "surname: cannot be blank.",
		},
		{
			name: "FAIL/LONGITUDE_MISSING",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Latitude:           -34.583863,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "longitude: cannot be blank.",
		},
		{
			name: "FAIL/LATITUDE_MISSING",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Longitude:          -58.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "latitude: cannot be blank.",
		},
		{
			name: "FAIL/COUNTRY_MISSING",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Latitude:           -34.583863,
				Longitude:          -58.3452797,
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "country: cannot be blank.",
		},
		{
			name: "FAIL/COUNTRY_OF_RESIDENCE_MISSING",
			incident: Incident{
				Name:      "Andres",
				Surname:   "Gomez",
				Latitude:  -34.583863,
				Longitude: -58.3452797,
				Country:   "Argentina",
				Date:      time.Now().Format(time.RFC3339),
			},
			errText: "country_of_residence: cannot be blank.",
		},
		{
			name: "FAIL/DATE_MISSING",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Latitude:           -34.583863,
				Longitude:          -58.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
			},
			errText: "date: cannot be blank.",
		},
		{
			name: "FAIL/INVALID_LATITUDE",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Latitude:           -134.583863,
				Longitude:          -58.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "latitude: must be no less than -90.",
		},
		{
			name: "FAIL/INVALID_LONGITUDE",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Latitude:           -34.583863,
				Longitude:          -258.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "longitude: must be no less than -180.",
		},
		{
			name: "OK",
			incident: Incident{
				Name:               "Andres",
				Surname:            "Gomez",
				Latitude:           -34.583863,
				Longitude:          -58.3452797,
				Country:            "Argentina",
				CountryOfResidence: "Argentina",
				Date:               time.Now().Format(time.RFC3339),
			},
			errText: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expectedError := c.errText
			err := c.incident.Validate()
			if expectedError != "" && expectedError != err.Error() {
				t.Errorf("%s: expected error message [%s] but got [%s]", c.name, expectedError, err.Error())
			}
		})
	}
}
