package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Payload struct {
	Secret string `json:"secret"`
}

func main() {
	// Create a new Echo server instance
	e := echo.New()
	// Create a new route that responds to a POST request by running postHandler
	e.POST("/hello/:name", postHandler)
	// Start the server on post 1337
	e.Logger.Fatal(e.Start(":1337"))
}

func postHandler(c echo.Context) error {
	// Bind JSON into the Payload structure and set to variable 'payload'
	payload := new(Payload)
	// If there's a binding error, return error
	if err := c.Bind(payload); err != nil {
		return err
	}
	// Set JSON "secret" to variable 'secret'
	secret := payload.Secret
	// Set custom response header
	c.Response().Header().Add("X-Secret", secret)
	// Get name parameter from URL and set it to variable 'name'
	name := c.Param("name")
	// Return HTML
	return c.HTML(http.StatusOK, "<h1>Hello "+name+"</h1>")
}
