package views

import (
	"time"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
)

type FootballLeague struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CountryId uuid.UUID `json:"country_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func OneFootballLeague(fl models.FootballLeague) FootballLeague {
	return FootballLeague{
		Id:        fl.Id,
		Name:      fl.Name,
		Slug:      fl.Slug,
		CountryId: fl.CountryId,
		CreatedAt: fl.CreatedAt,
		UpdatedAt: fl.UpdatedAt,
	}
}

func ManyFootballLeagues(fleagues []models.FootballLeague) []FootballLeague {
	response := []FootballLeague{}

	for _, fl := range fleagues {
		response = append(response, FootballLeague{
			Id:        fl.Id,
			Name:      fl.Name,
			Slug:      fl.Slug,
			CountryId: fl.CountryId,
			CreatedAt: fl.CreatedAt,
			UpdatedAt: fl.UpdatedAt,
		})
	}

	return response
}
