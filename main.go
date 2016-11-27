package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Payload struct {
		Secret string `json:"secret"`
}

func main() {
	e := echo.New()
	e.POST("/hello/:name", postHandler)
	e.Logger.Fatal(e.Start(":1337"))
}

func postHandler(c echo.Context) error {
	payload := new(Payload)

	if err := c.Bind(payload); err != nil {
		return err
	}

	secret := payload.Secret

	c.Response().Header().Add("X-Secret", secret)

	name := c.Param("name")
	return c.HTML(http.StatusOK, "<h1>Hello " + name + "</h1>")
}
