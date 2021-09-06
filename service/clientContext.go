package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/models"
	rclient "github.com/khorevaa/ras-client"
	"github.com/khorevaa/ras-client/serialize"
	"github.com/recoilme/pudge"
	uuid "github.com/satori/go.uuid"
)

//var _ rclient.Api = (ClientContext)(nil)

type ClientContext struct {
	client     rclient.Api
	App        *models.AppServer
	ctx        context.Context
	requestCtx *fiber.Ctx

	force bool
}

func (c ClientContext) GetApiClient() rclient.Api {
	return c.client
}

func (c ClientContext) Version() string {
	return c.client.Version()
}

func (c ClientContext) Close() error {
	return c.client.Close()
}

func (c ClientContext) AuthenticateAgent(user, password string) {
	c.client.AuthenticateAgent(user, password)
}

func (c ClientContext) AuthenticateCluster(cluster uuid.UUID, user, password string) {
	c.client.AuthenticateCluster(cluster, user, password)
}

func (c ClientContext) AuthenticateInfobase(infobase uuid.UUID, user, password string) {
	c.client.AuthenticateInfobase(infobase, user, password)
}

func (c ClientContext) GetClusters() ([]*serialize.ClusterInfo, error) {
	return c.client.GetClusters(c.ctx)
}

func (c ClientContext) GetAgentAdmins() (serialize.UsersList, error) {
	return c.client.GetAgentAdmins(c.ctx)
}

func (c ClientContext) GetAgentVersion() (string, error) {
	return c.client.GetAgentVersion(c.ctx)
}

func (c ClientContext) RegAgentAdmin(user serialize.UserInfo) error {
	return c.client.RegAgentAdmin(c.ctx, user)
}

func (c ClientContext) UnregAgentAdmin(user string) error {
	return c.client.UnregAgentAdmin(c.ctx, user)
}

func (c ClientContext) GetClusterAdmins(cluster uuid.UUID) (serialize.UsersList, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterAdmins(c.ctx, cluster)
}

func (c ClientContext) RegClusterAdmin(cluster uuid.UUID, user serialize.UserInfo) error {
	c.AddAuth(cluster)
	return c.client.RegClusterAdmin(c.ctx, cluster, user)
}

func (c ClientContext) UnregClusterAdmin(cluster uuid.UUID, user string) error {
	c.AddAuth(cluster)
	return c.client.UnregClusterAdmin(c.ctx, cluster, user)
}

func (c ClientContext) GetClusterInfo(cluster uuid.UUID) (serialize.ClusterInfo, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterInfo(c.ctx, cluster)
}

func (c ClientContext) GetClusterInfobases(cluster uuid.UUID) (serialize.InfobaseSummaryList, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterInfobases(c.ctx, cluster)
}

func (c ClientContext) GetClusterServices(cluster uuid.UUID) ([]*serialize.ServiceInfo, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterServices(c.ctx, cluster)
}

func (c ClientContext) GetClusterManagers(cluster uuid.UUID) ([]*serialize.ManagerInfo, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterManagers(c.ctx, cluster)
}

func (c ClientContext) GetClusterSessions(cluster uuid.UUID) (serialize.SessionInfoList, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterSessions(c.ctx, cluster)
}

func (c ClientContext) GetInfobaseSessions(cluster uuid.UUID, infobase uuid.UUID) (serialize.SessionInfoList, error) {
	c.AddAuth(cluster, infobase)
	return c.client.GetInfobaseSessions(c.ctx, cluster, infobase)
}

func (c ClientContext) TerminateSession(cluster uuid.UUID, session uuid.UUID, msg string) error {
	c.AddAuth(cluster)
	return c.client.TerminateSession(c.ctx, cluster, session, msg)
}

func (c ClientContext) GetClusterLocks(cluster uuid.UUID) (serialize.LocksList, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterLocks(c.ctx, cluster)
}

