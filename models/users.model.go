package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

func (m *UserModel) FetchAllUsers(userName string) (users []User, err error) {
	sqlStatement := "SELECT * FROM Users"
	if userName != "" {
		userName = fmt.Sprintf("%%%s%%", userName)
		sqlStatement = sqlStatement + " WHERE name LIKE ?"
	}

	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return
	}

	rows, err := stmt.Query()
	if userName != "" {
		rows, err = stmt.Query(userName)
	}
	if err != nil {
		return
	}
	defer rows.Close()

	users = make([]User, 0)
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Name, &user.UserType); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (m *UserModel) FetchSingleUser(id string) (users []User, err error) {
	sqlStatement := "SELECT * FROM Users WHERE id = ?"
	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return
	}
	defer rows.Close()

	users = make([]User, 0)
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Name, &user.UserType); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func (m *UserModel) CreateSingleUser(c echo.Context) (Response, error) {
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

func (m *UserModel) DeleteSingleUser(c echo.Context) (Response, error) {
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

func (m *UserModel) UpdateSingleUser(c echo.Context) (Response, error) {
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
