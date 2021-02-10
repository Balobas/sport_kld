package database

/*
Обертка пакета "github.com/leesper/couchdb-golang"
*/

import (
	"github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"
	"sport_kld/app/settings"
	"strings"
)

type CouchDB struct {
	Server couchdb.Server
	DB     couchdb.Database
}

var DATABASE CouchDB

func init() {
	DATABASE = CouchDB{}
	err := DATABASE.Init(settings.STORAGE_USERNAME, settings.STORAGE_PASSWORD, settings.STORAGE_HOST)
	if err != nil {
		panic("db init error " + err.Error())
	}
}

func (c *CouchDB) Init(login, password, url string) error {
	server, err := couchdb.NewServer(url)
	if err != nil {
		panic(err)
	}
	c.Server = *server
	token, err := server.Login(login, password)
	if err != nil {
		return errors.New("login db error " + err.Error())
	}
	err = server.VerifyToken(token)
	if err != nil {
		return errors.New("verify token error " + err.Error())
	}

	DB, err := server.Get(settings.STORAGE_NAME)

	if err != nil {
		return errors.WithStack(err)
	}
	if DB == nil {
		return errors.New("nil pointer database ")
	}
	c.DB = *DB
	return nil
}

func (c *CouchDB) Set(key string, obj interface{}) error {
	objMap, err := couchdb.ToJSONCompatibleMap(obj)
	if err != nil {
		return errors.Wrap(err, "marshalling error ")
	}
	savedObj, err := c.DB.Get(key, nil)
	if err != nil {
		if ! strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if savedObj != nil {
		objMap["_rev"] = savedObj["_rev"]
	}
	return c.DB.Set(key, objMap)
}

func (c *CouchDB) Get(key string) (map[string]interface{}, error) {
	return c.DB.Get(key, nil)
}

func (c *CouchDB) Unmarshal(obj interface{}, objMap map[string]interface{}) (interface{}, error) {
	err := couchdb.FromJSONCompatibleMap(obj, objMap)
	return obj, err
}

func (c *CouchDB) Delete(key string) error {
	return c.DB.Delete(key)
}

func (c *CouchDB) QueryAllFieldsWithSelector(selector string) ([]map[string] interface{}, error) {
	return c.DB.Query(nil, selector, nil, nil, nil, nil)

}
