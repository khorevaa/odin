package api

import (
	"github.com/khorevaa/odin/models"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type servicesSuite struct {
	baseSuite
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(servicesSuite))
}

func (s *servicesSuite) SetupSuite() {
	s.api().
		Post("/api/v1/app").
		JSON(&models.AppServer{
			Name: "test",
			Addr: "localhost",
			Port: "1546",
		}).
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).End()
}

func (s *servicesSuite) TestServiceList() {
	s.api().
		Debug().Report(apitest.SequenceDiagram()).
		Get("/api/v1/app/test/services").
		Expect(s.T()).
		Assert(jsonpath.Equal(`$.message`, "success")).
		Status(http.StatusOK).
		End()
}
