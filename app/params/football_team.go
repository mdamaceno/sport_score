package params

type CreateFootballTeamParams struct {
	Name      string `json:"name" validate:"required,min=3,max=100"`
	CountryId string `json:"country_id" validate:"required,uuid"`
}

type UpdateFootballTeamParams struct {
	Name      string `json:"name" validate:"required,min=3,max=100,omitempty"`
	CountryId string `json:"country_id" validate:"uuid,omitempty"`
}
