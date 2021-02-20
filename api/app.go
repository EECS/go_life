package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.Send([]byte("Welcome to the Game of Life!"))
	// })

	app.Static("/", "./public")

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	var _connections map[string]websocket.Conn = make(map[string]websocket.Conn)

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		var connectionID string = c.Params("id")
		_connections[connectionID] = *c

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s", msg)

			log.Println("current connectionID:", connectionID)
			for loopConID, conection := range _connections {
				log.Println("conLoop:", loopConID)
				if loopConID == connectionID {
					log.Println("Skipping:", loopConID)
					continue
				}
				conection.WriteMessage(1, msg) //<- For some reason this is STILL sending the message to the connection we skipped... (see console log in browser)
				log.Println("Sending to: "+conection.Params("id"), msg)
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
	// Access the websocket server: ws://localhost:8080/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
