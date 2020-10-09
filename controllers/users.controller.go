package controllers

import (
	"net/http"

	"github.com/khainan/models"
	"github.com/labstack/echo"
)

// Bikin struct yang punya UserModel untuk dipakai FetchAllUsers
type UserController struct {
	userModel *models.UserModel
}

func NewUserController(userModel *models.UserModel) *UserController {
	return &UserController{userModel: userModel}
}

func (cntrl *UserController) FetchAllUsers(c echo.Context) error {
	userName := c.QueryParam("name")
	result, err := cntrl.userModel.FetchAllUsers(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (cntrl *UserController) FetchSingleUser(c echo.Context) error {
	id := c.Param("id")
	result, err := cntrl.userModel.FetchSingleUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (cntrl *UserController) CreateSingleUser(c echo.Context) error {
	result, err := cntrl.userModel.CreateSingleUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (cntrl *UserController) DeleteSingleUser(c echo.Context) error {
	result, err := cntrl.userModel.DeleteSingleUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (cntrl *UserController) UpdateSingleUser(c echo.Context) error {
	result, err := cntrl.userModel.UpdateSingleUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
