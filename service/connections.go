package service

import (
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
)

func (s *service) GetConnections(client ClientContext) (serialize.ConnectionShortInfoList, error) {

	infobase, ok := client.GetInfobaseID()

	if ok {

		infobaseSummaryInfo, err := s.findInfobase(client, infobase.String(), client.GetClusterIDOrNil())

		if err != nil {
			return nil, err
		}

		return client.GetInfobaseConnections(infobaseSummaryInfo.ClusterID, infobaseSummaryInfo.UUID)
	}

	clusterID, err := s.getClusterID(client)

	if err != nil {
		return nil, err
	}

	return client.GetClusterConnections(clusterID)

}

func (s *service) GetInfobaseConnections(client ClientContext, infobase string) (serialize.ConnectionShortInfoList, error) {

	infobaseSummaryInfo, err := s.findInfobase(client, infobase)

	if err != nil {
		return nil, err
	}

	return client.GetInfobaseConnections(infobaseSummaryInfo.ClusterID, infobaseSummaryInfo.UUID)

}

func (s *service) TerminateConnection(ctt ClientContext, processID, connectionsID uuid.UUID) (models.ConnectionSig, error) {

	clusterID, err := s.getClusterID(ctt)

	if err != nil {
		return models.ConnectionSig{}, err
	}

	return models.ConnectionSig{
		ClusterID: clusterID,
		Process:   processID,
		UUID:      connectionsID,
	}, ctt.DisconnectConnection(clusterID, processID, connectionsID, uuid.Nil)

}
