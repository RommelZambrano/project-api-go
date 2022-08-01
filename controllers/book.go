package controllers

import (
	"context"
	"net/http"
	"team6-library/config"
	"team6-library/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = config.GetCollection(config.DB, "books")

func CreateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book
	defer cancel()
	c.BodyParser(&book)
	newBook := models.Book{
		Name:   book.Name,
		Gender: book.Gender,
		Author: book.Author,
		Price:  book.Price,
	}

	result, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	} else {
		return c.Status(http.StatusCreated).JSON(models.Response{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"data": result}})

	}
}

//GET

func GetAllBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var books []models.Book
	defer cancel()

	result, err := bookCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var singleBook models.Book
		if err := result.Decode(&singleBook); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
		}

		books = append(books, singleBook)
	}
	return c.Status(http.StatusOK).JSON(
		models.Response{Status: http.StatusOK,
			Message: "succes",
			Data:    &fiber.Map{"Data": books}})
}

//GET By ID
func GetById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bookId := c.Params("bookId")
	var book models.Book
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(bookId)

	err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&book)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Status:  http.StatusOK,
		Message: "succes",
		Data:    &fiber.Map{"data": book}})

}

//Delete
func Delete(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bookId := c.Params("bookId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(bookId)

	result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			models.Response{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "Book with specified ID not found!"}})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"data": "Book successfully deleted"}})
}
//update Book by id
func UpdateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bookId := c.Params("bookId")
	var book models.Book
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(bookId)
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}
	update := bson.M{
		"name":   book.Name,
		"gender": book.Gender,
		"author": book.Author,
		"price":  book.Price,
	}
	result, err := bookCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	var updateBook models.Book
	if result.MatchedCount == 1 {
		err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updateBook)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
		}
	}
	return c.Status(http.StatusOK).JSON(models.Response{
		Status:  http.StatusOK,
		Message: "Book Update Success",
		Data:    &fiber.Map{"data": updateBook}})
}