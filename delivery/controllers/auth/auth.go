package auth

import (
	"net/http"

	"github.com/furqonzt99/refresh-token/delivery/common"
	"github.com/furqonzt99/refresh-token/models"
	"github.com/furqonzt99/refresh-token/repository/auth"
	"github.com/furqonzt99/refresh-token/services"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Repository auth.AuthInterface
}

func NewAuthController(authInterface auth.AuthInterface) *AuthController {
	return &AuthController{Repository: authInterface}
}

func (ac AuthController) Register(c echo.Context) error {
	var registerRequest RegisterRequest

	c.Bind(&registerRequest)

	if err := c.Validate(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	password, _ := services.Hashpwd(registerRequest.Password)
	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: password,
	}

	_, err := ac.Repository.Register(user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, common.ErrorResponse(http.StatusNotAcceptable, "Email already exist"))
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (ac AuthController) Login(c echo.Context) error {
	var loginRequest LoginRequest

	c.Bind(&loginRequest)

	if err := c.Validate(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	user, err := ac.Repository.Login(loginRequest.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "User not found"))
	}

	ok, err := services.Checkpwd(loginRequest.Password, user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "Wrong password"))
	}

	var accessToken string
	var refreshToken string

	if ok {
		accessToken, _ = services.CreateAccessToken(user.ID, user.Email, user.Role)
		refreshToken, _ = services.CreateRefreshToken(accessToken)
	}

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(response))
}
