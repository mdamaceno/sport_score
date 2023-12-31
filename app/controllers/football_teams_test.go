package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/config"
	"github.com/stretchr/testify/assert"
)

var footballTeamJSON = struct {
	allValid         string
	noName           string
	noCountryId      string
	invalidCountryId string
}{
	allValid:         `{ "name": "Manchester United", "country_id": "ab005904-a7eb-11ee-ac0c-734d962dd9d1" }`,
	noName:           `{ "name": "", "country_id": "ab005904-a7eb-11ee-ac0c-734d962dd9d1" }`,
	noCountryId:      `{ "name": "Manchester United", "country_id": "" }`,
	invalidCountryId: `{ "name": "Manchester United", "country_id": "ab005904-a7eb-11ee-ac0c-734d962dd9d1" }`,
}

func TestFootballTeamsController(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Run("Should create a new football team", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/football_teams", strings.NewReader(footballTeamJSON.allValid))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &FootballTeamsController{DB: config.MockDB()}

			assert.NoError(t, c.Create(ctx))
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/football_teams", strings.NewReader(footballTeamJSON.noName))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				c := &FootballTeamsController{DB: config.MockDB()}

				c.Create(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})

			t.Run("When country_id is empty", func(t *testing.T) {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/football_teams", strings.NewReader(footballTeamJSON.noCountryId))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				c := &FootballTeamsController{DB: config.MockDB()}

				c.Create(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"countryid","reason":"required"}`)
			})
		})
	})

	t.Run("Index", func(t *testing.T) {
		t.Run("Should return all football teams", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/football_teams", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &FootballTeamsController{DB: config.MockDB()}

			assert.NoError(t, c.Index(ctx))
		})
	})

	t.Run("Show", func(t *testing.T) {
		t.Run("Should return a football team", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/football_teams/ab005904-a7eb-11ee-ac0c-734d962dd9d1", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &FootballTeamsController{DB: config.MockDB()}

			assert.NoError(t, c.Show(ctx))
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Should update a football team", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(
				http.MethodPatch,
				"/football_teams/ab005904-a7eb-11ee-ac0c-734d962dd9d1",
				strings.NewReader(footballTeamJSON.allValid),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &FootballTeamsController{DB: config.MockDB()}

			assert.NoError(t, c.Update(ctx))
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Should delete a football team", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/football_teams/ab005904-a7eb-11ee-ac0c-734d962dd9d1", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &FootballTeamsController{DB: config.MockDB()}

			assert.NoError(t, c.Delete(ctx))
		})
	})
}
