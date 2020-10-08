package models

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type UsersModel struct {
	db *sql.DB
}

func NewUsersModel(db *sql.DB) *UsersModel {
	return &UsersModel{db: db}
}

func (m *UsersModel) FetchAllUsers(c echo.Context) (Response, error) {
	var obj Users
	var newData []Users
	var res Response

	nameParams := c.QueryParam("name")

	sqlStatement := "SELECT * FROM Users"

	if nameParams != "" {
		sqlStatement = "SELECT * FROM Users WHERE name LIKE " + "'" + "%" + nameParams + "%" + "'"
	}

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

func (m *UsersModel) CreateSingleUser(c echo.Context) (Response, error) {
	var res Response

	name := c.FormValue("name")
	userType := c.FormValue("user_type")

	sqlStatement := "INSERT Users (name, user_type) VALUES (?, ?)"
	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, userType)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Adding User"
	res.Data = map[string]int64{
		"last_inserted_id:": lastInsertedId,
	}

	return res, nil
}

func (m *UsersModel) DeleteSingleUser(c echo.Context) (Response, error) {
	var res Response

	id := c.Param("id")
	sqlStatement := "DELETE FROM Users WHERE id = " + id
	stmt, err := m.db.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Deleting User"
	res.Data = stmt

	return res, nil
}

func (m *UsersModel) UpdateSingleUser(c echo.Context) (Response, error) {
	var res Response

	id := c.Param("id")
	newName := c.FormValue("name")

	sqlStatement := "UPDATE Users SET name = ? WHERE id = ?"
	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(newName, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected:": rowsAffected,
	}

	return res, nil
}
