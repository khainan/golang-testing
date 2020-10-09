package models

import (
	"database/sql"
	"fmt"
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

func (m *UserModel) CreateSingleUser(name string, userType string) (Response, err error) {
	sqlStatement := "INSERT Users (name, user_type) VALUES (?, ?)"
	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return
	}

	stmt.Exec(name, userType)

	return
}

func (m *UserModel) DeleteSingleUser(id string) (Response, err error) {
	sqlStatement := "DELETE FROM Users WHERE id = ?"
	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return
	}

	stmt.Query(id)

	return
}

func (m *UserModel) UpdateSingleUser(id string, newName string) (Response, err error) {
	sqlStatement := "UPDATE Users SET name = ? WHERE id = ?"
	stmt, err := m.db.Prepare(sqlStatement)
	if err != nil {
		return
	}

	stmt.Exec(newName, id)

	return
}
