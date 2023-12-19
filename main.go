package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app/controllers"
	config "github.com/mdmaceno/sport_score/config"
)

func main() {
	e := echo.New()
	DB := config.InitDB()
	countriesControllers := controllers.CountriesController{DB: DB}

	e.POST("/countries", countriesControllers.CreateCountry)

	e.Logger.Fatal(e.Start(":1323"))
}