package params

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCountryParams(t *testing.T) {
	t.Run("CreateCountryParams", func(t *testing.T) {
		t.Run("should not return an error", func(t *testing.T) {
			t.Run("when attributes are valid", func(t *testing.T) {
				params := CreateCountryParams{
					Name: "Brazil",
				}

				err := helpers.Validate.Struct(params)

				assert.Nil(t, err)
			})
		})

		t.Run("should return an error", func(t *testing.T) {
			t.Run("when name is empty", func(t *testing.T) {
				params := CreateCountryParams{}

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "required")
			})

			t.Run("when name is too short", func(t *testing.T) {
				params := CreateCountryParams{
					Name: "a",
				}

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "min")
			})

			t.Run("when name is too long", func(t *testing.T) {
				params := CreateCountryParams{
					Name: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.",
				}

				err := helpers.Validate.Struct(params)
				validationErrors := err.(validator.ValidationErrors)

				assert.NotNil(t, err)
				assert.Equal(t, validationErrors[0].Field(), "Name")
				assert.Equal(t, validationErrors[0].Tag(), "max")
			})
		})
	})
}
