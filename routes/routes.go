package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Dariiiii ECHOOOO")
	})

	// e.GET("/users", controllers.FetchAllUsers)

	return e
}
