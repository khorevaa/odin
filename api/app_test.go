package api

import (
	"github.com/khorevaa/odin/models"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type appSuite struct {
	baseSuite
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(appSuite))
}

func (s *appSuite) TestAppList() {

	s.api().
		Post("/api/v1/app").
		JSON(&models.AppServer{
			Name: "test",
			Addr: "localhost",
			Port: "1545",
		}).
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()

	s.api().Report(apitest.SequenceDiagram()).
		Get("/api/v1/app").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).
		End()
}

func (s *appSuite) TestAppReg() {
	s.api().
		Post("/api/v1/app").
		JSON(&models.AppServer{
			Name: "test",
			Addr: "localhost",
			Port: "1545",
		}).
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()
}

func (s *appSuite) TestAppUnreg() {
	s.api().
		Postf("/api/v1/app/%s", "test").
		JSON(&models.AppServer{
			Name: "test",
			Addr: "localhost",
			Port: "1546",
		}).
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).
		End()

	s.api().
		Deletef("/api/v1/app/%s", "test").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()

}

func (s *appSuite) TestHealthStatus() {
	s.api().
		Get("/api/v1/health").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.status`, true)).
		Status(http.StatusOK).
		End()
}

func (s *appSuite) TestHealthReadinessStatus() {
	s.api().
		Postf("/api/v1/app/%s", "test").
		JSON(&models.AppServer{
			Name: "test",
			Addr: "localhost",
			Port: "1546",
		}).
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()

	s.api().
		Get("/api/v1/health/readiness").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.status`, true)).
		Status(http.StatusOK).
		End()

	s.api().
		Deletef("/api/v1/app/%s", "test").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()

}

func (s *appSuite) TestHealthReadinessStatusBad() {

	s.api().
		Postf("/api/v1/app/%s", "test").
		JSON(&models.AppServer{
			Name: "test",
			Addr: "localhost",
			Port: "1545",
		}).
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()

	s.api().
		Get("/api/v1/health/readiness").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.status`, false)).
		Status(http.StatusServiceUnavailable).
		End()

}
