package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

type managersApi struct {
	service service.Service
}

func (a *managersApi) Routes(r fiber.Router) {

	r.Get("/managers", withClient(a.List))
	r.Get("/managers/:manager", withClient(a.Info))

}

// List получение списка менеджеров на кластере
//  Swagger-spec:
//		@Summary получение списка менеджеров на кластере
// 		@Description получение списка менеджеров на кластере
// 		@Tags managers
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ManagerInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/managers [get]
func (a *managersApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetManagers(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get managers")
	}

	return SuccessResponse(ctx, list)
}

// List получение списка менеджеров на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка менеджеров на сервере 1С Предприятие
// 		@Description получение списка менеджеров на сервере 1С Предприятие
// 		@Tags managers
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ManagerInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/managers [get]
func (a *managersApi) ListApp() {}

func (a *managersApi) Info(client service.ClientContext, ctx *fiber.Ctx) error {
	return NotImplemented(ctx)
}
