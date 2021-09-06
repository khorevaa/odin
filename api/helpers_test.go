package api

import (
	"context"
	"encoding/json"
	db "github.com/khorevaa/odin/database"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/odin/ras"
	"github.com/khorevaa/odin/service"
	"github.com/khorevaa/odin/service/cache"
	rclient "github.com/khorevaa/ras-client"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
	"github.com/steinfletcher/apitest"
	"time"
)

// WithRecorderRepository wraps an existing driver with a Recorder
func WithRecorderRepository(repository db.Repository, recorder *apitest.Recorder) db.Repository {

	recordingDriver := &recordingRepository{
		sourceName: "pudge DB",
		repository: repository,
		recorder:   recorder,
	}

	return recordingDriver
}

// WithRecorderService wraps an existing driver with a Recorder
func WithRecorderService(serv service.Service, recorder *apitest.Recorder) service.Service {

	recordingDriver := &recordingService{
		sourceName: "service",
		service:    serv,
		recorder:   recorder,
	}

	return recordingDriver
}

// WithRecorderService wraps an existing driver with a Recorder
func WithRecorderCache(c cache.Cache, recorder *apitest.Recorder) cache.Cache {

	recordingDriver := &recordingCache{
		sourceName: "cache",
		c:          c,
		recorder:   recorder,
	}

	return recordingDriver
}

// WithRecorderService wraps an existing driver with a Recorder
func WithRecorderClient(c rclient.Api, recorder *apitest.Recorder) rclient.Api {

	recordingDriver := &recordingClient{
		sourceName: "ras",
		Api:        c,
		recorder:   recorder,
	}

	return recordingDriver
}

// WithRecorderService wraps an existing driver with a Recorder
func RecorderStorage(recorder *apitest.Recorder) ras.Storage {

	recordingDriver := &recordingStorage{
		sourceName: "storage",
		recorder:   recorder,
	}

	return recordingDriver
}

type recordingStorage struct {
	recorder   *apitest.Recorder
	sourceName string
}

func (r recordingStorage) Store(id string, addr string, version string) (rclient.Api, bool) {

	client := WithRecorderClient(rclient.NewClient(addr, rclient.WithVersion(version)), r.recorder)
	return client, false
}

func (r recordingStorage) LoadOrInit(id string, addr string, version string) (rclient.Api, bool) {
	client := WithRecorderClient(rclient.NewClient(addr, rclient.WithVersion(version)), r.recorder)
	return client, false
}

func (r recordingStorage) Load(id string) (interface{}, bool) {

	return nil, false
}

type recordingClient struct {
	rclient.Api
	recorder   *apitest.Recorder
	sourceName string
}

