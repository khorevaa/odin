package service

import (
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
)

func (s service) findInfobaseInList(list serialize.InfobaseSummaryList, infobaseID string) (*serialize.InfobaseSummaryInfo, bool) {

	if id := uuid.FromStringOrNil(infobaseID); id != uuid.Nil {
		return list.ByID(id)
	}

	return list.ByName(infobaseID)

}

func (s *service) getClusterID(client ClientContext, clusterID ...uuid.UUID) (uuid.UUID, error) {

	if len(clusterID) > 0 &&
		clusterID[0] != uuid.Nil {
		return clusterID[0], nil
	}

	cluster, ok := client.GetClusterID()
	if ok {
		return cluster, nil
	}

	clusters, err := s.getClusters(client)

	if err != nil {
		return uuid.Nil, err
	}

	if len(clusters) == 1 {
		return clusters[0].UUID, nil
	}

	return uuid.Nil, errors.BadRequest.New("to many clusters. Set <cluster id> value manually")
}

func (s *service) getAnyClusterID(client ClientContext, clusterID ...uuid.UUID) (uuid.UUID, error) {

	if len(clusterID) > 0 &&
		clusterID[0] != uuid.Nil {
		return clusterID[0], nil
	}

	clusters, err := s.getClusters(client)

	if err != nil {
		return uuid.Nil, err
	}

	if len(clusters) > 0 {
		return clusters[0].UUID, nil
	}

	return uuid.Nil, errors.BadRequest.New("no registered clusters. Set <cluster id> value manually")
}

func (s service) findInfobase(client ClientContext, infobaseID string, clusterID ...uuid.UUID) (*serialize.InfobaseSummaryInfo, error) {

	clusters, err := s.getClusters(client)

	if err != nil {
		return nil, err
	}

	if len(clusterID) == 1 && clusterID[0] != uuid.Nil {
		list, err := s.getClusterInfobases(client, clusterID[0])
		if err != nil {
			return nil, err
		}

		summaryInfo, ok := s.findInfobaseInList(list, infobaseID)

		if !ok {
			return nil, errors.BadRequest.Newf("infobase not found by name or uuid <%s> on cluster <%s>",
				infobaseID, clusterID[0].String())
		}

		return summaryInfo, nil
	}

	var summaryInfoList []*serialize.InfobaseSummaryInfo

	for _, cluster := range clusters {
		list, err := s.getClusterInfobases(client, cluster.UUID)
		if err != nil {
			continue
		}

		summaryInfo, ok := s.findInfobaseInList(list, infobaseID)
		if ok {
			summaryInfoList = append(summaryInfoList, summaryInfo)
		}

	}

	switch len(summaryInfoList) {

	case 1:
		return summaryInfoList[0], nil
	case 0:
		return nil, errors.BadRequest.Newf("infobase not found by name or uuid <%s>", infobaseID)
	default:
		return nil, errors.BadRequest.Newf("find to many infobases with <%s>."+
			" Set <cluster-id> value manually", infobaseID)
	}
}

func (s *service) getClusterInfobases(client ClientContext, clusterID uuid.UUID) (serialize.InfobaseSummaryList, error) {

	cacheKey := clusterID.String()

	if list, ok := s.getCacheInfobases(cacheKey); ok && !client.Force() {
		return list, nil
	}

	list, err := client.GetClusterInfobases(clusterID)

	if err != nil {
		return nil, err
	}

	s.setCacheInfobases(cacheKey, list)

	return list, nil

}

func (s *service) getClusters(client ClientContext) ([]*serialize.ClusterInfo, error) {

	if err := s.validLicense(); err != nil {
		return nil, err
	}

	cacheKey := client.App.Name

	if list, ok := s.getCacheClusters(cacheKey); ok && !client.Force() {
		return list, nil
	}

	list, err := client.GetClusters()

	if err != nil {
		return nil, err
	}

	s.setCacheClusters(cacheKey, list)

	return list, nil

}
