package couchdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/OscarLlamas6/hex-arch-go/settings"
	"sync"
	"time"

	_ "github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
	logger "github.com/sirupsen/logrus"
	slicesUtils "golang.org/x/exp/slices"
)

var (
	couchDBClient *kivik.Client
	couchDB       *kivik.DB
	onceCouchLoad sync.Once
)

func connectDB() error {

	var err error

	onceCouchLoad.Do(func() {
		var i int
		for {
			if i >= 30 {
				fmt.Println("Error connecting to CouchDB :(")
				err = errors.New("error connecting to CouchDB :(")
				break
			}
			time.Sleep(3 * time.Second)
			couchDBClient, err = kivik.New("couch", GetDSN())
			if err != nil {
				fmt.Println(err)
				logger.Info("Retrying in 3 seconds...")
				i++
				continue
			}

			var allDBs []string
			allDBs, err = couchDBClient.AllDBs(context.TODO())
			if err != nil {
				fmt.Println(err)
				break
			}

			if !slicesUtils.Contains(allDBs, settings.AppConfig.CouchDBName) {
				err = couchDBClient.CreateDB(context.TODO(), settings.AppConfig.CouchDBName)
				if err != nil {
					fmt.Println(err)
					break
				}
			}

			couchDB = couchDBClient.DB(context.TODO(), settings.AppConfig.CouchDBName)

			logger.Info("CouchDB connected successfully :D")
			break
		}
	})

	return err
}
