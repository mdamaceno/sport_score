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
	"github.com/mdmaceno/sport_score/app/response"
	"github.com/mdmaceno/sport_score/app/views"
	"gorm.io/gorm"
)

type FootballTeamsController struct {
	DB *gorm.DB
}

func (c FootballTeamsController) Create(ctx echo.Context) error {
	footballTeamParams := new(params.CreateFootballTeamParams)

	if err := ctx.Bind(footballTeamParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, nil))
	}

	if err := helpers.Validate.Struct(footballTeamParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, mapErrors))
	}

	country := models.Country{}

	if err := c.DB.Where("id = ?", footballTeamParams.CountryId).First(&country).Error; err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
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
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(footballTeam.Slug+" already exists", nil))
		}

		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.GENERIC_MESSAGE, nil))
	}

	log.Println("FootballTeam created successfully with id: " + footballTeam.Id.String())

	return ctx.JSON(http.StatusCreated, response.NewAPIResponse(views.OneFootballTeam(footballTeam)))
}

func (c FootballTeamsController) Index(ctx echo.Context) error {
	footballTeams := []models.FootballTeam{}

	c.DB.Find(&footballTeams)

	return ctx.JSON(http.StatusOK, response.NewAPIResponse(views.ManyFootballTeams(footballTeams)))
}

func (c FootballTeamsController) Show(ctx echo.Context) error {
	footballTeam := models.FootballTeam{}

	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&footballTeam).Error; err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.FOOTBALL_TEAM_NOT_FOUND, nil))
	}

	return ctx.JSON(http.StatusOK, response.NewAPIResponse(views.OneFootballTeam(footballTeam)))
}

func (controller FootballTeamsController) Update(ctx echo.Context) error {
	footballTeamParams := new(params.UpdateFootballTeamParams)

	if err := ctx.Bind(footballTeamParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, nil))
	}

	if err := helpers.Validate.Struct(footballTeamParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, mapErrors))
	}

	footballTeam := models.FootballTeam{}

	if err := controller.DB.Where("id = ?", ctx.Param("id")).First(&footballTeam).Error; err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.FOOTBALL_TEAM_NOT_FOUND, nil))
	}

	if footballTeamParams.CountryId != "" {
		countryId, err := uuid.Parse(footballTeamParams.CountryId)

		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
		}

		country := models.Country{}

		if err := controller.DB.Where("id = ?", footballTeamParams.CountryId).First(&country).Error; err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
		}

		footballTeam.CountryId = countryId
	}

	footballTeam.Name = footballTeamParams.Name
	footballTeam.Slug = slug.Make(footballTeamParams.Name)

	if err := controller.DB.Save(&footballTeam).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(footballTeam.Slug+" already exists", nil))
		}

		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.GENERIC_MESSAGE, nil))
	}

	log.Println("FootballTeam updated successfully with id: " + footballTeam.Id.String())

	return ctx.JSON(http.StatusAccepted, response.NewAPIResponse(views.OneFootballTeam(footballTeam)))
}

func (controller FootballTeamsController) Delete(ctx echo.Context) error {
	footballTeam := models.FootballTeam{}
	id := ctx.Param("id")

	controller.DB.Where("id = ?", id).First(&footballTeam).Delete(&footballTeam)

	return ctx.JSON(http.StatusNoContent, nil)
}
