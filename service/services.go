package service

import "github.com/khorevaa/ras-client/serialize"

func (s *service) GetServices(ctt ClientContext) ([]*serialize.ServiceInfo, error) {

	clusterID, err := s.getClusterID(ctt)

	if err != nil {
		return nil, err
	}

	info, err := ctt.GetClusterServices(clusterID)

	if err != nil {
		return nil, err
	}
	return info, nil
}
