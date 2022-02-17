package routes

import (
	"github.com/furqonzt99/refresh-token/delivery/controllers/users"
	"github.com/labstack/echo/v4"
)

func RegisterUserPath(e *echo.Echo, userController *users.UserController) {
	e.POST("/users", userController.Create)
	e.GET("/users", userController.ReadAll)
	e.GET("/users/:id", userController.ReadOne)
	e.PUT("/users/:id", userController.Update)
	e.DELETE("/users/:id", userController.Delete)
}
