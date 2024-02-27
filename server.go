package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", homePage)

	e.GET("/spotify/song", spotifySong)
	e.GET("/spotify/album", spotifyAlbum)

	e.Logger.Fatal(e.Start(":1323"))
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func spotifySong(c echo.Context) error {
	return c.String(http.StatusOK, "Current song")
}

func spotifyAlbum(c echo.Context) error {
	return c.String(http.StatusOK, "Current album")
}
