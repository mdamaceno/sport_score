package params

type CreateCountryParams struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCountryParams struct {
	Name string `json:"name"`
}
