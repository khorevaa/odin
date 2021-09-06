package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/odin/service"
)

type blockerApi struct {
	service service.Service
}

func (a *blockerApi) Routes(r fiber.Router) {

	r.Get("/block", withClient(a.GetBlock))
	r.Post("/block", withClient(a.PostBlock))

	r.Get("/unblock", withClient(a.GetUnblock))
	r.Post("/unblock", withClient(a.PostUnblock))

}

// Block установка блокировки на информационную базу на кластере
//  Swagger-spec:
//		@Summary Установка блокировки на информационную базу на кластере
// 		@Description установка блокировки на информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase query string true "infobase uuid or name"
// 		@Param sessions-deny query string true "session deny"
// 		@Param message query string false "message to user"
// 		@Param permission-code query string false "permission code"
// 		@Param denied-parameter query string false "denied parameter"
// 		@Param permission-code query string false "permission code"
// 		@Param scheduled-jobs-deny query bool false "scheduled jobs deny"
// 		@Param denied-from query string false "denied from time"
// 		@Param denied-to query string false "denied to time"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/block [get]
func (a *blockerApi) GetBlock(client service.ClientContext, ctx *fiber.Ctx) error {

	var blocker models.InfobaseBlocker

	err := ctx.QueryParser(&blocker)

	if err != nil {
		return ErrorResponse(ctx, errors.BadRequest.Wrap(err, "bad request"), "error parse blocker query")
	}

	unblocker, err := a.service.Block(client, &blocker)

	if err != nil {
		return ErrorResponse(ctx, err, "block error")
	}

	return SuccessResponse(ctx, unblocker)

}

// Block Установка блокировки конкретную на информационную базу на кластере
//  Swagger-spec:
//		@Summary Установка блокировки на выбранную информационную базу на кластере
// 		@Description Установка блокировки на выбранную информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param sessions-deny query string true "session deny"
// 		@Param message query string false "message to user"
// 		@Param permission-code query string false "permission code"
// 		@Param denied-parameter query string false "denied parameter"
// 		@Param permission-code query string false "permission code"
// 		@Param scheduled-jobs-deny query bool false "scheduled jobs deny"
// 		@Param denied-from query string false "denied from time"
// 		@Param denied-to query string false "denied to time"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase}/block [get]
func (a *blockerApi) GetBlockInfobase() {}

// Block Установка блокировки на произвольную информационную базу на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary Установка блокировки на произвольную информационную базу на сервер 1С Предприятие
// 		@Description Установка блокировки на произвольную информационную базу на сервер 1С Предприятие
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase-id query string true "infobase uuid or name"
// 		@Param sessions-deny query string true "session deny"
// 		@Param message query string false "message to user"
// 		@Param permission-code query string false "permission code"
// 		@Param denied-parameter query string false "denied parameter"
// 		@Param permission-code query string false "permission code"
// 		@Param scheduled-jobs-deny query bool false "scheduled jobs deny"
// 		@Param denied-from query string false "denied from time" default("now")
// 		@Param denied-to query string false "denied to time"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/block [get]
func (a *blockerApi) GetlockApp() {}

// Block Установка блокировки конкретную на информационную базу на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary Установка блокировки конкретную на информационную базу на сервер 1С Предприятие
// 		@Description Установка блокировки конкретную на информационную базу на сервер 1С Предприятие
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param sessions-deny query string true "session deny"
// 		@Param message query string false "message to user"
// 		@Param permission-code query string false "permission code"
// 		@Param denied-parameter query string false "denied parameter"
// 		@Param permission-code query string false "permission code"
// 		@Param scheduled-jobs-deny query bool false "scheduled jobs deny"
// 		@Param denied-from query string false "denied from time"
// 		@Param denied-to query string false "denied to time"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases/{infobase}/block [get]
func (a *blockerApi) GetBlockAppInfobase() {}

// Block установка блокировки на информационную базу на кластере
//  Swagger-spec:
//		@Summary Установка блокировки на информационную базу на кластере
// 		@Description установка блокировки на информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param body body models.InfobaseBlocker true "block info"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/block [post]
func (a *blockerApi) PostBlock(client service.ClientContext, ctx *fiber.Ctx) error {

	var blocker models.InfobaseBlocker

	err := ctx.BodyParser(&blocker)

	if err != nil {
		return ErrorResponse(ctx, err, "error parse blocker body")
	}

	unblocker, err := a.service.Block(client, &blocker)

	if err != nil {
		return ErrorResponse(ctx, err, "block error")
	}

	return SuccessResponse(ctx, unblocker)

}

