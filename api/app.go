package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/odin/service"
)

type AppApi struct {
	service service.Service
}

func (a *AppApi) Routes(r fiber.Router) {

	r.Get("/app", a.List)
	r.Post("/app", a.Create)

	r.Get("/app/:app", a.Info)
	r.Post("/app/:app", a.Update)
	r.Delete("/app/:app", a.Delete)

	router := r.Group("/app/:app")

	parentApi := []route{
		&clusterApi{service: a.service},
		&infobasesApi{service: a.service},
		&sessionsApi{service: a.service},
		&connectionsApi{service: a.service},
		&managersApi{service: a.service},
		&servicesApi{service: a.service},
		&locksApi{service: a.service},
		&blockerApi{service: a.service},
		&processesApi{service: a.service},
		&licensesApi{service: a.service},
		&healthAppApi{service: a.service},
		&agentApi{service: a.service},
	}

	for _, api := range parentApi {
		api.Routes(router)
	}

}

// List Получение списка зарегистрированных серверов 1С.Предприятие
//  Swagger-spec:
//		@Summary Получение списка зарегистрированных серверов 1С.Предприятие
// 		@Description Получение списка зарегистрированных серверов 1С.Предприятие
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} Response{data=[]models.AppServer}
//		@Failure 400 {object} Response
// 		@Router /app [get]
func (a *AppApi) List(ctx *fiber.Ctx) error {

	apps, err := a.service.GetAppServers()

	if err != nil {
		return ErrorResponse(ctx, err, "error get app list")
	}

	return SuccessResponse(ctx, apps)
}

// Info получение информации о зарегистрированном сервере 1С.Предприятие
//  Swagger-spec:
//		@Summary Получение информации о зарегистрированном сервере 1С.Предприятие
// 		@Description Получение информации о зарегистрированном сервере 1С.Предприятие
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Success 200 {object} Response{data=models.AppServer}
//		@Failure 400 {object} Response
// 		@Router /app/{app} [get]
func (a *AppApi) Info(ctx *fiber.Ctx) error {

	name := ctx.Params("app")

	app, err := a.service.GetAppServer(name)

	if err != nil {
		return ErrorResponse(ctx, err, "error get app info")
	}

	return SuccessResponse(ctx, app)

}

// Update обновление информации о зарегистрированном сервере 1С.Предприятие
//  Swagger-spec:
//		@Summary Обновление информации о зарегистрированном сервере 1С.Предприятие
// 		@Description Обновление информации о зарегистрированном сервере 1С.Предприятие
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param req body models.AppServer true "app info"
// 		@Success 200 {object} Response{data=models.AppServer}
//		@Failure 400 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app/{app} [post]
func (a *AppApi) Update(ctx *fiber.Ctx) error {
	var newApp models.AppServer

	err := ctx.BodyParser(&newApp)

	if err != nil {
		return ErrorResponse(ctx, err, "error parse body to app info")
	}

	name := ctx.Params("app")

	newApp.Name = name
	err = a.service.SetAppServer(&newApp)

	return HttpResponse(ctx, newApp, err, "error update app")

}

// Create выполняет регистрацию сервера 1С.Предприятие в приложении
//  Swagger-spec:
//		@Summary выполняет регистрацию сервера 1С.Предприятие в приложении
// 		@Description выполняет регистрацию сервера 1С.Предприятие в приложении
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Param req body models.AppServer true "app info"
// 		@Success 200 {object} Response{data=models.AppServer}
//		@Failure 400 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app [post]
func (a *AppApi) Create(ctx *fiber.Ctx) error {

	var newApp models.AppServer

	err := ctx.BodyParser(&newApp)

	if err != nil {
		return ErrorResponse(ctx, err, "error parse body to app info")
	}
	err = a.service.AddAppServer(&newApp)

	if err != nil {
		return ErrorResponse(ctx, err, "error registry app")
	}

	return SuccessResponse(ctx, newApp)
}

// Delete выполняет отмену зарегистрирации сервера 1С.Предприятие
//  Swagger-spec:
//		@Summary Удаление информации о регистрации сервера приложений 1С.Предприятие
// 		@Description Удаление информации о регистрации сервера приложений 1С.Предприятие
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Success 200 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app/{app} [delete]
func (a *AppApi) Delete(ctx *fiber.Ctx) error {

	name := ctx.Params("app")

	err := a.service.DeleteAppServer(name)

	if err != nil {
		return ErrorResponse(ctx, err, "error delete app")
	}

	return SuccessResponse(ctx, "deleted")
}
