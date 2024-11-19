package controllers

import "github.com/gofiber/fiber/v2"

func ProxyHandler(route string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.SendString("Proxy to " + route)
	}
}
