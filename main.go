package main

import (
	config "github.com/furqonzt99/refresh-token/configs"
	"github.com/furqonzt99/refresh-token/delivery/common"
	ac "github.com/furqonzt99/refresh-token/delivery/controllers/auth"
	uc "github.com/furqonzt99/refresh-token/delivery/controllers/users"
	"github.com/furqonzt99/refresh-token/delivery/middlewares"
	"github.com/furqonzt99/refresh-token/delivery/routes"
	ar "github.com/furqonzt99/refresh-token/repository/auth"
	ur "github.com/furqonzt99/refresh-token/repository/users"
	"github.com/furqonzt99/refresh-token/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()

	db := utils.InitDB(config)

	utils.InitialMigrate(db)

	e := echo.New()

	// logging
	middlewares.LogMiddleware(e)

	// remove trailing slash
	e.Pre(middleware.RemoveTrailingSlash())

	// validator
	e.Validator = &common.Validator{Validator: validator.New()}

	// repository
	authRepository := ar.NewAuthRepository(db)
	userRepository := ur.NewUserRepository(db)

	// controller
	authController := ac.NewAuthController(authRepository)
	userController := uc.NewUserController(userRepository)

	// routes
	routes.RegisterAuthPath(e, authController)
	routes.RegisterUserPath(e, userController)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
