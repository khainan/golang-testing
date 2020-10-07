package models

import (
	"database/sql"
	"net/http"
)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"user-type"`
}

type UsersModel struct {
	DB *sql.DB
}

func (m *UsersModel) FetchAllUsers() (Response, error) {
	var obj Users
	var newData []Users
	var res Response

	sqlStatement := "SELECT * FROM Users"

	rows, err := m.DB.Query(sqlStatement)
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
