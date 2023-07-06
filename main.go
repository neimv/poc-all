package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neimv/poc-all/views"
)

var PORT = ":8000"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!")
	})
	e.GET("/users", views.GetUsers)
	e.GET("/users/:id", views.GetUsersId)

	e.Logger.Fatal(e.Start(PORT))
}
