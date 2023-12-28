package params

import "github.com/google/uuid"

type CreateFootballLeagueParams struct {
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	CountryId uuid.UUID `json:"country_id" validate:"required,uuid"`
}

type UpdateFootballLeagueParams struct {
	Name      string `json:"name"`
	CountryId string `json:"country_id" validate:"uuid"`
}
