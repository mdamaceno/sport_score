package controllers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/mdmaceno/sport_score/app/models"
	"gorm.io/gorm"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

type CountriesController struct {
	DB *gorm.DB
}

type CountryParams struct {
	Name string `json:"name" validate:"required"`
}

func (cc CountriesController) Create(c echo.Context) error {
	countryParams := new(CountryParams)

	if err := c.Bind(countryParams); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := validate.Struct(countryParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Data:          mapErrors,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	uuid := uuid.New()

	country := models.Country{
		Id:   uuid,
		Name: countryParams.Name,
		Slug: slug.Make(countryParams.Name),
	}

	if err := cc.DB.Create(&country).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return c.JSON(http.StatusConflict, map[string]helpers.Error{
				"error": {
					OriginalError: err,
					Message:       country.Slug + " already exists",
					Name:          http.StatusText(http.StatusConflict),
				},
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Something went wrong",
				Name:          http.StatusText(http.StatusInternalServerError),
			},
		})
	}

	log.Println("Country created successfully with id: " + uuid.String())

	return c.JSON(http.StatusCreated, &country)
}

func (cc CountriesController) Show(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	if err := cc.DB.Where("id = ?", id).First(&country).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Country not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	return c.JSON(http.StatusOK, &country)
}

func (cc CountriesController) Index(c echo.Context) error {
	countries := []models.Country{}

	cc.DB.Find(&countries)

	return c.JSON(http.StatusOK, &countries)
}

func (cc CountriesController) Update(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	if err := cc.DB.Where("id = ?", id).First(&country).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Country not found",
				Name:          http.StatusText(http.StatusNotFound),
			},
		})
	}

	countryParams := new(CountryParams)

	if err := c.Bind(countryParams); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	if err := validate.Struct(countryParams); err != nil {
		mapErrors := helpers.MapValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Data:          mapErrors,
				Message:       "Invalid request body",
				Name:          http.StatusText(http.StatusUnprocessableEntity),
			},
		})
	}

	country.Name = countryParams.Name
	country.Slug = slug.Make(countryParams.Name)

	if err := cc.DB.Save(&country).Error; err != nil {
		if helpers.PGConflictError(err) != nil {
			return c.JSON(http.StatusConflict, map[string]helpers.Error{
				"error": {
					OriginalError: err,
					Message:       country.Slug + " already exists",
					Name:          http.StatusText(http.StatusConflict),
				},
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]helpers.Error{
			"error": {
				OriginalError: err,
				Message:       "Something went wrong",
				Name:          http.StatusText(http.StatusInternalServerError),
			},
		})
	}

	return c.JSON(http.StatusAccepted, &country)
}

func (cc CountriesController) Delete(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	cc.DB.Where("id = ?", id).First(&country).Delete(&country)

	return c.JSON(http.StatusNoContent, nil)
}
