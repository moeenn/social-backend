package db

import (
	"context"
	"sandbox/config"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context, dbConfig *config.DatabaseConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dbConfig.ConnectionURI)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
