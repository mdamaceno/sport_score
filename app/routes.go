package app

import (
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/controllers"
	"github.com/mdmaceno/sport_score/app/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, DB *gorm.DB) *echo.Echo {
	countriesControllers := controllers.CountriesController{DB: DB}
	footballLeaguesControllers := controllers.FootballLeaguesController{DB: DB}
	apiV1 := e.Group("/api/v1")

	apiV1.Use(middleware.LogIncomingAccess)

	apiV1.GET("/countries", countriesControllers.Index)
	apiV1.GET("/countries/:id", countriesControllers.Show)
	apiV1.POST("/countries", countriesControllers.Create)
	apiV1.PATCH("/countries/:id", countriesControllers.Update)
	apiV1.DELETE("/countries/:id", countriesControllers.Delete)

	apiV1.GET("/football_leagues", footballLeaguesControllers.Index)
	apiV1.GET("/football_leagues/:id", footballLeaguesControllers.Show)
	apiV1.POST("/football_leagues", footballLeaguesControllers.Create)
	apiV1.PATCH("/football_leagues/:id", footballLeaguesControllers.Update)
	apiV1.DELETE("/football_leagues/:id", footballLeaguesControllers.Delete)

	return e
}
