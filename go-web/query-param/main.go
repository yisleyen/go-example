package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpoint istek geldi")
}

func getUser(c echo.Context) error {
	dataType := c.Param("data")

	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Username: %s Name: %s Surname: %s", username, name, surname))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	return c.String(http.StatusBadRequest, "Yalnızca JSON ve String tipinde istekten bulunabilirsiniz.")
}

func main() {
	e := echo.New()

	e.GET("/main", mainHandler)
	e.GET("/user/:data", getUser)

	e.Start(":8080")
}
