package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Username string "json:'username'"
	Name     string "json:'name'"
	Surname  string "json:'surname'"
}

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

func addUser(c echo.Context) error {
	user := User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	fmt.Println(user.Username)
	fmt.Println(user.Name)
	fmt.Println(user.Surname)
	return c.String(http.StatusOK, "Başarılı")
}

func main() {
	e := echo.New()

	e.GET("/main", mainHandler)
	e.GET("/user/:data", getUser)
	e.POST("/user", addUser)

	e.Start(":8080")
}
