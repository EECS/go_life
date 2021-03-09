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

	var _connections map[int]websocket.Conn = make(map[int]websocket.Conn)
	var conCounter = 1

	app.Get("/ws/", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed")) // true
		//log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		var connectionID int = conCounter
		conCounter++

		_connections[connectionID] = *c

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s", msg)

			log.Println("current connectionID:", connectionID)
			for loopConID, connection := range _connections {
				log.Println("conLoop:", loopConID)
				if loopConID == connectionID {
					log.Println("Skipping:", loopConID)
					continue
				}

				// Send Msg, print if error sending
				if err = connection.WriteMessage(1, msg); err != nil {
					log.Println("write:", err, mt)
					break
				}
				log.Println("Sending to: "+connection.Params("id"), msg)
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