func (c ClientContext) GetInfobaseLocks(cluster uuid.UUID, infobase uuid.UUID) (serialize.LocksList, error) {
	c.AddAuth(cluster, infobase)
	return c.client.GetInfobaseLocks(c.ctx, cluster, infobase)
}

func (c ClientContext) GetSessionLocks(cluster uuid.UUID, infobase uuid.UUID, session uuid.UUID) (serialize.LocksList, error) {
	c.AddAuth(cluster, infobase)
	return c.client.GetSessionLocks(c.ctx, cluster, infobase, session)
}

func (c ClientContext) GetConnectionLocks(cluster uuid.UUID, connection uuid.UUID) (serialize.LocksList, error) {
	c.AddAuth(cluster)
	return c.client.GetConnectionLocks(c.ctx, cluster, connection)
}

func (c ClientContext) GetClusterConnections(cluster uuid.UUID) (serialize.ConnectionShortInfoList, error) {
	c.AddAuth(cluster)
	return c.client.GetClusterConnections(c.ctx, cluster)
}

func (c ClientContext) GetInfobaseConnections(cluster uuid.UUID, infobase uuid.UUID) (serialize.ConnectionShortInfoList, error) {
	c.AddAuth(cluster, infobase)
	return c.client.GetInfobaseConnections(c.ctx, cluster, infobase)
}

func (c ClientContext) DisconnectConnection(cluster uuid.UUID, process uuid.UUID, connection uuid.UUID, infobase uuid.UUID) error {
	c.AddAuth(cluster, infobase)
	return c.client.DisconnectConnection(c.ctx, cluster, process, connection, infobase)
}

func (c ClientContext) CreateInfobase(cluster uuid.UUID, infobase serialize.InfobaseInfo, mode int) (serialize.InfobaseInfo, error) {

	c.AddAuth(cluster)
	return c.client.CreateInfobase(c.ctx, cluster, infobase, mode)
}

func (c ClientContext) UpdateSummaryInfobase(cluster uuid.UUID, infobase serialize.InfobaseSummaryInfo) error {
	c.AddAuth(cluster, infobase.UUID)
	return c.client.UpdateSummaryInfobase(c.ctx, cluster, infobase)
}

func (c ClientContext) UpdateInfobase(cluster uuid.UUID, infobase serialize.InfobaseInfo) error {
	c.AddAuth(cluster, infobase.UUID)
	return c.client.UpdateInfobase(c.ctx, cluster, infobase)
}

func (c ClientContext) DropInfobase(cluster uuid.UUID, infobase uuid.UUID, mode int) error {
	c.AddAuth(cluster, infobase)
	return c.client.DropInfobase(c.ctx, cluster, infobase, mode)
}

func (c ClientContext) GetInfobaseInfo(cluster uuid.UUID, infobase uuid.UUID) (serialize.InfobaseInfo, error) {
	c.AddAuth(cluster, infobase)
	return c.client.GetInfobaseInfo(c.ctx, cluster, infobase)
}

func (c ClientContext) GetWorkingProcesses(cluster uuid.UUID) (serialize.ProcessInfoList, error) {
	c.AddAuth(cluster)
	return c.client.GetWorkingProcesses(c.ctx, cluster)
}

func (c ClientContext) GetWorkingProcessInfo(cluster, process uuid.UUID) (*serialize.ProcessInfo, error) {
	c.AddAuth(cluster)
	return c.client.GetWorkingProcessInfo(c.ctx, cluster, process)
}

func (c ClientContext) GetWorkingServers(cluster uuid.UUID) ([]*serialize.ServerInfo, error) {
	c.AddAuth(cluster)
	return c.client.GetWorkingServers(c.ctx, cluster)
}

func (c ClientContext) GetWorkingServerInfo(cluster, serverID uuid.UUID) (*serialize.ServerInfo, error) {
	c.AddAuth(cluster)
	return c.client.GetWorkingServerInfo(c.ctx, cluster, serverID)
}

