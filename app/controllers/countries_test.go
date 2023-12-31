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

var countryJSON = struct {
	allValid string
	noName   string
}{
	allValid: `{ "name": "Brazil" }`,
	noName:   `{ "name": "" }`,
}

func TestCountrysController(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Run("Should create a new country", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/countries", strings.NewReader(countryJSON.allValid))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &CountriesController{DB: config.MockDB()}

			assert.NoError(t, c.Create(ctx))
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/countries", strings.NewReader(countryJSON.noName))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				c := &CountriesController{DB: config.MockDB()}

				c.Create(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})
		})
	})

	t.Run("Index", func(t *testing.T) {
		t.Run("Should return all countries", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/countries", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &CountriesController{DB: config.MockDB()}

			assert.NoError(t, c.Index(ctx))
		})
	})

	t.Run("Show", func(t *testing.T) {
		t.Run("Should return a countries", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/countries/ab005904-a7eb-11ee-ac0c-734d962dd9d1", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &CountriesController{DB: config.MockDB()}

			assert.NoError(t, c.Show(ctx))
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Should update a countries", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(
				http.MethodPatch,
				"/countries/ab005904-a7eb-11ee-ac0c-734d962dd9d1",
				strings.NewReader(countryJSON.allValid),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &CountriesController{DB: config.MockDB()}

			assert.NoError(t, c.Update(ctx))
		})

		t.Run("Should return an error", func(t *testing.T) {
			t.Run("When name is empty", func(t *testing.T) {
				e := echo.New()
				req := httptest.NewRequest(
					http.MethodPatch,
					"/countries/ab005904-a7eb-11ee-ac0c-734d962dd9d1",
					strings.NewReader(countryJSON.noName),
				)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				c := &CountriesController{DB: config.MockDB()}

				c.Update(ctx)

				assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				assert.Contains(t, rec.Body.String(), `{"field":"name","reason":"required"}`)
			})
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Should delete a countries", func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/countries/ab005904-a7eb-11ee-ac0c-734d962dd9d1", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			c := &CountriesController{DB: config.MockDB()}

			assert.NoError(t, c.Delete(ctx))
		})
	})
}
