package routes

import (
	"github.com/furqonzt99/refresh-token/constants"
	"github.com/furqonzt99/refresh-token/delivery/controllers/users"
	"github.com/furqonzt99/refresh-token/delivery/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserPath(e *echo.Echo, userController *users.UserController) {
	e.POST("/users", userController.Create, middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.AdminRole)
	e.GET("/users", userController.ReadAll, middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.AdminRole)
	e.GET("/users/:id", userController.ReadOne, middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.UserRole)
	e.PUT("/users/:id", userController.Update, middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.AdminRole)
	e.DELETE("/users/:id", userController.Delete, middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.AdminRole)
}
