package pg

import (
	"fmt"

	"github.com/OscarLlamas6/hex-arch-go/settings"
)

func GetDSN() string {

	var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		settings.AppConfig.DBHost,
		settings.AppConfig.DBUser,
		settings.AppConfig.DBPass,
		settings.AppConfig.DBName,
		settings.AppConfig.DBPort)

	return DSN
}
