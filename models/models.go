package models

import (
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"

	"github.com/absinsekt/pnk/lib/configuration"
	"github.com/absinsekt/pnk/lib/core"
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
	core.Check(err, true)
}
