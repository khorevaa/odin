package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/ras-client/serialize"
)

type clusterAdminApi struct {
	service service.Service
}

func (a *clusterAdminApi) Routes(r fiber.Router) {

	r.Get("/admins", withClient(a.List))
	r.Post("/admins", withClient(a.Create))
	r.Delete("/admins/:admin", withClient(a.Delete))

}

// List получение списка администраторов кластера
//  Swagger-spec:
//		@Summary получение списка администраторов кластера
// 		@Description получение списка администраторов кластера
// 		@Tags admins,clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.UsersList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/cluster/{cluster}/admins [get]
func (a *clusterAdminApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	list, err := a.service.GetClusterAdmins(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get cluster admins")
	}

	return SuccessResponse(ctx, list)
}

// Create выполняет регистрацию нового администратор на кластере
//  Swagger-spec:
//		@Summary выполняет регистрацию нового администратор на кластере
// 		@Description выполняет регистрацию нового администратор на кластере
// 		@Tags admins,clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param req body serialize.UserInfo true "user info"
// 		@Success 200 {object} Response{data=serialize.UserInfo}
//		@Failure 404 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app/{app}/cluster/{cluster}/admins [post]
func (a *clusterAdminApi) Create(client service.ClientContext, ctx *fiber.Ctx) error {

	var info serialize.UserInfo

	err := ctx.BodyParser(&info)

	if err != nil {
		return ErrorResponse(ctx, errors.BadRequest.Wrap(err, "parse user info"), "error parse body to user info")
	}

	_, err = a.service.RegClusterAdmin(client, info)

	if err != nil {
		return ErrorResponse(ctx, err, "error reg cluster admin")
	}

	return SuccessResponse(ctx, info)

}

// Delete Удаление администратора агента на кластере
//  Swagger-spec:
//		@Summary Удаление администратора агента на кластере
// 		@Description Удаление администратора агента на кластере
// 		@Tags admins,clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param admin path string true "admin name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/cluster/{cluster}/admins/{admin} [delete]
func (a *clusterAdminApi) Delete(client service.ClientContext, ctx *fiber.Ctx) error {

	name := ctx.Params("admin")

	err := a.service.UnregClusterAdmin(client, name)

	if err != nil {
		return ErrorResponse(ctx, err, "error unreg cluster admin")
	}

	return SuccessResponse(ctx, nil)

}
