package services

import (
	"errors"
	"fmt"
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
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix() // 1 minute expired
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_ACCESS_KEY))
}

func ExtractAccessToken(e echo.Context) (models.Payload, error) {
	user := e.Get("user").(*jwt.Token)
	fmt.Println(user)
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
	claims["exp"] = time.Now().Add(time.Hour * 24 * 60).Unix() // 60 days expired
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_REFRESH_KEY))
}

func ValidateRefreshToken(refreshToken string) (models.Payload, error) {

	// get accessToken from Refresh Token
	oldRefreshToken, _ := jwt.Parse(refreshToken, func(oldRefreshToken *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_REFRESH_KEY), nil
	})

	// get Payload from Old Access Token
	accToken := oldRefreshToken.Claims.(jwt.MapClaims)["accessToken"].(string)
	oldAccessToken, _ := jwt.Parse(accToken, func(oldAccessToken *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_ACCESS_KEY), nil
	})

	claims := oldAccessToken.Claims.(jwt.MapClaims)

	if claims["email"].(string) != "" {
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
