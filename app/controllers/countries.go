package controllers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
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

	country := models.Country{
		Id:   uuid,
		Name: json_map["name"].(string),
	}

	if cc.DB.Create(&country).Error != nil {
		log.Fatalln(err)
		return err
	}

	log.Println("Country created successfully with id: " + uuid.String())

	c.JSON(http.StatusOK, &country)

	return nil
}
