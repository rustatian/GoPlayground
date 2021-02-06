package transport

import (
	"github.com/48d90782/GoPlayground/secret/pkg/shared_data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
)

type Websocket struct {
	Id [16]byte
}

// TODO logger
func NewWSHandler(app *fiber.App, log *zap.Logger, shared shared_data.Shared) error {
	// websocket middleware
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
		log.Debug("allowed", zap.Bool("allowed", c.Locals("allowed").(bool)))
		log.Debug("locals", zap.String("params", c.Params("id")))
		log.Debug("query", zap.String("query", c.Query("v")))
		log.Debug("session", zap.String("session", c.Cookies("session")))

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)

		// set the shared data
		// at the moment we should have all data parsed via middlewares
		// all data should be in the locals
		shared.Set()
		// for every connection
		for {
			// returned messageType is either TextMessage or BinaryMessage.
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Debug("read:", zap.Error(err))
				break
			}
			log.Debug("receive", zap.String("recv", string(msg)))

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Debug("write:", zap.Error(err))
				break
			}
		}

	}))

	return nil
}
