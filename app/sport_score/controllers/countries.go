package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/sport_score/helpers"
	"github.com/mdmaceno/sport_score/app/sport_score/models"
	"gorm.io/gorm"
)

type CountriesController struct {
	DB *gorm.DB
}

func (cc CountriesController) CreateCountry(c echo.Context) error {
	json_map, err := helpers.DecodeRawJson(c)

	if err != nil {
		return err
	}

	country := models.Country{
		Id:   helpers.GenerateUUID(),
		Name: json_map["name"].(string),
	}

	if cc.DB.Create(&country).Error != nil {
		return err
	}

	c.JSON(http.StatusOK, &country)

	return nil
}
