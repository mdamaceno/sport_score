package views

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mdmaceno/sport_score/app/models"
	"github.com/stretchr/testify/assert"
)

func TestFootballTeamView(t *testing.T) {
	t.Run("OneFootballTeam", func(t *testing.T) {
		t.Run("should return a football team view", func(t *testing.T) {
			fl := models.FootballTeam{
				Id:        uuid.New(),
				Name:      "Campeonato Brasileiro",
				CountryId: uuid.New(),
			}

			oneFootballTeam := OneFootballTeam(fl)

			assert.Equal(t, oneFootballTeam.Id, fl.Id)
		})
	})

	t.Run("ManyFootballTeams", func(t *testing.T) {
		t.Run("should return a football team view", func(t *testing.T) {
			fls := []models.FootballTeam{
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

			manyFootballTeams := ManyFootballTeams(fls)

			assert.Equal(t, manyFootballTeams[0].Id, fls[0].Id)
			assert.Equal(t, manyFootballTeams[1].Id, fls[1].Id)
		})
	})
}
