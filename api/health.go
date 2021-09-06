package api

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/odin/service"
	"net"
	"sync"
	"time"
)

type healthApi struct {
	name    string
	version string
	route   string
	service service.Service
}

func (a *healthApi) Routes(r fiber.Router) {

	r.Get("/health", a.health)
	r.Get("/health/readiness", a.readiness)
}

// health запрос о состонии приложения
//  Swagger-spec:
//		@Summary запрос о состонии приложения
// 		@Description запрос о состонии приложения
// 		@Tags health
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} StatusResponse
//		@Failure 500 {object} StatusResponse
// 		@Router /health [get]
func (a *healthApi) health(ctx *fiber.Ctx) error {

	ok, err := a.service.HealthCheck()

	if ok && err == nil {
		return ctx.Status(fiber.StatusOK).JSON(StatusResponse{
			Status: true,
		})
	}

	return ctx.Status(fiber.StatusServiceUnavailable).JSON(StatusResponse{
		Status: false,
		Err:    err.Error(),
	})

}

// readiness запрос подробного состония приложения
//  Swagger-spec:
//		@Summary запрос подробного состония приложения
// 		@Description запрос подробного состония приложения
// 		@Tags health
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} ReadinessCheckStatus
//		@Failure 500 {object} ReadinessCheckStatus
// 		@Router /health/readiness [get]
func (a *healthApi) readiness(ctx *fiber.Ctx) error {

	var check readinessCheck
	check.Name = a.name
	check.Version = a.version

	apps, _ := a.service.GetAppServers()

	host := fmt.Sprintf("%s://%s%s", ctx.Protocol(), ctx.Hostname(), a.route)

	for _, app := range apps {

		check.Apps = append(check.Apps, AppServiceCheckConfig{
			App:     app,
			Host:    host,
			TimeOut: 10,
			Ctx:     ctx,
		})

	}

	result := check.Check()

	if result.Status {
		return ctx.Status(fiber.StatusOK).JSON(result)
	}

	return ctx.Status(fiber.StatusServiceUnavailable).JSON(result)

}

type readinessCheck struct {
	Name    string
	Version string
	service service.Service
	Apps    []AppServiceCheckConfig
}

type ReadinessCheckStatus struct {
	Name    string             `json:"name"`
	Version string             `json:"version"`
	Status  bool               `json:"status"`
	Apps    []AppServiceStatus `json:"apps"`
}

func (c readinessCheck) Check() ReadinessCheckStatus {

	check := ReadinessCheckStatus{
		Name:    c.Name,
		Version: c.Version,
		Status:  true,
	}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, app := range c.Apps {
		wg.Add(1)

		go func(a AppServiceCheckConfig) {
			res := checkAppServiceClient(a)
			mu.Lock()
			check.Apps = append(check.Apps, res)
			mu.Unlock()
			wg.Done()
		}(app)

	}

	wg.Wait()

	for _, app := range check.Apps {
		if !app.Status {
			check.Status = false
			break
		}
	}

	return check

}

type serviceStatus struct {
	Status bool   `json:"status"`
	Error  string `json:"errors,omitempty"`
}

type AppServiceCheckConfig struct {
	App     *models.AppServer
	Host    string
	TimeOut time.Duration `json:"timeout,omitempty"` // default value: 10
	Headers []HTTPHeader  `json:"headers,omitempty"`
	Ctx     *fiber.Ctx
}

type AppServiceStatus struct {
	Name         string  `json:"name"`
	Host         string  `json:"host"`
	Status       bool    `json:"status"`
	ResponseTime float64 `json:"response_time"`
	URL          string  `json:"url"`
	Error        string  `json:"errors,omitempty"`
}

// HTTPHeader used to setup webservices integrations
type HTTPHeader struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"Value,omitempty"`
}

func checkAppServiceHttpClient(config AppServiceCheckConfig) AppServiceStatus {

	var host = fmt.Sprintf("%s/app/%s/health", config.Host, config.App.Name)
	var timeout = time.Second * 10
	var myStatus = true

	if config.TimeOut > 0 {
		timeout = time.Second * config.TimeOut
	}

	start := time.Now()
	req := requests.Requests()
	req.SetTimeout(timeout)
	if len(config.Headers) > 0 {
		for _, v := range config.Headers {
			req.Header.Set(v.Key, v.Value)
		}
	}

	resp, err := req.Get(host)
	var checkError string
	switch {
	case err != nil:
		myStatus = false
		checkError = err.Error()
	case resp.R.StatusCode == fiber.StatusServiceUnavailable:

		myStatus = false
		statusErr := StatusResponse{}
		_ = resp.Json(&statusErr)

		checkError = statusErr.Err

	case resp.R.StatusCode == 200:
		// nothing do
	}

	elapsed := time.Now().Sub(start)
	return AppServiceStatus{
		Name:         config.App.Name,
		Host:         net.JoinHostPort(config.App.Addr, config.App.Port),
		Status:       myStatus,
		ResponseTime: elapsed.Seconds(),
		URL:          host,
		Error:        checkError,
	}

}

func checkAppServiceClient(config AppServiceCheckConfig) AppServiceStatus {

	var host = fmt.Sprintf("%s/app/%s/health", config.Host, config.App.Name)

	start := time.Now()

	client := service.NewClientContext(config.App, config.Ctx)

	ok, err := client.HealthCheck()
	checkError := ""
	if err != nil {
		checkError = err.Error()
	}

	elapsed := time.Now().Sub(start)
	return AppServiceStatus{
		Name:         config.App.Name,
		Host:         net.JoinHostPort(config.App.Addr, config.App.Port),
		Status:       ok,
		ResponseTime: elapsed.Seconds(),
		URL:          host,
		Error:        checkError,
	}

}

type StatusResponse struct {
	Status bool   `json:"status"`
	Err    string `json:"errors,omitempty"`
}

func (r StatusResponse) Error() string {
	return r.Err
}
