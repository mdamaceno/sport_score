package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgconn"
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := validate.Struct(countryParams); err != nil {
        mapErrors := helpers.MapValidationErrors(err)
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"errors": mapErrors})
	}

	uuid := uuid.New()

	country := models.Country{
		Id:   uuid,
		Name: countryParams.Name,
		Slug: slug.Make(countryParams.Name),
	}

	if err := cc.DB.Create(&country).Error; err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Country already exists"})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	log.Println("Country created successfully with id: " + uuid.String())

	c.JSON(http.StatusCreated, &country)

	return nil
}

func (cc CountriesController) Show(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	if err := cc.DB.Where("id = ?", id).First(&country).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Country not found"})
	}

	return c.JSON(http.StatusOK, &country)
}

func (cc CountriesController) Index(c echo.Context) error {
    countries := []models.Country{}

    cc.DB.Find(&countries)

    return c.JSON(http.StatusOK, &countries)
}
