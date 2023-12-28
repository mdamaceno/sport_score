package views

import (
	"time"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
)

type Country struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func OneCountry(country models.Country) Country {
	return Country{
		Id:        country.Id,
		Name:      country.Name,
		Slug:      country.Slug,
		CreatedAt: country.CreatedAt,
		UpdatedAt: country.UpdatedAt,
	}
}

func ManyCountries(countries []models.Country) []Country {
	response := []Country{}

	for _, country := range countries {
		response = append(response, Country{
			Id:        country.Id,
			Name:      country.Name,
			Slug:      country.Slug,
			CreatedAt: country.CreatedAt,
			UpdatedAt: country.UpdatedAt,
		})
	}

	return response
}
