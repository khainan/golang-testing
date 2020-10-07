package controllers

import (
	"net/http"

	"github.com/khainan/models"
	"github.com/labstack/echo"
)

// Bikin struct yang punya UsersModel untuk dipakai FetchAllUsers
type UsersController struct {
	UsersModel models.UsersModel
}

func (cntrl *UsersController) FetchAllUsers(c echo.Context) error {
	result, err := cntrl.UsersModel.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
