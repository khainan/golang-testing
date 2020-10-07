package controllers

import (
	"net/http"

	"github.com/khainan/models"
	"github.com/labstack/echo"
)

// Bikin struct yang punya UsersModel untuk dipakai FetchAllUsers
type UsersController struct {
	usersModel *models.UsersModel
}

func NewUsersController(usersModel *models.UsersModel) *UsersController {
	return &UsersController{usersModel: usersModel}
}

func (cntrl *UsersController) FetchAllUsers(c echo.Context) error {
	result, err := cntrl.usersModel.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (cntrl *UsersController) FetchSingleUser(c echo.Context) error {
	result, err := cntrl.usersModel.FetchSingleUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
