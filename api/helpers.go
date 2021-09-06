package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

func withClient(fn func(client service.ClientContext, ctx *fiber.Ctx) error) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		client, err := service.GetClientContext(ctx)

		if err != nil {
			return ErrorResponse(ctx, err, "error get context client")
		}

		return fn(client, ctx)

	}
}
