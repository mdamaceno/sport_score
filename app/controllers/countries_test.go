package controllers

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
	"github.com/mdmaceno/sport_score/config"
	"github.com/mdmaceno/sport_score/tests"
	"github.com/stretchr/testify/assert"
)

func TestCountriesController(t *testing.T) {
	DB := *tests.InitDB(config.Envs())
	setupOpts := &tests.SetupOptions{DB: &DB}

	country := models.Country{
		Id:   uuid.New(),
		Name: "Brazil",
	}

	t.Run("Create", func(t *testing.T) {
		t.Run("Should create a new country", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			body := `{ "name": "Brazil" }`

			ctx, rec := tests.PrepareRequest(http.MethodPost, body)

			ctx.SetPath("/countries")
			c := CountriesController{DB: &DB}
			c.Create(ctx)

			assert.Equal(t, http.StatusCreated, rec.Code)
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				teardown := tests.SetupTest(t, setupOpts)
				defer teardown(t)

				body := `{ "name": "" }`

				ctx, rec := tests.PrepareRequest(http.MethodPost, body)

				ctx.SetPath("/countries")
				c := CountriesController{DB: &DB}
				c.Create(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})
		})
	})

	t.Run("Index", func(t *testing.T) {
		t.Run("Should return all countries", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			ctx, rec := tests.PrepareRequest(http.MethodGet, "")

			ctx.SetPath("/countries")
			c := CountriesController{DB: &DB}
			c.Index(ctx)

			assert.Equal(t, http.StatusOK, rec.Code)
		})
	})

	t.Run("Show", func(t *testing.T) {
		t.Run("Should return a country", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)

			ctx, rec := tests.PrepareRequest(http.MethodGet, "")

			ctx.SetPath("/countries/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(country.Id.String())

			c := CountriesController{DB: &DB}
			c.Show(ctx)

			assert.Equal(t, http.StatusOK, rec.Code)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Should update a country", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)

			body := `{ "name": "Argentina" }`

			ctx, rec := tests.PrepareRequest(http.MethodPatch, body)

			ctx.SetPath("/countries/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(country.Id.String())

			c := CountriesController{DB: &DB}
			c.Update(ctx)

			assert.Equal(t, http.StatusAccepted, rec.Code)
			assert.Contains(t, rec.Body.String(), `"name":"Argentina"`)
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				body := `{ "name": "" }`

				ctx, rec := tests.PrepareRequest(http.MethodPatch, body)

				ctx.SetPath("/countries/:id")
				ctx.SetParamNames("id")
				ctx.SetParamValues(country.Id.String())

				c := CountriesController{DB: &DB}
				c.Update(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Should delete a countries", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)

			ctx, rec := tests.PrepareRequest(http.MethodDelete, "")

			ctx.SetPath("/countries/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(country.Id.String())
			c := CountriesController{DB: &DB}
			c.Delete(ctx)

			assert.Equal(t, http.StatusNoContent, rec.Code)
		})
	})
}
