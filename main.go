package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mdmaceno/sport_score/app"
	"github.com/mdmaceno/sport_score/config"
)

func main() {
	e := echo.New()
	env := config.Envs()
	DB := config.InitDB(env)
	routes := app.InitRoutes(e, DB)

	e.Logger.Fatal(routes.Start(":1323"))
}
