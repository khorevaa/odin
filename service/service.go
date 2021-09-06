package service

import (
	"context"
	db "github.com/khorevaa/odin/database"
	_ "github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/odin/service/cache"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
)

const (
	clustersTpl  = "%s.clusters"
	infobasesTpl = "%s.infobases"
)

var _ Service = (*service)(nil)

//Service interface allows us to access the CRUD Operations
type Service interface {
	HealthCheck() (bool, error)
	Repository() db.Repository

	GetInfobase(ctt ClientContext) (*serialize.InfobaseInfo, error)
	CreateInfobase(ctt ClientContext, info *serialize.InfobaseInfo, createDB bool) (*serialize.InfobaseInfo, error)
	UpdateInfobase(ctt ClientContext, info *serialize.InfobaseInfo) (*serialize.InfobaseInfo, error)
	DropInfobase(ctt ClientContext, deleteDB bool) error

	GetSessions(ctt ClientContext) (serialize.SessionInfoList, error)
	TerminateSession(ctt ClientContext, sessionID uuid.UUID, msg string) error

	GetClusters(ctt ClientContext) ([]*serialize.ClusterInfo, error)
	GetClusterInfo(ctt ClientContext) (*serialize.ClusterInfo, error)
	GetManagers(ctt ClientContext) ([]*serialize.ManagerInfo, error)
	GetInfobases(ctt ClientContext) (serialize.InfobaseSummaryList, error)
	GetServices(ctt ClientContext) ([]*serialize.ServiceInfo, error)
	GetLocks(ctt ClientContext) (serialize.LocksList, error)
	GetConnections(ctt ClientContext) (serialize.ConnectionShortInfoList, error)
	GetInfobaseConnections(client ClientContext, infobase string) (serialize.ConnectionShortInfoList, error)
	Block(ctt ClientContext, blocker *models.InfobaseBlocker) (*models.InfobaseUnblocker, error)
	Unblock(ctt ClientContext, unblocker *models.InfobaseUnblocker) (*serialize.InfobaseSummaryInfo, error)
	GetProcesses(ctt ClientContext) (serialize.ProcessInfoList, error)
	GetProcessInfo(ctt ClientContext) (*serialize.ProcessInfo, error)
	TerminateConnection(ctt ClientContext, processID, connectionsID uuid.UUID) (models.ConnectionSig, error)

	GetAppServers() (apps []*models.AppServer, err error)
	GetAppServer(name string) (*models.AppServer, error)
	SetAppServer(app *models.AppServer) error
	DeleteAppServer(appName string) error

	GetAgentAdmins(client ClientContext) (serialize.UsersList, error)
	GetAgentVersion(client ClientContext) (string, error)
	RegAgentAdmin(client ClientContext, info serialize.UserInfo) error
	UnregAgentAdmin(client ClientContext, user string) error
	GetClusterAdmins(client ClientContext) (serialize.UsersList, error)
	RegClusterAdmin(client ClientContext, info serialize.UserInfo) (*serialize.UserInfo, error)
	UnregClusterAdmin(client ClientContext, user string) error
	AddAppServer(app *models.AppServer) error
	GetCache() cache.Cache
}

func NewService(cache cache.Cache, repository db.Repository) (Service, error) {

	return &service{
		cache:      cache,
		repository: repository,
	}, nil
}

type service struct {
	repository db.Repository
	cache      cache.Cache
}

func (s service) GetCache() cache.Cache {
	return s.cache
}

func (s service) validLicense() error {
	return nil
}

func (s service) Licensed() bool {

	return true
}

func (s service) AddAppServer(app *models.AppServer) error {

	return s.repository.AddAppServer(*app)

}

func (s *service) Repository() db.Repository {
	return s.repository
}

func (s *service) GetAppServers() (apps []*models.AppServer, err error) {
	return s.repository.GetAppServers()
}

func (s *service) GetAppServer(name string) (*models.AppServer, error) {
	return s.repository.GetAppServer(name)
}

func (s *service) SetAppServer(app *models.AppServer) error {

	api, err := app.Client()

	if err != nil {
		return err
	}

	app.AgentVersion, _ = api.GetAgentVersion(context.Background())
	app.Version = api.Version()

	return s.repository.SetAppServer(*app)
}

func (s *service) DeleteAppServer(appName string) error {
	return s.repository.DeleteAppServer(appName)
}

func (s *service) HealthCheck() (bool, error) {

	ok, err := s.cache.HealthCheck()
	return ok, err
}
