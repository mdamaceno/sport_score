package views

import (
	"time"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
)

type FootballTeam struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Logo      string    `json:"logo"`
	CountryId uuid.UUID `json:"country_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func OneFootballTeam(ft models.FootballTeam) FootballTeam {
	return FootballTeam{
		Id:        ft.Id,
		Name:      ft.Name,
		Slug:      ft.Slug,
		Logo:      ft.Logo,
		CountryId: ft.CountryId,
		CreatedAt: ft.CreatedAt,
		UpdatedAt: ft.UpdatedAt,
	}
}

func ManyFootballTeams(fteams []models.FootballTeam) []FootballTeam {
	response := []FootballTeam{}

	for _, ft := range fteams {
		response = append(response, FootballTeam{
			Id:        ft.Id,
			Name:      ft.Name,
			Slug:      ft.Slug,
			Logo:      ft.Logo,
			CountryId: ft.CountryId,
			CreatedAt: ft.CreatedAt,
			UpdatedAt: ft.UpdatedAt,
		})
	}

	return response
}
