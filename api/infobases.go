package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/ras-client/serialize"
	"strconv"
)

type infobasesApi struct {
	service service.Service
}

func (a *infobasesApi) Routes(r fiber.Router) {

	r.Get("/infobases", withClient(a.List))
	r.Post("/infobases", withClient(a.Create))
	r.Get("/infobases/:infobase", withClient(a.Info))
	r.Post("/infobases/:infobase", withClient(a.Update))
	r.Delete("/infobases/:infobase", withClient(a.Drop))

	router := r.Group("/infobases/:infobase")

	parentApi := []route{
		&sessionsApi{service: a.service},
		&connectionsApi{service: a.service},
		&blockerApi{service: a.service},
		&locksApi{service: a.service},
	}

	for _, api := range parentApi {
		api.Routes(router)
	}

}

// List получение списка информационных баз с кластера
//  Swagger-spec:
//		@Summary получение списка информационных баз с кластера
// 		@Description получение списка информационных баз с кластера
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseSummaryList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases [get]
func (a *infobasesApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetInfobases(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get infobases")

	}
	return SuccessResponse(ctx, list)

}

// List получение списка информационных баз с сервера 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка информационных баз с сервера 1С Предприятие
// 		@Description получение списка информационных баз с сервера 1С Предприятие
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseSummaryList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases [get]
func (a *infobasesApi) ListApp() {}

// Info получение информации об информационной базе с кластера
//  Swagger-spec:
//		@Summary получение информации об информационной базе с кластера
// 		@Description получение информации об информационной базе с кластера
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase name or uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.InfobaseInfo}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase} [get]
func (a *infobasesApi) Info(client service.ClientContext, ctx *fiber.Ctx) error {

	info, err := a.service.GetInfobase(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get infobase info")

	}
	return SuccessResponse(ctx, info)

}

// Info получение информации об информационной базе с сервера 1С Предприятие
//  Swagger-spec:
//		@Summary получение информации об информационной базе с сервера 1С Предприятие
// 		@Description получение информации об информационной базе с сервера 1С Предприятие
// 		@Tags infobases
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
// 		@Success 200 {object} Response{data=serialize.InfobaseInfo}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/infobases/{infobase} [get]
func (a *infobasesApi) InfoApp() {}

// Update обновление информации об информационной базе на кластере
//  Swagger-spec:
//		@Summary обновление информации об информационной базе на кластере
// 		@Description побновление информации об информационной базе на кластере
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase name or uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
// 		@Param info body serialize.InfobaseInfo true "new info"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 400 {object} Response{data=string}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase} [post]
func (a *infobasesApi) Update(client service.ClientContext, ctx *fiber.Ctx) error {

	var body serialize.InfobaseInfo

	err := ctx.BodyParser(&body)
	if err != nil {
		return ErrorResponse(ctx, err, "error body parse infobase info")
	}

	err = validateUpdateInfobaseInfo(body)
	if err != nil {
		return ErrorResponse(ctx, err, "validate error")
	}

	info, err := a.service.UpdateInfobase(client, &body)

	if err != nil {
		return ErrorResponse(ctx, err, "error update infobase")

	}
	return SuccessResponse(ctx, info)

}

// Update обновление информации об информационной базе на сервера 1С Предприятие
//  Swagger-spec:
//		@Summary обновление информации об информационной базе на сервера 1С Предприятие
// 		@Description обновление информации об информационной базе на сервера 1С Предприятие
// 		@Tags infobases
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
// 		@Param info body serialize.InfobaseInfo true "new info"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 400 {object} Response{data=string}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/infobases/{infobase} [post]
func (a *infobasesApi) UpdateApp() {}

func validateUpdateInfobaseInfo(body serialize.InfobaseInfo) error {

	// TODO Сделать валидатор информации об информационной базе

	return nil

}

// Create создание информационной базы на кластере
//  Swagger-spec:
//		@Summary оздание информационной базы на кластере
// 		@Description оздание информационной базы на кластере
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param create-db query bool false "create server db"
// 		@Param info body serialize.InfobaseInfo true "new info"
// 		@Success 200 {object} Response{data=serialize.InfobaseInfo}
// 		@Failure 400 {object} Response{data=string}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/clusters/{cluster}/infobases [post]
func (a *infobasesApi) Create(client service.ClientContext, ctx *fiber.Ctx) error {

	var body serialize.InfobaseInfo

	err := ctx.BodyParser(&body)
	if err != nil {
		return ErrorResponse(ctx, err, "error body parse infobase info")
	}

	err = validateCreateInfobaseInfo(body)
	if err != nil {
		return ErrorResponse(ctx, err, "validate error")
	}

	createDB, _ := strconv.ParseBool(ctx.Query("create-db", "false"))

	info, err := a.service.CreateInfobase(client, &body, createDB)

	if err != nil {
		return ErrorResponse(ctx, err, "create infobase error")

	}
	return SuccessResponse(ctx, info)
}

// Create создание информационной базы на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary создание информационной базы на сервере 1С Предприятие
// 		@Description создание информационной базы на сервере 1С Предприятие
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param create-db query bool false "create server db"
// 		@Param info body serialize.InfobaseInfo true "new info"
// 		@Success 200 {object} Response{data=serialize.InfobaseInfo}
// 		@Failure 400 {object} Response{data=string}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/infobases [post]
func (a *infobasesApi) CreateApp() {}

func validateCreateInfobaseInfo(body serialize.InfobaseInfo) error {
	// TODO Сделать валидатор информации об информационной базе
	return nil
}

// Drop удаляет информационную базу с сервера кластера
//  Swagger-spec:
//		@Summary удаляет информационную базу с сервера кластера
// 		@Description удаляет информационную базу с сервера кластера
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase name or uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
//		@Param delete-db query bool false "delete server db"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase} [delete]
func (a *infobasesApi) Drop(client service.ClientContext, ctx *fiber.Ctx) error {

	deleteDB, _ := strconv.ParseBool(ctx.Query("delete-db", "false"))
	err := a.service.DropInfobase(client, deleteDB)

	if err != nil {
		return ErrorResponse(ctx, err, "drop infobase error")

	}
	return SuccessResponse(ctx, nil)
}

// DropApp удаляет информационную базу с сервера 1С Предприятие
//  Swagger-spec:
//		@Summary удаляет информационную базу с сервера 1С Предприятие
// 		@Description удаляет информационную базу с сервера 1С Предприятие
// 		@Tags infobases
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase name or uuid"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
//		@Param delete-db query bool false "delete server db"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response{data=string}
// 		@Router /app/{app}/infobases/{infobase} [delete]
func (a *infobasesApi) DropApp() {}
