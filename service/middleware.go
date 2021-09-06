package service

import (
	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func Middleware(s Service) fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) (err error) {

		c.Context().SetUserValue("service", s)

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
