package services

import (
	"errors"
	"time"

	"github.com/furqonzt99/refresh-token/constants"
	"github.com/furqonzt99/refresh-token/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateAccessToken(userId string, email, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_ACCESS_KEY))
}

func ExtractAccessToken(e echo.Context) (models.Payload, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"]
		email := claims["email"]
		role := claims["role"]
		return models.Payload{
			UserID: userId.(string),
			Email:  email.(string),
			Role:   role.(string),
		}, nil
	}
	return models.Payload{}, errors.New("invalid token")
}

func CreateRefreshToken(accessToken string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["accessToken"] = accessToken
	claims["exp"] = time.Now().Add(time.Hour * 24 * 60).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_ACCESS_KEY))
}

func ValidateRefreshToken(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		tokenId := claims["tokenId"]
		return tokenId.(string), nil
	}

	return "", errors.New("invalid token")
}
