package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/khainan/config"
	"github.com/khainan/controllers"
	"github.com/khainan/models"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to %s: %s\n", connectionString, err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("DNS Invalid")
	}

	// Bikin instance UsersModel dengan parameter db
	usersModel := models.UsersModel{DB: db}
	usersCntrl := controllers.UsersController{UsersModel: usersModel}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Dariiiii ECHOOOO")
	})
	e.GET("/users", usersCntrl.FetchAllUsers)
	e.Logger.Fatal(e.Start(":1234"))
}
