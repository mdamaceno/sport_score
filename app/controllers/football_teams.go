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

type FootballTeamsController struct {
	DB *gorm.DB
}

func (c FootballTeamsController) Create(ctx echo.Context) error {
	footballTeamParams := new(params.CreateFootballTeamParams)

	if err := ctx.Bind(footballTeamParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := helpers.Validate.Struct(footballTeamParams); err != nil {
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

	if err := c.DB.Where("id = ?", footballTeamParams.CountryId).First(&country).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Country not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	countryId, _ := uuid.Parse(footballTeamParams.CountryId)

	footballTeam := models.FootballTeam{
		Id:        uuid.New(),
		Name:      footballTeamParams.Name,
		Slug:      slug.Make(footballTeamParams.Name),
		CountryId: countryId,
	}

	if err := c.DB.Create(&footballTeam).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return ctx.JSON(http.StatusConflict, map[string]helpers.Error{
				"error": {
					OriginalError: err,
					Message:       footballTeam.Slug + " already exists",
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

	log.Println("FootballTeam created successfully with id: " + footballTeam.Id.String())

	return ctx.JSON(http.StatusCreated, views.OneFootballTeam(footballTeam))
}

func (c FootballTeamsController) Index(ctx echo.Context) error {
	footballTeams := []models.FootballTeam{}

	c.DB.Find(&footballTeams)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse{
		Data: views.ManyFootballTeams(footballTeams),
	})
}

func (c FootballTeamsController) Show(ctx echo.Context) error {
	footballTeam := models.FootballTeam{}

	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&footballTeam).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "FootballTeam not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	return ctx.JSON(http.StatusOK, views.OneFootballTeam(footballTeam))
}

func (controller FootballTeamsController) Update(ctx echo.Context) error {
	footballTeamParams := new(params.UpdateFootballTeamParams)

	if err := ctx.Bind(footballTeamParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := helpers.Validate.Struct(footballTeamParams); err != nil {
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

	footballTeam := models.FootballTeam{}

	if err := controller.DB.Where("id = ?", ctx.Param("id")).First(&footballTeam).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Football Team not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	if footballTeamParams.CountryId != "" {
		countryId, err := uuid.Parse(footballTeamParams.CountryId)

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

		if err := controller.DB.Where("id = ?", footballTeamParams.CountryId).First(&country).Error; err != nil {
			return ctx.JSON(http.StatusNotFound, map[string]helpers.Error{
				"error": {
					OriginalError: err,
					Message:       "Country not found",
					Name:          http.StatusText(http.StatusNotFound),
				},
			})
		}

		footballTeam.CountryId = countryId
	}

	footballTeam.Name = footballTeamParams.Name
	footballTeam.Slug = slug.Make(footballTeamParams.Name)

	if err := controller.DB.Save(&footballTeam).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return ctx.JSON(http.StatusConflict, map[string]helpers.Error{
				"error": {
					OriginalError: err,
					Message:       footballTeam.Slug + " already exists",
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

	log.Println("FootballTeam updated successfully with id: " + footballTeam.Id.String())

	return ctx.JSON(http.StatusOK, views.OneFootballTeam(footballTeam))
}

func (controller FootballTeamsController) Delete(ctx echo.Context) error {
	footballTeam := models.FootballTeam{}
	id := ctx.Param("id")

	controller.DB.Where("id = ?", id).First(&footballTeam).Delete(&footballTeam)

	return ctx.JSON(http.StatusNoContent, nil)
}
