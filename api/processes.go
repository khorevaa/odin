package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

type processesApi struct {
	service service.Service
}

func (a *processesApi) Routes(r fiber.Router) {

	r.Get("/processes", withClient(a.List))
	r.Get("/processes/:process", withClient(a.Info))

}

// List получение списка процессов на кластере
//  Swagger-spec:
//		@Summary получение списка процессов на кластере
// 		@Description получение списка процессов на кластере
// 		@Tags processes
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ProcessInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/processes [get]
func (a *processesApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetProcesses(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get processes")
	}

	return SuccessResponse(ctx, list)
}

// List получение списка процессов на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка процессов на сервере 1С Предприятие
// 		@Description получение списка процессов на сервере 1С Предприятие
// 		@Tags processes
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ProcessInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/processes [get]
func (a *processesApi) ListApp() {}

// Info получение информации опроцессе на кластере
//  Swagger-spec:
//		@Summary получение информации опроцессе на кластере
// 		@Description получение информации опроцессе на кластере
// 		@Tags processes
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param process path string true "uuid process"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=serialize.ProcessInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/processes/{process} [get]
func (a *processesApi) Info(client service.ClientContext, ctx *fiber.Ctx) error {

	info, err := a.service.GetProcessInfo(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get process info")
	}

	return SuccessResponse(ctx, info)

}

// Info получение информации опроцессе на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение информации опроцессе на сервере 1С Предприятие
// 		@Description получение информации опроцессе на сервере 1С Предприятие
// 		@Tags processes
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param process path string true "uuid process"
// 		@Param cluster-id query string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=serialize.ProcessInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/processes/{process} [get]
func (a *processesApi) InfoApp() {}