func (r recordingClient) addReq(header string, body ...interface{}) {

	var data string

	if len(body) > 0 {
		data = toJson(body[0])
	}

	r.recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    "service",
		Target:    r.sourceName,
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingClient) addResp(header string, body interface{}, err error) {

	var data string

	if err != nil {
		data = err.Error()
	} else {
		data = toJson(body)
	}

	r.recorder.AddMessageResponse(apitest.MessageResponse{
		Source:    r.sourceName,
		Target:    "service",
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingClient) GetClusters(ctx context.Context) ([]*serialize.ClusterInfo, error) {

	r.addReq("GetClusters")
	ok, err := r.Api.GetClusters(ctx)
	r.addResp("GetClusters", ok, err)

	return ok, err

}

func (r recordingClient) GetClusterServices(ctx context.Context, cluster uuid.UUID) ([]*serialize.ServiceInfo, error) {

	r.addReq("GetClusterServices",
		map[string]interface{}{"cluster": cluster})
	ok, err := r.Api.GetClusterServices(ctx, cluster)
	r.addResp("GetClusterServices", ok, err)

	return ok, err

}

type recordingCache struct {
	c          cache.Cache
	recorder   *apitest.Recorder
	sourceName string
}

func (r recordingCache) Connect() {
	r.Connect()
}

func (r recordingCache) Get(key string) (interface{}, bool) {
	r.addReq("cache get", key)
	data, ok := r.c.Get(key)
	r.addResp("cache resp",
		map[string]interface{}{"data": data, "ok": ok},
		nil)
	return data, ok
}

func (r recordingCache) Set(key string, value interface{}) {
	r.addReq("cache set", map[string]interface{}{key: value})
	r.c.Set(key, value)
}

func (r recordingCache) Clear(key string) {
	r.addReq("cache clear", key)
	r.c.Clear(key)

}

func (r recordingCache) HealthCheck() (bool, error) {

	r.addReq("cache HealthCheck")
	ok, err := r.c.HealthCheck()
	r.addResp("cache HealthCheck", ok, err)

	return ok, err
}

func (r recordingCache) addReq(header string, body ...interface{}) {

	var data string

	if len(body) > 0 {
		data = toJson(body[0])
	}

	r.recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    "service",
		Target:    r.sourceName,
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingCache) addResp(header string, body interface{}, err error) {

	var data string

	if err != nil {
		data = err.Error()
	} else {
		data = toJson(body)
	}

	r.recorder.AddMessageResponse(apitest.MessageResponse{
		Source:    r.sourceName,
		Target:    "service",
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

type recordingService struct {
	service    service.Service
	recorder   *apitest.Recorder
	sourceName string
}

func (r recordingService) GetCache() cache.Cache {
	panic("implement me")
}

func (r recordingService) addReq(header string, body ...interface{}) {

	var data string

	if len(body) > 0 {
		data = toJson(body[0])
	}

	r.recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    apitest.SystemUnderTestDefaultName,
		Target:    r.sourceName,
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingService) addResp(header string, body interface{}, err error) {

	var data string

	if err != nil {
		data = err.Error()
	} else {
		data = toJson(body)
	}

	r.recorder.AddMessageResponse(apitest.MessageResponse{
		Source:    r.sourceName,
		Target:    apitest.SystemUnderTestDefaultName,
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingService) HealthCheck() (bool, error) {

	r.addReq("HealthCheck", nil)
	ok, err := r.service.HealthCheck()
	r.addResp("HealthCheck", ok, err)

	return ok, err

}

func (r recordingService) Repository() db.Repository {
	return r.service.Repository()
}

func (r recordingService) GetInfobase(ctt service.ClientContext) (*serialize.InfobaseInfo, error) {
	panic("implement me")
}

func (r recordingService) CreateInfobase(ctt service.ClientContext, info *serialize.InfobaseInfo, createDB bool) (*serialize.InfobaseInfo, error) {
	panic("implement me")
}

func (r recordingService) UpdateInfobase(ctt service.ClientContext, info *serialize.InfobaseInfo) (*serialize.InfobaseInfo, error) {
	panic("implement me")
}

func (r recordingService) DropInfobase(ctt service.ClientContext, deleteDB bool) error {
	panic("implement me")
}

func (r recordingService) GetSessions(ctt service.ClientContext) (serialize.SessionInfoList, error) {
	panic("implement me")
}

func (r recordingService) TerminateSession(ctt service.ClientContext, sessionID uuid.UUID, msg string) error {
	panic("implement me")
}

func (r recordingService) GetClusters(ctt service.ClientContext) ([]*serialize.ClusterInfo, error) {
	panic("implement me")
}

func (r recordingService) GetClusterInfo(ctt service.ClientContext) (*serialize.ClusterInfo, error) {
	panic("implement me")
}

func (r recordingService) GetManagers(ctt service.ClientContext) ([]*serialize.ManagerInfo, error) {
	panic("implement me")
}

func (r recordingService) GetInfobases(ctt service.ClientContext) (serialize.InfobaseSummaryList, error) {
	panic("implement me")
}

func (r recordingService) GetServices(ctt service.ClientContext) ([]*serialize.ServiceInfo, error) {
	r.addReq("GetServices")
	ok, err := r.service.GetServices(ctt)
	r.addResp("GetServices", ok, err)
	return ok, err
}

func (r recordingService) GetLocks(ctt service.ClientContext) (serialize.LocksList, error) {
	panic("implement me")
}

func (r recordingService) GetConnections(ctt service.ClientContext) (serialize.ConnectionShortInfoList, error) {
	panic("implement me")
}

func (r recordingService) GetInfobaseConnections(client service.ClientContext, infobase string) (serialize.ConnectionShortInfoList, error) {
	panic("implement me")
}

func (r recordingService) Block(ctt service.ClientContext, blocker *models.InfobaseBlocker) (*models.InfobaseUnblocker, error) {
	panic("implement me")
}

func (r recordingService) Unblock(ctt service.ClientContext, unblocker *models.InfobaseUnblocker) (*serialize.InfobaseSummaryInfo, error) {
	panic("implement me")
}

func (r recordingService) GetProcesses(ctt service.ClientContext) (serialize.ProcessInfoList, error) {
	panic("implement me")
}

func (r recordingService) GetProcessInfo(ctt service.ClientContext) (*serialize.ProcessInfo, error) {
	panic("implement me")
}

func (r recordingService) TerminateConnection(ctt service.ClientContext, processID, connectionsID uuid.UUID) (models.ConnectionSig, error) {
	panic("implement me")
}

func (r recordingService) GetAppServers() (apps []*models.AppServer, err error) {

	r.addReq("GetAppServers")

	apps, err = r.service.GetAppServers()

	r.addResp("GetAppServers", apps, err)

	return
}

func (r recordingService) GetAppServer(name string) (*models.AppServer, error) {
	r.addReq("GetAppServer", name)

	app, err := r.service.GetAppServer(name)

	r.addResp("GetAppServer", app, err)

	return app, err
}

func (r recordingService) SetAppServer(app *models.AppServer) error {
	r.addReq("SetAppServer", app)

	err := r.service.SetAppServer(app)

	r.addResp("GetAppServer", nil, err)

	return err
}

func (r recordingService) DeleteAppServer(appName string) error {
	r.addReq("DeleteAppServer", appName)
	err := r.service.DeleteAppServer(appName)
	r.addResp("DeleteAppServer", nil, err)

	return err
}

func (r recordingService) GetAgentAdmins(client service.ClientContext) (serialize.UsersList, error) {
	panic("implement me")
}

func (r recordingService) GetAgentVersion(client service.ClientContext) (string, error) {
	panic("implement me")
}

func (r recordingService) RegAgentAdmin(client service.ClientContext, info serialize.UserInfo) error {
	panic("implement me")
}

func (r recordingService) UnregAgentAdmin(client service.ClientContext, user string) error {
	panic("implement me")
}

func (r recordingService) GetClusterAdmins(client service.ClientContext) (serialize.UsersList, error) {
	panic("implement me")
}

func (r recordingService) RegClusterAdmin(client service.ClientContext, info serialize.UserInfo) (*serialize.UserInfo, error) {
	panic("implement me")
}

func (r recordingService) UnregClusterAdmin(client service.ClientContext, user string) error {
	panic("implement me")
}

func (r recordingService) AddAppServer(app *models.AppServer) error {
	r.addReq("AddAppServer", app)

	err := r.service.AddAppServer(app)

	r.addResp("AddAppServer", nil, err)

	return err
}

type recordingRepository struct {
	repository db.Repository
	recorder   *apitest.Recorder
	sourceName string
}

func (r recordingRepository) addReq(header string, body interface{}) {

	var data string
	errBody, ok := body.(error)

	if ok {

		data = errBody.Error()
	} else {
		data = toJson(body)
	}

	r.recorder.AddMessageRequest(apitest.MessageRequest{
		Source:    "service",
		Target:    r.sourceName,
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingRepository) addResp(header string, body interface{}, err error) {

	var data string

	if err != nil {
		data = err.Error()
	} else {
		data = toJson(body)
	}

	r.recorder.AddMessageResponse(apitest.MessageResponse{
		Source:    r.sourceName,
		Target:    "service",
		Header:    header,
		Body:      data,
		Timestamp: time.Now().UTC(),
	})
}

func (r recordingRepository) AddAppServer(app models.AppServer) error {

	r.addReq("AddAppServer", app)

	err := r.repository.AddAppServer(app)

	r.addResp("AddAppServer", nil, err)

	return err
}

func (r recordingRepository) Clear() error {
	panic("implement me")
}

func (r recordingRepository) Db() string {
	return r.repository.Db()
}

func (r recordingRepository) GetAppServers() (apps []*models.AppServer, err error) {

	r.addReq("GetAppServers", nil)

	apps, err = r.repository.GetAppServers()

	r.addResp("GetAppServers", apps, err)

	return
}

func (r recordingRepository) GetAppServer(name string) (*models.AppServer, error) {

	r.addReq("GetAppServer", name)

	app, err := r.repository.GetAppServer(name)

	r.addResp("GetAppServer", app, err)

	return app, err
}

func (r recordingRepository) SetAppServer(app models.AppServer) error {
	r.addReq("SetAppServer", app)

	err := r.repository.SetAppServer(app)

	r.addResp("SetAppServer", nil, err)

	return err
}

func (r recordingRepository) DeleteAppServer(appName string) error {

	r.addReq("DeleteAppServer", appName)
	err := r.repository.DeleteAppServer(appName)
	r.addResp("DeleteAppServer", nil, err)

	return err
}

func toJson(data interface{}) string {
	raw, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(raw)
}
