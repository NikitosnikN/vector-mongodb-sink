package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandlerHealthcheck(c *fiber.Ctx) error {
	return c.Status(200).SendString("OK")
}
