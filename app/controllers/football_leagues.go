package controllers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/mdmaceno/sport_score/app/models"
	"gorm.io/gorm"
)

type FootballLeaguesController struct {
	DB *gorm.DB
}

type FootballLeagueParams struct {
	Name      string    `json:"name" validate:"required"`
	CountryId uuid.UUID `json:"country_id" validate:"required,uuid"`
}

func (controller FootballLeaguesController) Create(ctx echo.Context) error {
	footballLeagueParams := new(FootballLeagueParams)

	if err := ctx.Bind(footballLeagueParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := validate.Struct(footballLeagueParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Data:          mapErrors,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	country := models.Country{}

	if err := controller.DB.Where("id = ?", footballLeagueParams.CountryId).First(&country).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Country not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	footballLeague := models.FootballLeague{
		Name:      footballLeagueParams.Name,
		Slug:      slug.Make(footballLeagueParams.Name),
		CountryId: footballLeagueParams.CountryId,
	}

	if err := controller.DB.Create(&footballLeague).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return ctx.JSON(http.StatusConflict, map[string]helpers.Error{
				"error": {
					OriginalError: err,
					Message:       footballLeague.Slug + " already exists",
					Name:          http.StatusText(http.StatusConflict),
				},
			})
		}

		return ctx.JSON(http.StatusInternalServerError, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Something went wrong",
				Name:          http.StatusText(http.StatusInternalServerError),
			},
		})
	}

	log.Println("FootballLeague created successfully with id: " + footballLeague.Id.String())

	return ctx.JSON(http.StatusCreated, &footballLeague)
}
