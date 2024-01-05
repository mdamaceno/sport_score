package controllers

import (
	"log"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/mdmaceno/sport_score/app/models"
	"github.com/mdmaceno/sport_score/app/params"
	"github.com/mdmaceno/sport_score/app/response"
	"github.com/mdmaceno/sport_score/app/views"
	"gorm.io/gorm"
)

type CountriesController struct {
	DB *gorm.DB
}

func (cc CountriesController) Create(c echo.Context) error {
	countryParams := new(params.CreateCountryParams)

	if err := c.Bind(countryParams); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, nil))
	}

	if err := helpers.Validate.Struct(countryParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, mapErrors))
	}

	country := models.Country{
		Name: countryParams.Name,
		Slug: slug.Make(countryParams.Name),
	}

	if err := cc.DB.Create(&country).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(country.Slug+" already exists", nil))
		}
	}

	log.Println("Country created successfully with id: " + country.Id.String())

	return c.JSON(http.StatusCreated, response.NewAPIResponse(views.OneCountry(country)))
}

func (cc CountriesController) Show(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	if err := cc.DB.Where("id = ?", id).First(&country).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
	}

	return c.JSON(http.StatusOK, response.NewAPIResponse(views.OneCountry(country)))
}

func (cc CountriesController) Index(c echo.Context) error {
	countries := []models.Country{}

	cc.DB.Find(&countries)

	return c.JSON(http.StatusOK, response.NewAPIResponse(views.ManyCountries(countries)))
}

func (cc CountriesController) Update(c echo.Context) error {
	countryParams := new(params.UpdateCountryParams)

	if err := c.Bind(countryParams); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, nil))
	}

	if err := helpers.Validate.Struct(countryParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.INVALID_REQUEST, mapErrors))
	}

	country := models.Country{}

	if err := cc.DB.Where("id = ?", c.Param("id")).First(&country).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.COUNTRY_NOT_FOUND, nil))
	}

	country.Name = countryParams.Name
	country.Slug = slug.Make(countryParams.Name)

	if err := cc.DB.Save(&country).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(country.Slug+" already exists", nil))
		}

		return c.JSON(http.StatusUnprocessableEntity, response.NewAPIErrorResponse(response.GENERIC_MESSAGE, nil))
	}

	return c.JSON(http.StatusAccepted, response.NewAPIResponse(views.OneCountry(country)))
}

func (cc CountriesController) Delete(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	cc.DB.Where("id = ?", id).First(&country).Delete(&country)

	return c.JSON(http.StatusNoContent, nil)
}
