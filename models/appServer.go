package models

import (
	"github.com/khorevaa/odin/ras"
	rclient "github.com/khorevaa/ras-client"
	"net"
)

type AppServer struct {
	Name         string            `json:"name"`
	Addr         string            `json:"addr"`
	Port         string            `json:"port"`
	Version      string            `json:"version"`
	Description  string            `json:"descr,omitempty" yaml:"descr"`
	AgentUsr     string            `json:"agent_usr,omitempty"`
	AgentPwd     string            `json:"agent_pwd,omitempty"`
	AgentVersion string            `json:"agent_version,omitempty"`
	Properties   map[string]string `json:"properties,omitempty"`
} // @Name AppServer

func (a *AppServer) init(api rclient.Api, inited bool) {
	if inited {
		return
	}

	if len(a.AgentUsr) > 0 {
		api.AuthenticateAgent(a.AgentUsr, a.AgentPwd)
	}
}

func (a *AppServer) Client() (rclient.Api, error) {

	api, inited := ras.LoadOrInit(a.Name, net.JoinHostPort(a.Addr, a.Port), a.Version)
	a.init(api, inited)

	return api, nil
}

func (a *AppServer) Reload() {

	c, err := a.Client()
	if err == nil {
		_ = c.Close()
	}

	api, inited := ras.Store(a.Name, net.JoinHostPort(a.Addr, a.Port), a.Version)
	a.init(api, inited)

}
