package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxskaink/proyect-gateway/controllers"
)

func AddRoutes(app *fiber.App) {
	group := app.Group("/api")

	group.Get("/articles", controllers.ProxyHandler("http://localhost:3000/api/articles"))
}
