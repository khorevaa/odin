package service

import (
	errors2 "github.com/khorevaa/odin/errors"
	"github.com/khorevaa/ras-client/serialize"
)

func (s *service) GetProcesses(client ClientContext) (serialize.ProcessInfoList, error) {

	clusterID, err := s.getClusterID(client)

	if err != nil {
		return nil, err
	}

	info, err := client.GetWorkingProcesses(clusterID)

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *service) GetProcessInfo(client ClientContext) (*serialize.ProcessInfo, error) {

	clusterID, err := s.getClusterID(client)

	if err != nil {
		return nil, err
	}

	process, ok := client.GetContextValue("process process-id")

	if !ok {
		return nil, errors2.BadRequest.New("process id must be set")
	}

	processID, err := process.UUID()
	if err != nil {
		return nil, errors2.BadCommand.Wrapf(err, "process-id <%s> is incorrect", process.String())
	}

	info, err := client.GetWorkingProcessInfo(clusterID, processID)

	if err != nil {
		return nil, err
	}
	return info, nil
}