func (c ClientContext) RegWorkingServer(cluster uuid.UUID, info *serialize.ServerInfo) (*serialize.ServerInfo, error) {
	c.AddAuth(cluster)
	return c.client.RegWorkingServer(c.ctx, cluster, info)
}

func (c ClientContext) UnRegWorkingServer(cluster, serverID uuid.UUID) error {
	c.AddAuth(cluster)
	return c.client.UnRegWorkingServer(c.ctx, cluster, serverID)
}

func (c ClientContext) Force() bool {
	return c.force
}

func (c ClientContext) Context() context.Context {
	return c.ctx
}

func appServerFromContext(ctx *fiber.Ctx) (*models.AppServer, error) {

	name := ctx.Params("app")

	serviceInterface := ctx.Context().UserValue("service")

	s, ok := serviceInterface.(Service)

	if !ok {
		return nil, errors.Internal.New("cannot get service from context ")
	}

	app, err := s.GetAppServer(name)

	if err != nil {

		if err == pudge.ErrKeyNotFound {
			return nil, errors.BadRequest.Newf("app <%s> not registered", name)
		}

		return nil, errors.BadRequest.Wrapf(err, "cannot get app <%s>", name)
	}

	return app, err
}

func GetClientContext(ctx *fiber.Ctx) (ClientContext, error) {

	app, err := appServerFromContext(ctx)
	if err != nil {
		return ClientContext{}, err
	}

	apiClient, err := app.Client()
	if err != nil {
		return ClientContext{}, err
	}

	client := ClientContext{
		App:        app,
		client:     apiClient,
		requestCtx: ctx,
		ctx:        ctx.Context(),
		force:      GetContextValueOrNil(ctx, "force").Bool(false),
	}

	return client, nil
}

func NewClientContext(app *models.AppServer, ctx *fiber.Ctx) *ClientContext {

	apiClient, _ := app.Client()

	return &ClientContext{
		App:        app,
		client:     apiClient,
		requestCtx: ctx,
		ctx:        ctx.Context(),
		force:      true,
	}
}

func (c ClientContext) GetContextValue(name string, unescape ...bool) (ContextValue, bool) {
	return GetContextValue(c.requestCtx, name, unescape...)
}

var NeedClusterID = errors.BadRequest.New("need set cluster id")

func (c ClientContext) GetClusterID() (uuid.UUID, bool) {

	value := GetContextValueOrNil(c.requestCtx, "cluster cluster-id", true)

	id, err := value.UUID()

	if err != nil {
		return id, false
	}

	return id, true

}

func (c ClientContext) GetClusterIDOrNil() uuid.UUID {
	id, _ := c.GetClusterID()
	return id

}

func (c ClientContext) GetInfobaseID() (ContextValue, bool) {

	val := GetContextValueOrNil(c.requestCtx, "infobase infobase-id", true)

	return val, !val.Empty()

}

func (c ClientContext) AddAuth(cluster uuid.UUID, infobase ...uuid.UUID) {

	c.authCluster(cluster)

	if len(infobase) == 1 {
		c.authInfobase(infobase[0])
	}

}
func (c ClientContext) authCluster(cluster uuid.UUID) {

	if cluster == uuid.Nil {
		return
	}

	user, _ := c.GetContextValue("cluster-usr", true)

	if len(user) == 0 {
		return
	}

	pwd, _ := c.GetContextValue("cluster-pwd", true)

	c.client.AuthenticateCluster(cluster, user.String(), pwd.String())

}

func (c ClientContext) authInfobase(infobase uuid.UUID) {

	if infobase == uuid.Nil {
		return
	}
	user, _ := c.GetContextValue("infobase-usr", true)

	if len(user) == 0 {
		return
	}

	pwd, _ := c.GetContextValue("infobase-pwd", true)

	c.client.AuthenticateInfobase(infobase, user.String(), pwd.String())

}

func (c ClientContext) HealthCheck() (bool, error) {

	_, err := c.client.GetAgentVersion(c.ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}
