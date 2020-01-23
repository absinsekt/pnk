package models

import (
	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/utils"
	"github.com/go-pg/pg"
	//  "github.com/go-pg/pg/orm"
)

// DB main database connection
var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		Database: configuration.DbName,
		User:     configuration.DbUser,
		Password: configuration.DbPassword,
	})

	_, err := DB.ExecOne("SELECT 1")
	utils.Check(err, true)
}
