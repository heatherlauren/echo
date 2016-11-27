package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", getHandler)
	e.Logger.Fatal(e.Start(":1337"))
}

func getHandler(c echo.Context) error {
	c.Response().Header().Add("X-Secret", "shotsfired")
	return c.HTML(http.StatusOK, "<h1>Hello jonny</h1>")
}
