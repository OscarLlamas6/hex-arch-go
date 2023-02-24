package couchdb

import (
	"fmt"

	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func GetDSN() string {

	var DSN = fmt.Sprintf("http://%s:%s@%s:%s/",
		settings.AppConfig.CouchDBUser,
		settings.AppConfig.CouchDBPass,
		settings.AppConfig.CouchDBHost,
		settings.AppConfig.CouchDBPort)

	return DSN
}
