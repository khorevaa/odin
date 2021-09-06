package api

import (
	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(Response{
		Code:    404,
		Message: "not found",
	})
}

type HTTPError struct {
	Status  string
	Message string
}
