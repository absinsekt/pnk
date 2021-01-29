package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	log "github.com/sirupsen/logrus"

	"github.com/absinsekt/pnk/lib/core"
)

// Pool main database connection pool
var Pool *pg.DB

// InitConnection initializes db connection
func InitConnection() {
	Pool = pg.Connect(&pg.Options{
		Database: core.Config.DbName,
		User:     core.Config.DbUser,
		Password: core.Config.DbPassword,
	})

	if core.Config.Debug {
		Pool.AddQueryHook(core.QueryLogger{})
	}

	CheckConnection()
}

// CheckConnection checks if the data base is available
func CheckConnection() {
	log.Info("Checking DB connection")

	_, err := Pool.ExecOne("SELECT 1")
	core.Check(err, true)
}

// RegisterTable todo
func RegisterTable(strct interface{}) {
	orm.RegisterTable(strct)
}
