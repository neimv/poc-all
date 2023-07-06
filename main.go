package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/neimv/poc-all/views"
)

var PORT = ":8000"

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	envi := os.Getenv("environment")
	fmt.Println(envi)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		response := fmt.Sprintf("%s %s", "Hello World!!", envi)
		return c.String(http.StatusOK, response)
	})
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "esto llego hasta prod")
	})
	e.GET("/users", views.GetUsers)
	e.GET("/users/:id", views.GetUsersId)

	e.Logger.Fatal(e.Start(PORT))
}
