package controller

import (
	entity "github.com/dinhnguyen1812002/go-crud/Entity"
	"github.com/dinhnguyen1812002/go-crud/config"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var books []entity.Book
	config.Database.Find(&books)
	return c.Status(200).JSON(books)
}
func GetAllBooks(c *fiber.Ctx) error  {
	var books []entity.Book
	config.Database.Find(&books)
	return c.Status(200).JSON(books)
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var books entity.Book
	result := config.Database.Find(&books, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&books)

}

func AddBook(c *fiber.Ctx) error {
	book := new(entity.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	config.Database.Create(&book)
	return c.Status(201).JSON(book)
}
func Update(c *fiber.Ctx) error {
	book := new(entity.Book)
	id := c.Params("id")

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&book)
	return c.Status(200).JSON(book)
}
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var book entity.Book

	result := config.Database.Delete(&book, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
