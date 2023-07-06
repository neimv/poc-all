package views

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/neimv/poc-all/models"
)

// Data fake
func initDataFake() [10]models.User {
	users := [10]models.User{
		{Id: 0, Name: "prueba", Email: "prueba@test.com"},
		{Id: 1, Name: "emma", Email: "emma@test.com"},
		{Id: 2, Name: "su", Email: "su@test.com"},
		{Id: 3, Name: "zu", Email: "zu@test.com"},
		{Id: 4, Name: "gaby", Email: "gaby@test.com"},
		{Id: 5, Name: "sasori", Email: "sasori@test.com"},
		{Id: 6, Name: "raquel", Email: "raquel@test.com"},
		{Id: 7, Name: "sti", Email: "sti@test.com"},
		{Id: 8, Name: "leo", Email: "leo@test.com"},
		{Id: 9, Name: "casa", Email: "casa@test.com"},
	}

	return users
}

func GetUsers(c echo.Context) error {
	users := initDataFake()

	return c.JSON(http.StatusOK, users)
}

func GetUsersId(c echo.Context) error {
	users := initDataFake()

	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, users[id])
}
