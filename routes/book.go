package routes

import (
	"team6-library/controllers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Post("/books", controllers.CreateBook)
	app.Get("/books", controllers.GetAllBooks)
	app.Get("/books/:bookId", controllers.GetById)
	app.Delete("/books/:bookId", controllers.Delete)
	app.Put("/books/:bookId", controllers.UpdateBook)
}
