package views

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
	"github.com/stretchr/testify/assert"
)

func TestCountryView(t *testing.T) {
	t.Run("OneCountry", func(t *testing.T) {
		t.Run("should return a country view", func(t *testing.T) {
			country := models.Country{
				Id:   uuid.New(),
				Name: "Brazil",
			}

			oneCountry := OneCountry(country)

			assert.Equal(t, oneCountry.Id, country.Id)
		})
	})

	t.Run("ManyCountries", func(t *testing.T) {
		t.Run("should return a country view", func(t *testing.T) {
			countries := []models.Country{
				{
					Id:   uuid.New(),
					Name: "Brazil",
				},
				{
					Id:   uuid.New(),
					Name: "Argentina",
				},
			}

			manyCountries := ManyCountries(countries)

			assert.Equal(t, manyCountries[0].Id, countries[0].Id)
			assert.Equal(t, manyCountries[1].Id, countries[1].Id)
		})
	})
}
