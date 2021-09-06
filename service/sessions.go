package service

import (
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
)

func (s *service) GetSessions(client ClientContext) (serialize.SessionInfoList, error) {

	infobase, ok := client.GetInfobaseID()

	if ok {

		summaryInfo, err := s.findInfobase(client, infobase.String())

		if err != nil {
			return nil, err
		}

		return client.GetInfobaseSessions(summaryInfo.ClusterID, summaryInfo.UUID)

	}

	clusterID, err := s.getClusterID(client)

	if err != nil {
		return nil, err
	}

	list, err := client.GetClusterSessions(clusterID)
	if err != nil {
		return nil, err
	}

	return list, nil

}

func (s *service) TerminateSession(client ClientContext, sessionID uuid.UUID, msg string) error {

	clusterID, err := s.getClusterID(client)

	if err != nil {
		return err
	}

	err = client.TerminateSession(clusterID, sessionID, msg)

	return err
}
