package main

import (
	api_fiber "github.com/maxskaink/proyect-gateway/apiFiber"
)

func main() {
	app := api_fiber.NewApiFiber()
	app.Listen()
}
