package main

import (
	"log"
	"net/http"

	"github.com/khainan/controllers"
	"github.com/khainan/db"
	"github.com/khainan/models"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Bikin instance UsersModel dengan parameter db
	usersModel := models.NewUsersModel(database)
	usersCntrl := controllers.NewUsersController(usersModel)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Dariiiii ECHOOOO")
	})

	e.GET("/users", usersCntrl.FetchAllUsers)
	e.GET("/users/:id", usersCntrl.FetchSingleUser)
	e.POST("/users", usersCntrl.CreateSingleUser)
	e.DELETE("/users/:id", usersCntrl.DeleteSingleUser)
	e.PUT("/users/:id", usersCntrl.UpdateSingleUser)

	e.Logger.Fatal(e.Start(":1234"))
}
