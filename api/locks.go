package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
)

type locksApi struct {
	service service.Service
}

func (a *locksApi) Routes(r fiber.Router) {

	r.Get("/locks", withClient(a.List))

}

// List получение списка блокировок на кластере
//  Swagger-spec:
//		@Summary получение списка блокировок на кластере
// 		@Description получение списка блокировок на кластере
// 		@Tags locks
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=serialize.LocksList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/locks [get]
func (a *locksApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetLocks(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get locks")
	}

	return SuccessResponse(ctx, list)
}

// List получение списка блокировок на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка блокировок на сервере 1С Предприятие
// 		@Description получение списка блокировок на сервере 1С Предприятие
// 		@Tags locks
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=serialize.LocksList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/locks [get]
func (a *locksApi) ListApp() {}

// List получение списка блокировок для информационной базы на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary пполучение списка блокировок для информационной базы на сервере 1С Предприятие
// 		@Description получение списка блокировок для информационной базы на сервере 1С Предприятие
// 		@Tags locks
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase name or uuid"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.LocksList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases/{infobase}/locks [get]
func (a *locksApi) ListAppInfobase() {}

// List получение списка блокировок для информационной базы на кластере
//  Swagger-spec:
//		@Summary получение списка блокировок для информационной базы на кластере
// 		@Description получение списка блокировок для информационной базы на кластере
// 		@Tags locks
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase name or uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.LocksList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase}/locks [get]
func (a *locksApi) LisClusterInfobase() {}
