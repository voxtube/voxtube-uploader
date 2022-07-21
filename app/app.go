package app

import (
	"log"
	"voxtube-uploader/app/route"
	"voxtube-uploader/config"

	"github.com/gofiber/fiber/v2"
)

func InitialiseApp() {
	app := fiber.New(fiber.Config{
		BodyLimit: 5120 * 1024 * 1024,
	})

	// initialising middleware
	InitialiseMiddleWares(app)

	// global prefix
	api := app.Group("/api")

	// routes
	route.UploadRoute(api)

	log.Fatal(app.Listen(config.GetPort()))
}
