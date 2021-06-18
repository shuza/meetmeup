package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctc context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.Query)
	return nil
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
