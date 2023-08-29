package main

import (
	"log"

	config "github.com/dinhnguyen1812002/go-crud/config"
	controller "github.com/dinhnguyen1812002/go-crud/Controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func SetupRouter(app *fiber.App)  {
	app.Get("/api/book", controller.GetAllBooks)
	app.Get("/api/book/:id", controller.GetById)
	app.Post("/api/book",controller.AddBook)
	app.Put("/api/book/:id",controller.Update)
	app.Delete("/api/book/:id",controller.Delete)

}
func main() {

	app := fiber.New()
	config.Connect()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))
	SetupRouter(app)
	log.Fatal(app.Listen(":8080"))
}