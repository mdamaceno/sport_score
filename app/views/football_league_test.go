package views

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
	"github.com/stretchr/testify/assert"
)

func TestFootballLeagueView(t *testing.T) {
	t.Run("OneFootballLeague", func(t *testing.T) {
		t.Run("should return a football league view", func(t *testing.T) {
			fl := models.FootballLeague{
				Id:        uuid.New(),
				Name:      "Campeonato Brasileiro",
				CountryId: uuid.New(),
			}

			oneFootballLeague := OneFootballLeague(fl)

			assert.Equal(t, oneFootballLeague.Id, fl.Id)
		})
	})

	t.Run("ManyFootballLeagues", func(t *testing.T) {
		t.Run("should return a football league view", func(t *testing.T) {
			fls := []models.FootballLeague{
				{
					Id:        uuid.New(),
					Name:      "Campeonato Brasileiro",
					CountryId: uuid.New(),
				},
				{
					Id:        uuid.New(),
					Name:      "Campeonato Argentino",
					CountryId: uuid.New(),
				},
			}

			manyFootballLeagues := ManyFootballLeagues(fls)

			assert.Equal(t, manyFootballLeagues[0].Id, fls[0].Id)
			assert.Equal(t, manyFootballLeagues[1].Id, fls[1].Id)
		})
	})
}
