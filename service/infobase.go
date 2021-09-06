package service

import (
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
)

func (s *service) GetInfobases(client ClientContext) (serialize.InfobaseSummaryList, error) {

	cluster, _ := client.GetClusterID()
	clusterID, _ := s.getClusterID(client, cluster)

	if clusterID != uuid.Nil {
		return s.getClusterInfobases(client, clusterID)
	}

	clusters, err := s.getClusters(client)

	if err != nil {
		return nil, err
	}

	var listAll serialize.InfobaseSummaryList

	for _, cluster := range clusters {
		list, err := s.getClusterInfobases(client, cluster.UUID)
		if err != nil {
			continue
		}
		listAll = append(listAll, list...)
	}

	return listAll, nil

}

func (s *service) CreateInfobase(client ClientContext, info *serialize.InfobaseInfo, createDB bool) (*serialize.InfobaseInfo, error) {

	cluster, _ := client.GetClusterID()
	clusterID, err := s.getAnyClusterID(client, cluster)

	if err != nil {
		return nil, err
	}

	mode := 0

	if createDB {
		mode = 1
	}

	infobaseInfo, err := client.CreateInfobase(clusterID, *info, mode)

	if err != nil {
		return nil, err
	}

	s.clearCacheInfobases(clusterID.String())

	return &infobaseInfo, nil
}

func (s *service) UpdateInfobase(client ClientContext, info *serialize.InfobaseInfo) (*serialize.InfobaseInfo, error) {

	err := client.UpdateInfobase(info.ClusterID, *info)

	if err != nil {
		return nil, err
	}

	newInfo, err := client.GetInfobaseInfo(info.ClusterID, info.UUID)

	return &newInfo, err

}

func (s *service) DropInfobase(client ClientContext, dropDB bool) error {

	infobaseID, _ := client.GetInfobaseID()

	if infobaseID.Empty() {
		return errors.BadRequest.New("infobase id or name must be set")
	}

	summaryInfo, err := s.findInfobase(client, infobaseID.String())
	if err != nil {
		return err
	}

	mode := 0

	if dropDB {
		mode = 1
	}

	err = client.DropInfobase(summaryInfo.ClusterID, summaryInfo.UUID, mode)

	return err

}

func (s *service) GetInfobase(client ClientContext) (*serialize.InfobaseInfo, error) {

	infobaseID, _ := client.GetInfobaseID()

	if infobaseID.Empty() {
		return nil, errors.BadRequest.New("infobase id or name must be set")
	}

	summaryInfo, err := s.findInfobase(client, infobaseID.String())
	if err != nil {
		return nil, err
	}

	info, err := client.GetInfobaseInfo(summaryInfo.ClusterID, summaryInfo.UUID)

	return &info, err
}
