package main

import (
	"team6-library/config"
	"team6-library/routes"

	"github.com/gofiber/fiber/v2"
)

func main (){
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx) error{
		return c.SendString("Hello, World!")
	})
	config.ConnectDB()
	routes.BookRoutes(app)
	app.Listen(":8080")
}