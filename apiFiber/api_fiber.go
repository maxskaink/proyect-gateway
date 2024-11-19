package api_fiber

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxskaink/proyect-gateway/routes"
)

type ApiFiber struct {
	app  *fiber.App
	PORT string
}

// Create a new instance of ApiFiber
// Make sure you have the PORT_API enviroment variable set
func NewApiFiber() *ApiFiber {
	// Create a new fiber app
	app := fiber.New()

	// Add the necesary routes
	routes.AddRoutes(app)

	// Config the port to the api
	PORT_API := os.Getenv("PORT_API")

	if PORT_API == "" {
		PORT_API = "8000"
		fmt.Printf("warning: PORT_API not found en enviroment variables, using default %s\n", PORT_API)
	}

	//Return the new instance
	return &ApiFiber{
		app,
		PORT_API,
	}
}

// Start listening the api
// This function is blocking
func (app *ApiFiber) Listen() {
	log.Fatal(app.app.Listen(":" + app.PORT))
}
