package models

import "github.com/gofiber/fiber/v2"

type Book struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Autor string `json:"author"`
	Price float32 `json:"price"`
}

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data  *fiber.Map `json:"data"`
}