package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	errors2 "github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/ras-client/serialize"
	"sync"
)

type connectionsApi struct {
	service service.Service
}

func (a *connectionsApi) Routes(r fiber.Router) {

	r.Get("/connections", withClient(a.List))
	r.Post("/connections/terminate", withClient(a.Terminate))
	r.Delete("/connections/:connection.:process", withClient(a.TerminateOne))

}

// List получение списка подключений на кластере
//  Swagger-spec:
//		@Summary получение списка подключений на кластере
// 		@Description пполучение списка подключений на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/connections [get]
func (a *connectionsApi) List(client service.ClientContext, ctx *fiber.Ctx) error {
	list, err := a.service.GetConnections(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get connections")
	}

	return SuccessResponse(ctx, list)
}

// List получение списка подключений для информационной базы на кластере
//  Swagger-spec:
//		@Summary получение списка подключений для информационной базы на кластере
// 		@Description получение списка подключений для информационной базы на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
//		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase}/connections [get]
func (a *connectionsApi) ListClusterInfobase() {}

// List получение списка подключений на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка подключений на сервере 1С Предприятие
// 		@Description получение списка подключений на сервере 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
//		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/connections [get]
func (a *connectionsApi) ListApp() {}

// List получение списка подключений для информационной базы на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка подключений для информационной базы на сервер 1С Предприятие
// 		@Description получение списка подключений для информационной базы на сервер 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
//		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases/{infobase}/connections [get]
func (a *connectionsApi) ListAppInfobase() {}

// TerminateOne отключение подключения на кластере
//  Swagger-spec:
//		@Summary отключение подключения на кластере
// 		@Description отключение подключения на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param connection path string true "connection uuid"
// 		@Param process path string true "process uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=models.TerminateConnectionSig}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/connections/{connection}.{process} [delete]
func (a *connectionsApi) TerminateOne(client service.ClientContext, ctx *fiber.Ctx) error {

	connection, ok := service.GetContextValue(ctx, "connection")
	if !ok {
		return ErrorResponse(ctx, errors2.BadRequest.New("connection id must be set"), "error terminate connection")
	}

	process, ok := service.GetContextValue(ctx, "process")

	if !ok {
		return ErrorResponse(ctx, errors2.BadRequest.New("process id must be set"), "error terminate connection")
	}

	connectionID, err := connection.UUID()
	processID, err := process.UUID()

	connectionSig, err := a.service.TerminateConnection(client, processID, connectionID)

	if err != nil {
		return ErrorResponse(ctx, err, "error terminate connection")
	}

	return SuccessResponse(ctx, connectionSig)
}

// TerminateOne отключение подключения на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary отключение подключения на сервере 1С Предприятие
// 		@Description отключение подключения на сервере 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param connection path string true "connection uuid"
// 		@Param process path string true "process uuid"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=models.TerminateConnectionSig}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/connections/{connection}.{process} [delete]
func (a *connectionsApi) TerminateOneApp() {}

// Terminate отключение списка подключений на кластере
//  Swagger-spec:
//		@Summary отключение списка подключений на кластере
// 		@Description отключение списка подключений или по информационной базе на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param req body models.TerminateConnectionsRequest true "request"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
//		@Success 200 {object} Response{data=models.TerminateConnectionsResponse}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/connections/terminate [post]
func (a *connectionsApi) Terminate(client service.ClientContext, ctx *fiber.Ctx) error {

	var body models.TerminateConnectionsRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		return ErrorResponse(ctx, err, "error parse body TerminateConnectionsRequest")
	}

	respond := &models.TerminateConnectionsResponse{}
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	switch {

	case len(body.InfobaseID) > 0:

		list, err := a.service.GetInfobaseConnections(client, body.InfobaseID)
		if err != nil {
			return ErrorResponse(ctx, err, "error get infobase connections")
		}

		list.Filter(filterServiceConnections).Each(func(info *serialize.ConnectionShortInfo) {

			wg.Add(1)

			go func(info *serialize.ConnectionShortInfo) {
				err := client.DisconnectConnection(info.ClusterID, info.Process, info.UUID, info.InfobaseID)

				mu.Lock()
				respond.AddResult(models.ConnectionSig{
					ClusterID: info.ClusterID, Process: info.Process, UUID: info.UUID,
				}, err)
				mu.Unlock()
				wg.Done()
			}(info)

		})

		wg.Wait()

	case len(body.Connections) > 0:

		EachConnection(body.Connections).Do(func(info models.ConnectionSig) {
			err := client.DisconnectConnection(info.ClusterID, info.Process, info.UUID, info.InfobaseID)
			respond.AddResult(models.ConnectionSig{
				ClusterID: info.ClusterID, Process: info.Process, UUID: info.UUID,
			}, err)
		})

	default:
		return ErrorResponse(ctx, errors.New("InfobaseID or Connections must be set"))
	}

	return SuccessResponse(ctx, respond)

}

// Terminate отключение списка подключений на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary отключение списка подключений на сервере 1С Предприятие
// 		@Description отключение списка подключений или по информационной базе на сервере 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param req body models.TerminateConnectionsRequest true "request"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
//		@Success 200 {object} Response{data=models.TerminateConnectionsResponse}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/connections/terminate [post]
func (a *connectionsApi) TerminateApp() {}

type EachConnection []models.ConnectionSig

func (l EachConnection) Do(fn func(info models.ConnectionSig)) {
	for _, sig := range l {
		fn(sig)
	}
}

func filterServiceConnections(info *serialize.ConnectionShortInfo) bool {
	switch info.Application {
	case "AgentStandardCall", "JobScheduler":
		return false
	default:
		return true
	}
}
