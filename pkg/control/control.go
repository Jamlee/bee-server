package control

import (
	"net/http"

	"github.com/labstack/echo"
)

func StartServer(peerAddress string) {
	e := echo.New()
	e.Static("/", "static")
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(peerAddress))
}
