package routes

import (
	"github.com/furqonzt99/refresh-token/delivery/controllers/auth"
	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, authController *auth.AuthController) {
	e.POST("/auth/register", authController.Register)
	e.POST("/auth/login", authController.Login)
	e.POST("/auth/refresh", authController.Refresh)
}
