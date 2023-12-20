package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/models"
	"gorm.io/gorm"
)

type CountriesController struct {
	DB *gorm.DB
}

type CountryParams struct {
	Name string `json:"name"`
}

func (cc CountriesController) Create(c echo.Context) error {
	countryParams := new(CountryParams)

	err := c.Bind(countryParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
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
