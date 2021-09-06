package api

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/khorevaa/odin/database"
	"github.com/khorevaa/odin/ras"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/odin/service/cache"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"sync"
	"time"
)

var once = sync.Once{}

type baseSuite struct {
	suite.Suite
}

func (s *baseSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseSuite) api(recorder ...*apitest.Recorder) *apitest.APITest {

	rec := apitest.NewTestRecorder()
	if len(recorder) > 0 {
		rec = recorder[0]
	}
	once.Do(func() {

	})

	return apitest.New().
		HandlerFunc(FiberToHandlerFunc(newTestApp(rec))).
		//Debug().
		//Report(apitest.SequenceDiagram())
		Recorder(rec)

}

func newTestApp(rec *apitest.Recorder) *fiber.App {

	server := fiber.New()

	memoryCache := &cache.Memory{
		Expiration: 30 * time.Minute,
	}

	memoryCache.Connect()

	ras.SetLocalStorage(RecorderStorage(rec))

	rep := WithRecorderRepository(db.NewRepository("./tests/"), rec)
	s, _ := service.NewService(WithRecorderCache(memoryCache, rec), rep)
	serv := WithRecorderService(s, rec)
	server.Use(service.Middleware(serv))

	Routes(server, serv)
	return server

}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
