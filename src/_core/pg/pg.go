package pg

import (
	"context"
	myContext "exex-chart/src/_core/context"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DB struct {
	conn *pgx.Conn
}

var Ctx = context.Background()

func Connect() (*DB, error) {
	host := myContext.Config.Pg.Host
	conn, err := pgx.Connect(Ctx, host)

	if err != nil {
		log.Errorf("FAILED TO CONNECT TO DATABASE: %w", err)
		return nil, err
	}
	return &DB{conn: conn}, nil
}

func (pg *DB) Close() {
	pg.conn.Close(Ctx)
}

func (pg *DB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return pg.conn.QueryRow(ctx, sql, args...)
}

func (pg *DB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return pg.conn.Query(ctx, sql, args...)
}

func (pg *DB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return pg.conn.Exec(ctx, sql, args...)
}
