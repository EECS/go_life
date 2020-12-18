package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte ("Welcome to the Game of Life!"))
	})

	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}

	_ = app.Listen(":"+port)
}