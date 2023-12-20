package config

import (
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/controllers"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, DB *gorm.DB) *echo.Echo {
	countriesControllers := controllers.CountriesController{DB: DB}
	apiV1 := e.Group("/api/v1")

	apiV1.GET("/countries/:id", countriesControllers.Show)
	apiV1.POST("/countries", countriesControllers.Create)

	return e
}
