package database

import (
	"github.com/khorevaa/odin/models"
	"github.com/recoilme/pudge"
	"log"
	"path/filepath"
)

const pathPrefix = ".db"

var _ Repository = (*pudgeDb)(nil)

func NewRepository(path string) Repository {

	dbPath := path

	return &pudgeDb{
		db: filepath.Join(dbPath, pathPrefix, "app"),
	}
}

type pudgeDb struct {
	db string
}

func (r *pudgeDb) Clear() error {
	return pudge.DeleteFile(r.db)
}

func (r *pudgeDb) Db() string {
	return r.db
}

func (r *pudgeDb) GetAppServers() (apps []*models.AppServer, err error) {
	appdb, err := pudge.Open(r.db, pudge.DefaultConfig)

	if err != nil {
		return apps, err
	}

	defer appdb.Close()

	keys, err := appdb.Keys(nil, 0, 0, false)

	if err == nil {
		for _, k := range keys {
			val := &models.AppServer{}
			err := appdb.Get(k, val)

			if err != nil {
				log.Println("err", err)
				continue
			}
			apps = append(apps, val)

		}

	}
	return
}

func (r *pudgeDb) GetAppServer(name string) (*models.AppServer, error) {
	val := &models.AppServer{}
	err := pudge.Get(r.db, name, val)

	return val, err
}

func (r *pudgeDb) SetAppServer(app models.AppServer) error {

	err := pudge.Set(r.db, app.Name, app)

	app.Reload()

	return err
}

func (r *pudgeDb) AddAppServer(app models.AppServer) error {

	app = prepareAppServer(app)

	err := pudge.Set(r.db, app.Name, app)

	return err
}

func (r *pudgeDb) DeleteAppServer(appName string) error {
	err := pudge.Delete(r.db, appName)

	return err
}
