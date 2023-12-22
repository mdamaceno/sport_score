package config

import (
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/controllers"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, DB *gorm.DB) *echo.Echo {
	countriesControllers := controllers.CountriesController{DB: DB}
	apiV1 := e.Group("/api/v1")

    apiV1.GET("/countries", countriesControllers.Index)
	apiV1.GET("/countries/:id", countriesControllers.Show)
	apiV1.POST("/countries", countriesControllers.Create)
    apiV1.PATCH("/countries/:id", countriesControllers.Update)
    apiV1.DELETE("/countries/:id", countriesControllers.Delete)

	return e
}
