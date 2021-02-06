package main

import (
	"log"
	"os"

	"github.com/48d90782/GoPlayground/secret"
	"github.com/48d90782/GoPlayground/secret/application"
	"github.com/48d90782/GoPlayground/secret/pkg/shared_data"
	"github.com/48d90782/GoPlayground/secret/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	// configuration file path
	var configPath string
	// debug mode -> use pprof, and microservice unsafe mode (turned off some checks)
	// TODO figure out, what checks should be turned off
	var debug bool
	cliApp := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"l"},
				Value:       "appspace.json",
				Usage:       "appspace configuration file",
				Destination: &configPath,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "microservice debug mode",
				Destination: &debug,
			},
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		// we can't use uber/zap here, because it's not initialized at the time
		log.Fatal("failed to initialize CLI")
	}

	logger, err := zap.NewProduction()
	if debug {
		// NewDevelopment builds a development Logger that writes DebugLevel and above
		// logs to standard error in a human-friendly format.
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatal("failed to initialize development zap logger")
		}
	}
	defer func() {
		// ignore error explicitly with _
		_ = logger.Sync()
	}()

	// log.Fatal will call os.Exit(1)
	if err != nil {
		log.Fatal("failed to initialize zap logger")
	}

	// Initialize shared data between Websocket <-> Rabbitmq
	shared := shared_data.NewSharedData()

	// application layer
	appL := application.NewAppLayer()
	// domain layer
	dom := secret.NewDomainLayer(appL)

	// TODO temporary
	_ = dom

	// initialize infrastructure level (fiber, websockets)
	app := &fiber.App{}
	// initialize websocket transport
	// TODO options
	err = transport.NewWSHandler(app, logger, shared)
	if err != nil {
		logger.Error("failed to initialize websocket handler")
		return
	}

	// TODO add to the config
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Error("failed to instantiate RabbitMQ connection", zap.Error(err))
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	err = transport.NewRabbitMQHandler(conn, logger, shared)
	if err != nil {
		logger.Error("failed to initialize websocket handler", zap.Error(err))
		return
	}

	// TODO listen address to the config
	log.Fatal(app.Listen(":3000"))
}
