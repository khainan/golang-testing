package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khainan/config"
)

func Init() (*sql.DB, error) {
	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
