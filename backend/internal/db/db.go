package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPGConn(dsn string) (*pgx.Conn, error) {
	if conn, err := pgx.Connect(context.Background(), dsn); err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}
