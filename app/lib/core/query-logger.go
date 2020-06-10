package core

import (
	"context"

	"github.com/go-pg/pg/v9"
	log "github.com/sirupsen/logrus"
)

// QueryLogger logs formatted queries
type QueryLogger struct{}

// BeforeQuery dummy
func (ql QueryLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// AfterQuery logs formatted query after it finishes
func (ql QueryLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	log.Info(q.FormattedQuery())
	return nil
}
