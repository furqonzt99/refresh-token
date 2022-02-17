package main

import (
	config "github.com/furqonzt99/refresh-token/configs"
	"github.com/furqonzt99/refresh-token/delivery/middlewares"
	"github.com/furqonzt99/refresh-token/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()

	db := utils.InitDB(config)

	utils.InitialMigrate(db)

	e := echo.New()

	// logging
	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
