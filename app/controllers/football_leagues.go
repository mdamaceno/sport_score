package controllers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/mdmaceno/sport_score/app/models"
	"github.com/mdmaceno/sport_score/app/params"
	"github.com/mdmaceno/sport_score/app/views"
	"gorm.io/gorm"
)

type FootballLeaguesController struct {
	DB *gorm.DB
}

func (controller FootballLeaguesController) Create(ctx echo.Context) error {
	footballLeagueParams := new(params.CreateFootballLeagueParams)

	if err := ctx.Bind(footballLeagueParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := helpers.Validate.Struct(footballLeagueParams); err != nil {
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

	return ctx.JSON(http.StatusCreated, views.OneFootballLeague(footballLeague))
}

func (controller FootballLeaguesController) Index(ctx echo.Context) error {
	footballLeagues := []models.FootballLeague{}

	controller.DB.Find(&footballLeagues)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse{
		Data: views.ManyFootballLeagues(footballLeagues),
	})
}

func (controller FootballLeaguesController) Show(ctx echo.Context) error {
	footballLeague := models.FootballLeague{}

	if err := controller.DB.Where("id = ?", ctx.Param("id")).First(&footballLeague).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "FootballLeague not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	return ctx.JSON(http.StatusOK, views.OneFootballLeague(footballLeague))
}

func (controller FootballLeaguesController) Update(ctx echo.Context) error {
	footballLeagueParams := new(params.UpdateFootballLeagueParams)

	if err := ctx.Bind(footballLeagueParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := helpers.Validate.Struct(footballLeagueParams); err != nil {
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

	footballLeague := models.FootballLeague{}

	if err := controller.DB.Where("id = ?", ctx.Param("id")).First(&footballLeague).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Football League not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	if footballLeagueParams.CountryId != "" {
		countryId, err := uuid.Parse(footballLeagueParams.CountryId)

		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
				"error": {
					OriginalError: err,
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

		footballLeague.CountryId = countryId
	}

	footballLeague.Name = footballLeagueParams.Name
	footballLeague.Slug = slug.Make(footballLeagueParams.Name)

	if err := controller.DB.Save(&footballLeague).Error; err != nil {
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

	log.Println("FootballLeague updated successfully with id: " + footballLeague.Id.String())

	return ctx.JSON(http.StatusOK, views.OneFootballLeague(footballLeague))
}

func (controller FootballLeaguesController) Delete(ctx echo.Context) error {
	footballLeague := models.FootballLeague{}
	id := ctx.Param("id")

	controller.DB.Where("id = ?", id).First(&footballLeague).Delete(&footballLeague)

	return ctx.JSON(http.StatusNoContent, nil)
}
