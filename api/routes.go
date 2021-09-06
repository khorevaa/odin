package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

type route interface {
	Routes(r fiber.Router)
}

func Routes(app *fiber.App, s service.Service) {

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")
	routes := []route{
		&AppApi{
			service: s,
		},
		&healthApi{
			service: s,
			name:    "API Remote Administration for 1S.Enterprise Application Servers",
			version: "1.0",
			route:   "/api/v1",
		},
	}
	for _, r := range routes {
		r.Routes(v1)
	}

}
