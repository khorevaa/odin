package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

type healthAppApi struct {
	service service.Service
}

func (a *healthAppApi) Routes(r fiber.Router) {

	r.Get("/health", withClient(a.health))
}

// health запрос о состонии сервера приложений 1С Предприятие
//  Swagger-spec:
//		@Summary запрос о состонии сервера приложений 1С Предприятие
// 		@Description запрос о состонии сервера приложений 1С Предприятие
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} StatusResponse
//		@Failure 500 {object} StatusResponse
// 		@Router /app/{app}/health [get]
func (a *healthAppApi) health(client service.ClientContext, ctx *fiber.Ctx) error {

	ok, err := client.HealthCheck()

	if ok && err == nil {
		return ctx.Status(fiber.StatusOK).JSON(StatusResponse{
			Status: true,
		})
	}

	return ctx.Status(fiber.StatusServiceUnavailable).JSON(StatusResponse{
		Status: false,
		Err:    err.Error(),
	})

}
