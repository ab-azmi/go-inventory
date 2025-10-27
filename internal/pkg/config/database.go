package config

import (
	xtremedb "github.com/globalxtreme/go-core/v2/database"
	"gorm.io/gorm"
	"os"
)

var (
	PgSQL *gorm.DB
)

func InitDB() func() {
	var DBClose func()

	PgSQL, DBClose = xtremedb.Connect(xtremedb.DBConf{
		Driver:    xtremedb.POSTGRESQL_DRIVER,
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		ParseTime: true,
	})

	return DBClose
}
