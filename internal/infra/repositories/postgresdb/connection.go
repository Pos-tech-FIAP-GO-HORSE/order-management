package postgresdb

import (
	"context"
	"database/sql"
)

func Connect(ctx context.Context, driver, uri string) (*sql.DB, error) {
	conn, err := sql.Open(driver, uri)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
