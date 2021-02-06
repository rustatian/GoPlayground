package transport

import (
	"fmt"
	"sync"

	"github.com/48d90782/GoPlayground/secret/pkg/shared_data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
)

type Storage struct {
	Data *sync.Map
}

type Websocket struct {
	Id [16]byte
}

// TODO logger
func NewWSHandler(app *fiber.App, log *zap.Logger, shared shared_data.Shared) error {
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		fmt.Println(c.Locals("allowed"))  // true
		fmt.Println(c.Params("id"))       // 123
		fmt.Println(c.Query("v"))         // 1.0
		fmt.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		// for every connection
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				fmt.Println("read:", err)
				break
			}
			fmt.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				fmt.Println("write:", err)
				break
			}
		}

	}))

	return nil
}
