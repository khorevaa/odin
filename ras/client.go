package ras

import (
	rclient "github.com/khorevaa/ras-client"
	"sync"
)

type Storage interface {
	Store(id string, addr string, version string) (rclient.Api, bool)
	LoadOrInit(id string, addr string, version string) (rclient.Api, bool)
	Load(id string) (interface{}, bool)
}

var localStorage = newStorage()

func newStorage() Storage {
	return &storage{
		sMap: &sync.Map{},
	}
}

type storage struct {
	sMap *sync.Map
}

func (s *storage) Store(id string, addr string, version string) (rclient.Api, bool) {
	client := rclient.NewClient(addr, rclient.WithVersion(version))

	s.sMap.Store(id, client)

	return client, false
}

func (s *storage) LoadOrInit(id string, addr string, version string) (rclient.Api, bool) {
	client, ok := s.sMap.Load(id)

	if !ok {
		return Store(id, addr, version)
	}

	c := client.(rclient.Api)

	return c, true
}

func (s *storage) Load(id string) (interface{}, bool) {
	return s.sMap.Load(id)
}

func SetLocalStorage(s Storage) {
	localStorage = s
}

func Store(id string, addr string, version string) (rclient.Api, bool) {

	client, _ := localStorage.Store(id, addr, version)

	return client, false
}

func LoadOrInit(id string, addr string, version string) (rclient.Api, bool) {

	client, ok := localStorage.Load(id)

	if !ok {
		return localStorage.Store(id, addr, version)
	}

	c := client.(rclient.Api)

	return c, true
}
