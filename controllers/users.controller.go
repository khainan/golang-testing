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
	name := c.FormValue("name")
	userType := c.FormValue("user_type")
	result, err := cntrl.userModel.CreateSingleUser(name, userType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (cntrl *UserController) DeleteSingleUser(c echo.Context) error {
	id := c.Param("id")
	result, err := cntrl.userModel.DeleteSingleUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (cntrl *UserController) UpdateSingleUser(c echo.Context) error {
	id := c.Param("id")
	newName := c.FormValue("name")
	result, err := cntrl.userModel.UpdateSingleUser(id, newName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
