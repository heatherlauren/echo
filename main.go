package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/hello/:name", postHandler)
	e.Logger.Fatal(e.Start(":1337"))
}

func postHandler(c echo.Context) error {
	c.Response().Header().Add("X-Secret", "shotsfired")

	name := c.Param("name")

	return c.HTML(http.StatusOK, "<h1>Hello " + name + "</h1>")
}
