package main

import (
	"godocker/api"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, api.GetApi())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
