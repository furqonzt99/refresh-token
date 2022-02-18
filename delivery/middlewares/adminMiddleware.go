package middlewares

import (
	"net/http"

	"github.com/furqonzt99/refresh-token/delivery/common"
	"github.com/furqonzt99/refresh-token/services"
	"github.com/labstack/echo/v4"
)

func AdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := services.ExtractAccessToken(c)

		if user.Role != "admin" {
			return c.JSON(http.StatusUnauthorized, common.NewUnauthorizeResponse())
		}
		return next(c)
	}
}
