package params

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/stretchr/testify/assert"
)

func TestFootballTeamParams(t *testing.T) {
	t.Run("CreateFootballTeamParams", func(t *testing.T) {
		var validCreateParams = CreateFootballTeamParams{
			Name:      "Flamengo",
			CountryId: uuid.New().String(),
		}

		t.Run("should not return an error", func(t *testing.T) {
			t.Run("when attributes are valid", func(t *testing.T) {
				err := helpers.Validate.Struct(validCreateParams)

				assert.Nil(t, err)
			})
		})

		t.Run("should return an error", func(t *testing.T) {
			t.Run("when name is empty", func(t *testing.T) {
				params := validCreateParams
				params.Name = ""

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "required")
			})

			t.Run("when name is too short", func(t *testing.T) {
				params := validCreateParams
				params.Name = "a"

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "min")
			})

			t.Run("when name is too long", func(t *testing.T) {
				params := validCreateParams
				params.Name = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor."

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "max")
			})

			t.Run("when country_id is empty", func(t *testing.T) {
				params := validCreateParams
				params.CountryId = ""

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "CountryId")
				assert.Equal(t, validationErrors[0].Tag(), "required")
			})
		})
	})

	t.Run("UpdateFootballTeamParams", func(t *testing.T) {
		var validUpdateParams = UpdateFootballTeamParams{
			Name:      "Flamengo",
			CountryId: uuid.New().String(),
		}

		t.Run("should not return an error", func(t *testing.T) {
			t.Run("when attributes are valid", func(t *testing.T) {
				err := helpers.Validate.Struct(validUpdateParams)

				assert.Nil(t, err)
			})
		})

		t.Run("should return an error", func(t *testing.T) {
			t.Run("when name is empty", func(t *testing.T) {
				params := validUpdateParams
				params.Name = ""

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "required")
			})

			t.Run("when name is too short", func(t *testing.T) {
				params := validUpdateParams
				params.Name = "a"

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "min")
			})

			t.Run("when name is too long", func(t *testing.T) {
				params := validUpdateParams
				params.Name = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor."

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "max")
			})

			t.Run("when country_id is empty", func(t *testing.T) {
				params := validUpdateParams
				params.CountryId = ""

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "CountryId")
				assert.Equal(t, validationErrors[0].Tag(), "uuid")
			})
		})
	})
}
