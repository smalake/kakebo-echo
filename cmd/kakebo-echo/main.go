package main

import (
	"net/http"
	// Import Echo v4.
	"github.com/labstack/echo/v4"
)

func main() {
	// Create an Echo instance.
	e := echo.New()

	// Set Route to helloWorld.
	e.GET("/", helloWorld)

	// Start an Echo server.
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
