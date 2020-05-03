package models

import (
	"context"
	"fmt"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/utils"
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

	DB.AddQueryHook(dbLogger{})
}

// CheckConnection checks if the data base is available
func CheckConnection() {
	log.Info("Checking DB connection")

	_, err := DB.ExecOne("SELECT 1")
	utils.Check(err, true)
}

// TODO
type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}
