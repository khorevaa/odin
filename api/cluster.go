package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

type clusterApi struct {
	service service.Service
}

func (a *clusterApi) Routes(r fiber.Router) {

	r.Get("/clusters", withClient(a.List))
	r.Get("/clusters/:cluster", withClient(a.Info))
	r.Post("/clusters", withClient(a.Reg))
	r.Delete("/clusters/:cluster", withClient(a.Unreg))

	router := r.Group("/clusters/:cluster")

	parentApi := []route{
		&infobasesApi{service: a.service},
		&sessionsApi{service: a.service},
		&connectionsApi{service: a.service},
		&managersApi{service: a.service},
		&servicesApi{service: a.service},
		&locksApi{service: a.service},
		&processesApi{service: a.service},
		&licensesApi{service: a.service},
		&clusterAdminApi{service: a.service},
	}

	for _, api := range parentApi {
		api.Routes(router)
	}

}

// List получение списка кластеров на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка кластеров на сервере 1С Предприятие
// 		@Description получение списка кластеров на сервере 1С Предприятие
// 		@Tags clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ClusterInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters [get]
func (a *clusterApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetClusters(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get clusters")
	}

	return SuccessResponse(ctx, list)
}

// Info получение информации о кластере на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение информации о кластере на сервере 1С Предприятие
// 		@Description получение информации о кластере на сервере 1С Предприятие
// 		@Tags clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.ClusterInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster} [get]
func (a *clusterApi) Info(client service.ClientContext, ctx *fiber.Ctx) error {
	val, err := a.service.GetClusterInfo(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get cluster info")
	}

	return SuccessResponse(ctx, val)
}

func (a *clusterApi) Reg(client service.ClientContext, ctx *fiber.Ctx) error {
	return NotImplemented(ctx)
}

func (a *clusterApi) Unreg(client service.ClientContext, ctx *fiber.Ctx) error {
	return NotImplemented(ctx)
}
