package models

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"user-type"`
}

type UsersModel struct {
	db *sql.DB
}

func NewUsersModel(db *sql.DB) *UsersModel {
	return &UsersModel{db: db}
}

func (m *UsersModel) FetchAllUsers() (Response, error) {
	var obj Users
	var newData []Users
	var res Response

	sqlStatement := "SELECT * FROM Users"

	rows, err := m.db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.UserType)
		if err != nil {
			return res, err
		}
		newData = append(newData, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = newData

	return res, nil
}

func (m *UsersModel) FetchSingleUser(c echo.Context) (Response, error) {
	var obj Users
	var newData []Users
	var res Response
	id := c.Param("id")

	sqlStatement := "SELECT * FROM Users WHERE id = " + id

	rows, err := m.db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.UserType)
		if err != nil {
			return res, err
		}
		newData = append(newData, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = newData

	return res, nil
}
