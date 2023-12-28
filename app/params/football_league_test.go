package params

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/stretchr/testify/assert"
)

var validParams = CreateFootballLeagueParams{
	Name:      "Campeonato Brasileiro",
	CountryId: uuid.New(),
}

func TestFootballLeagueParams(t *testing.T) {
	t.Run("CreateFootballLeagueParams", func(t *testing.T) {
		t.Run("should not return an error", func(t *testing.T) {
			t.Run("when attributes are valid", func(t *testing.T) {
				err := helpers.Validate.Struct(validParams)

				assert.Nil(t, err)
			})
		})

		t.Run("should return an error", func(t *testing.T) {
			t.Run("when name is empty", func(t *testing.T) {
				params := validParams
				params.Name = ""

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "required")
			})

			t.Run("when name is too short", func(t *testing.T) {
				params := validParams
				params.Name = "a"

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "min")
			})

			t.Run("when name is too long", func(t *testing.T) {
				params := validParams
				params.Name = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor."

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "max")
			})

			t.Run("when country_id is empty", func(t *testing.T) {
				params := validParams
				params.CountryId = uuid.UUID{}

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "CountryId")
				assert.Equal(t, validationErrors[0].Tag(), "required")
			})
		})
	})
}
