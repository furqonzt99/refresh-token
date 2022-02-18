package middlewares

import (
	"net/http"

	"github.com/furqonzt99/refresh-token/delivery/common"
	"github.com/furqonzt99/refresh-token/services"
	"github.com/labstack/echo/v4"
)

func UserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := services.ExtractAccessToken(c)

		id := c.Param("id")

		if user.Role == "user" && user.UserID != id {
			return c.JSON(http.StatusUnauthorized, common.NewUnauthorizeResponse())
		}
		return next(c)
	}
}
