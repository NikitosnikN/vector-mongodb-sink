package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"vector-mongodb-sink/internal/handlers"
)
import "vector-mongodb-sink/internal/config"

func initMiddlewares(app *fiber.App) {
	// initialize basic auth middleware if credentials are provided
	if config.Env.AuthUsername != "" && config.Env.AuthPassword != "" {
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				config.Env.AuthUsername: config.Env.AuthPassword,
			},
		}))
	}

}

func initRoutes(app *fiber.App) {
	app.Get("/ht", handlers.HandlerHealthcheck)

	app.Post("/", handlers.HandlerStoreLogs)
	app.Post("/:collection", handlers.HandlerStoreLogs)

}

func main() {
	app := fiber.New()

	initMiddlewares(app)
	initRoutes(app)

	app.Listen(fmt.Sprintf("%s:%d", config.Env.Host, config.Env.Port))
}
