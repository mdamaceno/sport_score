package params

type CreateCountryParams struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateCountryParams struct {
	Name string `json:"name" validate:"required,min=3,max=100,omitempty"`
}
