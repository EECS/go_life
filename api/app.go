package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main(){
	app := fiber.New()

	defaultPort := "8080"
	port := os.Getenv("PORT")

	if port == ""{
		port = defaultPort
	}

	app.Use(cors.New())

	// Optional middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if c.Get("host") == "localhost:"+port {
			c.Locals("Host", "Localhost:"+port)
			return c.Next()
		}
		return c.Status(403).SendString("Request origin not allowed")
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte ("Welcome to the Game of Life!"))
	})

	// Upgraded websocket request
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host"))
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	err := app.Listen(":"+port)

	if err != nil{
		log.Fatal(err.Error())
	}
}