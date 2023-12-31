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

var footballTeamJSON string = `{ "name": "Manchester United", "country_id": "ab005904-a7eb-11ee-ac0c-734d962dd9d1" }`

func TestFootballTeamsController(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Run("Should create a new football team", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/football_teams", strings.NewReader(footballTeamJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &FootballTeamsController{DB: config.MockDB()}

			assert.NoError(t, c.Create(ctx))
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
				strings.NewReader(footballTeamJSON),
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
