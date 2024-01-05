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

type FootballLeaguesController struct {
	DB *gorm.DB
}

func (c FootballLeaguesController) Create(ctx echo.Context) error {
	footballLeagueParams := new(params.CreateFootballLeagueParams)

	if err := ctx.Bind(footballLeagueParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, nil))
	}

	if err := helpers.Validate.Struct(footballLeagueParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, mapErrors))
	}

	country := models.Country{}

	if err := c.DB.Where("id = ?", footballLeagueParams.CountryId).First(&country).Error; err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
	}

	countryId, _ := uuid.Parse(footballLeagueParams.CountryId)

	footballLeague := models.FootballLeague{
		Id:        uuid.New(),
		Name:      footballLeagueParams.Name,
		Slug:      slug.Make(footballLeagueParams.Name),
		CountryId: countryId,
	}

	if err := c.DB.Create(&footballLeague).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(footballLeague.Slug+" already exists", nil))
		}

		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.GENERIC_MESSAGE, nil))
	}

	log.Println("FootballLeague created successfully with id: " + footballLeague.Id.String())

	return ctx.JSON(http.StatusCreated, response.NewAPIResponse(views.OneFootballLeague(footballLeague)))
}

func (c FootballLeaguesController) Index(ctx echo.Context) error {
	footballLeagues := []models.FootballLeague{}

	c.DB.Find(&footballLeagues)

	return ctx.JSON(http.StatusOK, response.NewAPIResponse(views.ManyFootballLeagues(footballLeagues)))
}

func (c FootballLeaguesController) Show(ctx echo.Context) error {
	footballLeague := models.FootballLeague{}

	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&footballLeague).Error; err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
	}

	return ctx.JSON(http.StatusOK, response.NewAPIResponse(views.OneFootballLeague(footballLeague)))
}

func (c FootballLeaguesController) Update(ctx echo.Context) error {
	footballLeagueParams := new(params.UpdateFootballLeagueParams)

	if err := ctx.Bind(footballLeagueParams); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, nil))
	}

	if err := helpers.Validate.Struct(footballLeagueParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, mapErrors))
	}

	footballLeague := models.FootballLeague{}

	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&footballLeague).Error; err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.FOOTBALL_LEAGUE_NOT_FOUND, nil))
	}

	if footballLeagueParams.CountryId != "" {
		countryId, _ := uuid.Parse(footballLeagueParams.CountryId)

		country := models.Country{}

		if err := c.DB.Where("id = ?", footballLeagueParams.CountryId).First(&country).Error; err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
		}

		footballLeague.CountryId = countryId
	}

	footballLeague.Name = footballLeagueParams.Name
	footballLeague.Slug = slug.Make(footballLeagueParams.Name)

	if err := c.DB.Save(&footballLeague).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(footballLeague.Slug+" already exists", nil))
		}

		return ctx.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.GENERIC_MESSAGE, nil))
	}

	log.Println("FootballLeague updated successfully with id: " + footballLeague.Id.String())

	return ctx.JSON(http.StatusAccepted, response.NewAPIResponse(views.OneFootballLeague(footballLeague)))
}

func (c FootballLeaguesController) Delete(ctx echo.Context) error {
	footballLeague := models.FootballLeague{}
	id := ctx.Param("id")

	c.DB.Where("id = ?", id).First(&footballLeague).Delete(&footballLeague)

	return ctx.JSON(http.StatusNoContent, nil)
}
