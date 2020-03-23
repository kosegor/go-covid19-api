package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Incident struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Surname            string  `json:"surname"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Country            string  `json:"country"`
	CountryOfResidence string  `json:"country_of_residence"`
	Date               string  `json:"date"`
}

func (i Incident) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.Surname, validation.Required),
		validation.Field(&i.Latitude, validation.Required, validation.Min(-90.0), validation.Max(90.0)),
		validation.Field(&i.Longitude, validation.Required, validation.Min(-180.0), validation.Max(180.0)),
		validation.Field(&i.Country, validation.Required),
		validation.Field(&i.CountryOfResidence, validation.Required),
		validation.Field(&i.Date, validation.Required, validation.Date(time.RFC3339)))
}
