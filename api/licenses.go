package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/ras-client/serialize"
)

type licensesApi struct {
	service service.Service
}

func (a *licensesApi) Routes(r fiber.Router) {

	r.Get("/licenses", withClient(a.List))

}

// List получение списка лицензий на кластере
//  Swagger-spec:
//		@Summary получение списка лицензий на кластере
// 		@Description получение списка лицензий на кластере
// 		@Tags licenses
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.LicenseInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/licenses [get]
func (a *licensesApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	licenses := serialize.LicenseInfoList{}

	list, err := a.service.GetSessions(client)

	if err != nil {
		return ErrorResponse(ctx, err, "error get licenses")
	}

	list.Each(func(info *serialize.SessionInfo) {
		if info.Licenses == nil {
			return
		}
		licenses = append(licenses, *info.Licenses...)
	})

	return SuccessResponse(ctx, licenses)

}

// List получение списка лицензий на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка лицензий на сервере 1С Предприятие
// 		@Description получение списка лицензий на сервере 1С Предприятие
// 		@Tags licenses
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.LicenseInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/licenses [get]
func (a *licensesApi) ListApp() {}
