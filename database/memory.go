package database

import (
	"github.com/khorevaa/odin/models"
	"sync"
)

var _ Repository = (*InMemory)(nil)

type InMemory struct {
	m *sync.Map
}

func (i *InMemory) GetAppServers() (apps []*models.AppServer, err error) {

	i.m.Range(func(_, val interface{}) bool {
		app := val.(models.AppServer)
		apps = append(apps, &app)
		return true
	})

	return
}

func (i *InMemory) GetAppServer(name string) (*models.AppServer, error) {
	app, ok := i.m.Load(name)
	if !ok {
		return nil, ErrorNotFound
	}

	appServer := app.(models.AppServer)

	return &appServer, nil
}

func (i InMemory) SetAppServer(app models.AppServer) error {

	i.m.Store(app.Name, app)
	return nil
}

func (i *InMemory) AddAppServer(app models.AppServer) error {

	app = prepareAppServer(app)

	i.m.Store(app.Name, app)

	return nil

}

func (i *InMemory) DeleteAppServer(appName string) error {
	_, ok := i.m.LoadAndDelete(appName)
	if !ok {
		return ErrorNotFound
	}

	return nil
}

func (i *InMemory) Db() string {
	return ""
}

func (i *InMemory) Clear() error {
	return nil
}

func NewMemoryRepository() Repository {

	return &InMemory{
		m: &sync.Map{},
	}

}
