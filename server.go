package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", homePage)
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	for i := 0; i < 1; i++ {

	}
	return c.String(http.StatusOK, id)
}
