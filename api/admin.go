package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/ras-client/serialize"
)

type agentApi struct {
	service service.Service
}

func (a *agentApi) Routes(r fiber.Router) {

	r.Get("/agent/version", withClient(a.Version))
	r.Get("/agent/admins", withClient(a.List))
	r.Post("/agent/admins", withClient(a.RegAgentAdmin))
	r.Delete("/agent/admins/:admin", withClient(a.UnregAgentAdmin))

}

// List получение списка администраторов агента на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка администраторов агента на сервере 1С Предприятие
// 		@Description получение списка администраторов агента на сервере 1С Предприятие
// 		@Tags admins,agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.UsersList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/admins [get]
func (a *agentApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetAgentAdmins(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get agent admins")
	}

	return SuccessResponse(ctx, list)
}

// Version получение версии агента на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение версии агента на сервере 1С Предприятие
// 		@Description получение версии агента на сервере 1С Предприятие
// 		@Tags agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/version [get]
func (a *agentApi) Version(client service.ClientContext, ctx *fiber.Ctx) error {
	version, err := a.service.GetAgentVersion(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get agent version")
	}

	return SuccessResponse(ctx, version)
}

// RegAgentAdmin выполняет регистрацию нового адмнимистратор на агенте сервера 1С Предприятие
//  Swagger-spec:
//		@Summary выполняет регистрацию нового адмнимистратор на агенте сервера 1С Предприятиеи
// 		@Description выполняет регистрацию нового адмнимистратор на агенте сервера 1С Предприятие
// 		@Tags admins,agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param req body serialize.UserInfo true "user info"
// 		@Success 200 {object} Response{data=serialize.UserInfo}
//		@Failure 404 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/admins [post]
func (a *agentApi) RegAgentAdmin(client service.ClientContext, ctx *fiber.Ctx) error {

	var info serialize.UserInfo

	err := ctx.BodyParser(&info)

	if err != nil {
		return ErrorResponse(ctx, errors.BadRequest.Wrap(err, "parse user info"), "error parse body to user info")
	}

	err = a.service.RegAgentAdmin(client, info)

	if err != nil {
		return ErrorResponse(ctx, err, "error reg agent admin")
	}

	return SuccessResponse(ctx, info)

}

// UnregAgentAdmin Удаление администратора агента на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary Удаление администратора агента на сервере 1С Предприятие
// 		@Description Удаление администратора агента на сервере 1С Предприятие
// 		@Tags admins,agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param admin path string true "admin name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/admins/{admin} [delete]
func (a *agentApi) UnregAgentAdmin(client service.ClientContext, ctx *fiber.Ctx) error {

	name := ctx.Params("admin")

	err := a.service.UnregAgentAdmin(client, name)

	if err != nil {
		return ErrorResponse(ctx, err, "error unreg agent admin")
	}

	return SuccessResponse(ctx, nil)

}
