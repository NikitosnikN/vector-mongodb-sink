package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"vector-mongodb-sink/internal/app"
)

const (
	typeTextPlain       string = "text/plain"
	typeApplicationJson string = "application/json"
)

func HandlerStoreLogs(c *fiber.Ctx) error {
	contentType := c.GetReqHeaders()["Content-Type"]

	collection := c.Params("collection", "default")

	if contentType == typeTextPlain {
		app.StorePlainText(c.Body(), collection)
	} else if contentType == typeApplicationJson {
		app.StoreJson(c.Body(), collection)
	} else {
		return c.Status(415).SendString(fmt.Sprintf("'%s' unsupported", contentType))
	}

	return c.Status(200).SendString("OK")
}
