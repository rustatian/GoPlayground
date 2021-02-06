package main

import (
	"log"
	"os"

	"github.com/48d90782/GoPlayground/secret/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	var configPath string
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
		log.Fatal("failed to initialize CLI")
	}

	logger, _ := zap.NewProduction()
	defer func() {
		// ignore error explicitly with _
		_ = logger.Sync()
	}()

	app := &fiber.App{}

	err := transport.NewWSHandler(app)
	if err != nil {
		panic(err)
	}

	log.Fatal(app.Listen(":3000"))
}
