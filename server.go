package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", homePage)

	e.GET("/spotify", spotify)
	e.GET("/spotify/song", spotifySong)

	e.Logger.Fatal(e.Start(":1323"))
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func spotify(c echo.Context) error {
	return c.String(http.StatusOK, "Spofity API endpoint")
}

func spotifySong(c echo.Context) error {
	return c.String(http.StatusOK, "Current song")
}
