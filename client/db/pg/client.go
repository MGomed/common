package pg

import (
	"context"
	"fmt"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"

	db "github.com/MGomed/common/client/db"
	logger "github.com/MGomed/common/logger"
)

type pgClient struct {
	dbc db.DB
}

// New is a pgClient constructor
func New(ctx context.Context, log logger.Interface, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		dbc: &pg{
			log: log,
			dbc: dbc,
		},
	}, nil
}

// DB is dbc getter
func (c *pgClient) DB() db.DB {
	return c.dbc
}

// Close closes dbc
func (c *pgClient) Close() error {
	if c.dbc != nil {
		c.dbc.Close()
	}

	return nil
}
