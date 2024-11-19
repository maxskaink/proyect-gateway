package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maxskaink/proyect-gateway/routes"
)

func main() {
	app := fiber.New()

	routes.AddRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
