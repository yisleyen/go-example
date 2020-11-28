package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpoint istek geldi")
}

func main() {
	e := echo.New()

	e.GET("/main", mainHandler)

	e.Start(":8080")
}
