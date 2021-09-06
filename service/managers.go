package service

import (
	"github.com/khorevaa/ras-client/serialize"
)

func (s *service) GetManagers(ctt ClientContext) ([]*serialize.ManagerInfo, error) {

	clusterID, err := s.getClusterID(ctt)

	if err != nil {
		return nil, err
	}

	info, err := ctt.GetClusterManagers(clusterID)

	if err != nil {
		return nil, err
	}
	return info, nil
}
