package service

import (
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/ras-client/serialize"
)

func (s *service) GetClusters(client ClientContext) ([]*serialize.ClusterInfo, error) {

	return s.getClusters(client)
}

func (s *service) GetClusterInfo(client ClientContext) (*serialize.ClusterInfo, error) {

	clusterID, ok := client.GetClusterID()

	if !ok {
		return nil, errors.BadRequest.New("incorrect or not set <cluster-id>")
	}

	info, err := client.GetClusterInfo(clusterID)

	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (s *service) RegCluster(client ClientContext, info *serialize.ClusterInfo) (*serialize.ClusterInfo, error) {

	return nil, nil
}

func (s *service) UnRegCluster(client ClientContext) error {

	return nil
}

func (s *service) GetClusterAdmins(client ClientContext) (serialize.UsersList, error) {

	clusterID, ok := client.GetClusterID()

	if !ok {
		return nil, errors.BadRequest.New("incorrect or not set <cluster-id>")
	}

	return client.GetClusterAdmins(clusterID)
}

func (s *service) RegClusterAdmin(client ClientContext, info serialize.UserInfo) (*serialize.UserInfo, error) {

	clusterID, ok := client.GetClusterID()

	if !ok {
		return nil, errors.BadRequest.New("incorrect or not set <cluster-id>")
	}

	return &info, client.RegClusterAdmin(clusterID, info)
}

func (s *service) UnregClusterAdmin(client ClientContext, user string) error {

	clusterID, ok := client.GetClusterID()

	if !ok {
		return errors.BadRequest.New("incorrect or not set <cluster-id>")
	}

	return client.UnregClusterAdmin(clusterID, user)
}
