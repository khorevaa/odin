package service

import (
	"github.com/khorevaa/ras-client/serialize"
)

func (s *service) GetAgentAdmins(client ClientContext) (serialize.UsersList, error) {

	list, err := client.GetAgentAdmins()

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *service) GetAgentVersion(client ClientContext) (string, error) {

	version, err := client.GetAgentVersion()

	if err != nil {
		return "", err
	}
	return version, nil
}

func (s *service) RegAgentAdmin(client ClientContext, info serialize.UserInfo) error {

	err := client.RegAgentAdmin(info)

	return err
}

func (s *service) UnregAgentAdmin(client ClientContext, user string) error {

	return client.UnregAgentAdmin(user)
}
