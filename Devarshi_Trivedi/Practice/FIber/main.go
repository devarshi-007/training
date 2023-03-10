package main

import (
	"main/book"
	"main/database"

	"github.com/gofiber/fiber"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
}

func main() {
	app := fiber.New()
	database.InitDatabase()
	defer database.Dbconn.Close()
	SetUpRoutes(app)
	app.Listen(3000)
}
