package service

import (
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
)

func (s *service) GetLocks(client ClientContext) (serialize.LocksList, error) {

	if session, ok := client.GetContextValue("session session-id"); ok {
		return s.getSessionLocks(client, session)
	}

	if infobase, ok := client.GetInfobaseID(); ok {
		return s.getInfobaseLocks(client, infobase)
	}

	if connection, ok := client.GetContextValue("connection connection-id"); ok {
		return s.getSessionLocks(client, connection)
	}

	clusterID, err := s.getClusterID(client)
	if err != nil {
		return nil, err
	}
	info, err := client.GetClusterLocks(clusterID)

	if err != nil {
		return nil, err
	}
	return info, nil

}

func (s *service) getSessionLocks(client ClientContext, session ContextValue) (serialize.LocksList, error) {

	sessionID, err := session.UUID()

	if err != nil {
		return nil, errors.BadRequest.Newf("session id <%s> is incorrect", session.String())
	}

	infobase, ok := client.GetInfobaseID()

	if !ok {
		return nil, errors.BadRequest.New("infobase id or name must be set")
	}

	summaryInfo, err := s.findInfobase(client, infobase.String())
	if err != nil {
		return nil, err
	}

	info, err := client.GetSessionLocks(summaryInfo.ClusterID, summaryInfo.UUID, sessionID)

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *service) getInfobaseLocks(client ClientContext, infobase ContextValue) (serialize.LocksList, error) {

	summaryInfo, err := s.findInfobase(client, infobase.String())
	if err != nil {
		return nil, err
	}

	info, err := client.GetInfobaseLocks(summaryInfo.ClusterID, summaryInfo.UUID)

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *service) getConnectionsLocks(client ClientContext, connection ContextValue) (serialize.LocksList, error) {

	connectionID, err := connection.UUID()

	if err != nil {
		return nil, errors.BadRequest.Newf("connection id <%s> is incorrect", connection.String())
	}

	clusterID, err := s.getClusterID(client)
	if err != nil {
		return nil, err
	}

	info, err := client.GetConnectionLocks(clusterID, connectionID)

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *service) getClusterLocks(client ClientContext, clusterID uuid.UUID) (serialize.LocksList, error) {

	info, err := client.GetClusterLocks(clusterID)

	if err != nil {
		return nil, err
	}
	return info, nil
}