// Block установка блокировки на конкретную информационную базу на кластере
//  Swagger-spec:
//		@Summary Установка блокировки на конкретную информационную базу на кластере
// 		@Description установка блокировки на конкретную информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param body body models.InfobaseBlocker true "block info"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase}/block [post]
func (a *blockerApi) PostBlockInfobase() {}

// Block Установка блокировки на произвольную информационную базу на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary Установка блокировки на произвольную информационную базу на сервер 1С Предприятие
// 		@Description Установка блокировки на произвольную информационную базу на сервер 1С Предприятие
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param body body models.InfobaseBlocker true "block info"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/block [post]
func (a *blockerApi) PostBlockApp() {}

// Block Установка блокировки конкретную на информационную базу на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary Установка блокировки конкретную на информационную базу на сервер 1С Предприятие
// 		@Description Установка блокировки конкретную на информационную базу на сервер 1С Предприятие
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param body body models.InfobaseBlocker true "block info"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=models.InfobaseUnblocker}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases/{infobase}/block [post]
func (a *blockerApi) PostBlockAppInfobase() {}

// GetUnblock Снятие блокировки на информационную базу на кластере
//  Swagger-spec:
//		@Summary Снятие блокировки на информационную базу на кластере
// 		@Description Снятие блокировки на информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase query string true "infobase uuid or name"
// 		@Param sessions-deny query string true "session deny"
// 		@Param denied-parameter query string false "denied parameter"
// 		@Param permission-code query string false "permission code"
// 		@Param scheduled-jobs-deny query bool false "scheduled jobs deny"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseSummaryInfo}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/unblock [get]
func (a *blockerApi) GetUnblock(client service.ClientContext, ctx *fiber.Ctx) error {

	var unblocker models.InfobaseUnblocker

	err := ctx.QueryParser(&unblocker)

	if err != nil {
		return ErrorResponse(ctx, err, "error parse unblocker query")
	}

	info, err := a.service.Unblock(client, &unblocker)

	if err != nil {
		return ErrorResponse(ctx, err, "unblock error")
	}

	return SuccessResponse(ctx, info)
}

// GetUnblock Снятие блокировки на информационную базу на кластере
//  Swagger-spec:
//		@Summary Снятие блокировки на информационную базу на кластере
// 		@Description Снятие блокировки на информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase query string true "infobase uuid or name"
// 		@Param sessions-deny query string true "session deny"
// 		@Param denied-parameter query string false "denied parameter"
// 		@Param permission-code query string false "permission code"
// 		@Param scheduled-jobs-deny query bool false "scheduled jobs deny"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseSummaryInfo}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/unblock [get]
func (a *blockerApi) GetUnblockApp() {}

// PostUnblock Снимает блокировку на информационную базу на кластере
//  Swagger-spec:
//		@Summary Снимает блокировку на информационную базу на кластере
// 		@Description Снимает блокировку на информационную базу на кластере
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param body body models.InfobaseUnblocker true "unblock info"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseSummaryInfo}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/unblock [post]
func (a *blockerApi) PostUnblock(client service.ClientContext, ctx *fiber.Ctx) error {

	var unblocker models.InfobaseUnblocker

	err := ctx.BodyParser(&unblocker)

	if err != nil {
		return ErrorResponse(ctx, err, "error parse unblocker query")
	}

	info, err := a.service.Unblock(client, &unblocker)

	if err != nil {
		return ErrorResponse(ctx, err, "unblock error")
	}

	return SuccessResponse(ctx, info)
}

// PostUnblockApp Снимает блокировку на информационную базу на сервере 1С Прдприятие
//  Swagger-spec:
//		@Summary Снимает блокировку на информационную базу на сервере 1С Прдприятие
// 		@Description Снимает блокировку на информационную базу на сервере 1С Прдприятие
// 		@Tags blocker
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param body body models.InfobaseUnblocker true "unblock info"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseSummaryInfo}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/unblock [post]
func (a *blockerApi) PostUnblockApp() {}
