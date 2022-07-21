package app

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func InitialiseMiddleWares(app *fiber.App) {
	// cors
	app.Use(cors.New())

	// compressor
	app.Use(compress.New())

	// logger
	app.Use(logger.New())

	// route middle wares
	app.Get("/monitor", monitor.New())

	// swagger
	app.Get("/docs/*", swagger.Handler) // default
}
