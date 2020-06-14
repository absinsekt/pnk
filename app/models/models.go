package models

import (
	"github.com/absinsekt/pnk/lib"
	"github.com/absinsekt/pnk/lib/core"

	"github.com/absinsekt/pnk/configuration"
	"github.com/go-pg/pg/v9"
	log "github.com/sirupsen/logrus"
)

// DB main database connection
var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		Database: configuration.DbName,
		User:     configuration.DbUser,
		Password: configuration.DbPassword,
	})

	if configuration.Debug {
		DB.AddQueryHook(core.QueryLogger{})
	}
}

// CheckConnection checks if the data base is available
func CheckConnection() {
	log.Info("Checking DB connection")

	_, err := DB.ExecOne("SELECT 1")
	lib.Check(err, true)
}
