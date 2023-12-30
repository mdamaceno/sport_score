package main

import (
	"github.com/labstack/echo/v4"
	config "github.com/mdmaceno/sport_score/config"
)

func main() {
	e := echo.New()
	appConfig := config.InitConfig()
	DB := config.InitDB(appConfig)
	routes := config.InitRoutes(e, DB)

	e.Logger.Fatal(routes.Start(":1323"))
}
