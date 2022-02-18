package users

import (
	"net/http"

	"github.com/furqonzt99/refresh-token/delivery/common"
	"github.com/furqonzt99/refresh-token/models"
	"github.com/furqonzt99/refresh-token/repository/users"
	"github.com/furqonzt99/refresh-token/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repository users.UserInterface
}

func NewUserController(userInterface users.UserInterface) *UserController {
	return &UserController{Repository: userInterface}
}

func (uc UserController) Create(c echo.Context) error {
	var userRequest CreateUserRequest

	c.Bind(&userRequest)

	if err := c.Validate(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	password, _ := services.Hashpwd(userRequest.Password)
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: password,
	}

	userDB, err := uc.Repository.Create(user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, common.ErrorResponse(http.StatusNotAcceptable, "Email already exist"))
	}

	response := UserResponse{
		ID:    userDB.ID,
		Name:  userDB.Name,
		Email: userDB.Email,
		Role:  userDB.Role,
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(response))
}

func (uc UserController) ReadAll(c echo.Context) error {
	users, _ := uc.Repository.ReadAll()

	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(response))
}

func (uc UserController) ReadOne(c echo.Context) error {
	id := c.Param("id")

	user, err := uc.Repository.ReadOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	response := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(response))
}

func (uc UserController) Update(c echo.Context) error {
	id := c.Param("id")

	var userRequest UpdateUserRequest

	c.Bind(&userRequest)

	if err := c.Validate(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	var password string
	if userRequest.Password != "" {
		password, _ = services.Hashpwd(userRequest.Password)
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: password,
	}

	_, err := uc.Repository.Update(id, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (uc UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	_, err := uc.Repository.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
