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

func TestFootballLeaguesController(t *testing.T) {
	DB := *tests.InitDB(config.Envs())
	setupOpts := &tests.SetupOptions{DB: &DB}

	country := models.Country{
		Id:   uuid.New(),
		Name: "Brazil",
	}

	footballLeague := models.FootballLeague{
		Id:        uuid.New(),
		Name:      "Brasileirão",
		CountryId: country.Id,
	}

	t.Run("Create", func(t *testing.T) {
		t.Run("Should create a new football league", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)

			body := `{ "name": "Flamengo", "country_id": "` + country.Id.String() + `" }`

			ctx, rec := tests.PrepareRequest(http.MethodPost, body)

			ctx.SetPath("/football_leagues")
			c := &FootballLeaguesController{DB: &DB}
			c.Create(ctx)

			assert.Equal(t, http.StatusCreated, rec.Code)
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				teardown := tests.SetupTest(t, setupOpts)
				defer teardown(t)

				DB.Create(&country)

				body := `{ "name": "", "country_id": "` + country.Id.String() + `" }`

				ctx, rec := tests.PrepareRequest(http.MethodPost, body)

				ctx.SetPath("/football_leagues")
				c := &FootballLeaguesController{DB: &DB}

				c.Create(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})

			t.Run("When country_id is empty", func(t *testing.T) {
				teardown := tests.SetupTest(t, setupOpts)
				defer teardown(t)

				DB.Create(&country)

				body := `{ "name": "Brasileirão Série A", "country_id": "" }`

				ctx, rec := tests.PrepareRequest(http.MethodPost, body)

				ctx.SetPath("/football_leagues")
				c := &FootballLeaguesController{DB: &DB}

				c.Create(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"countryid","reason":"required"}`)
			})
		})
	})

	t.Run("Index", func(t *testing.T) {
		t.Run("Should return all football leagues", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)
			DB.Create(&footballLeague)

			ctx, rec := tests.PrepareRequest(http.MethodGet, "")

			ctx.SetPath("/football_leagues")
			c := &FootballLeaguesController{DB: &DB}

			assert.NoError(t, c.Index(ctx))
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	})

	t.Run("Show", func(t *testing.T) {
		t.Run("Should return a football league", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)
			DB.Create(&footballLeague)

			ctx, rec := tests.PrepareRequest(http.MethodGet, "")

			ctx.SetPath("/football_leagues/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(footballLeague.Id.String())
			c := &FootballLeaguesController{DB: &DB}
			c.Show(ctx)

			assert.Equal(t, http.StatusOK, rec.Code)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Should update a football league", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)
			DB.Create(&footballLeague)

			body := `{ "name": "Brasileirão Série A", "country_id": "` + country.Id.String() + `" }`

			ctx, rec := tests.PrepareRequest(http.MethodPatch, body)

			ctx.SetPath("/football_leagues/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(footballLeague.Id.String())
			c := &FootballLeaguesController{DB: &DB}
			c.Update(ctx)

			assert.Equal(t, http.StatusAccepted, rec.Code)
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				teardown := tests.SetupTest(t, setupOpts)
				defer teardown(t)

				DB.Create(&country)
				DB.Create(&footballLeague)

				body := `{ "name": "", "country_id": "` + country.Id.String() + `" }`

				ctx, rec := tests.PrepareRequest(http.MethodPatch, body)

				ctx.SetPath("/football_leagues/:id")
				ctx.SetParamNames("id")
				c := &FootballLeaguesController{DB: &DB}

				c.Update(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Should delete a football league", func(t *testing.T) {
			teardown := tests.SetupTest(t, setupOpts)
			defer teardown(t)

			DB.Create(&country)
			DB.Create(&footballLeague)

			ctx, rec := tests.PrepareRequest(http.MethodDelete, "")

			ctx.SetPath("/football_leagues/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(footballLeague.Id.String())
			c := &FootballLeaguesController{DB: &DB}
			c.Delete(ctx)

			assert.Equal(t, http.StatusNoContent, rec.Code)
		})
	})
}
