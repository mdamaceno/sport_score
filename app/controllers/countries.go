package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/helpers"
	"github.com/mdmaceno/sport_score/app/models"
	"gorm.io/gorm"
)

type CountriesController struct {
	DB *gorm.DB
}

func (cc CountriesController) CreateCountry(c echo.Context) error {
	json_map, err := helpers.DecodeRawJson(c.Request().Body)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	uuid := uuid.New()
	name := json_map["name"].(string)

	country := models.Country{
		Id:   uuid,
		Name: name,
		Slug: slug.Make(name),
	}

	if err := cc.DB.Create(&country).Error; err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, map[string]string{"error": "Country already exists"})
		}

		return err
	}

	log.Println("Country created successfully with id: " + uuid.String())

	c.JSON(http.StatusCreated, &country)

	return nil
}

func (cc CountriesController) FindCountryById(c echo.Context) error {
	country := models.Country{}
	id := c.Param("id")

	if err := cc.DB.Where("id = ?", id).First(&country).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Country not found"})
	}

	return c.JSON(http.StatusOK, &country)
}
