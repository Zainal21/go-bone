package mysql

import (
	"context"
	"database/sql"
)

type Adapter interface {
	Ping() error
	InTransaction() bool
	Close() error
	Query(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	QueryRow(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	QueryX(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowX(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(ctx context.Context, query string, args ...interface{}) (_ int64, err error)
	Transact(ctx context.Context, iso sql.IsolationLevel, txFunc func(*DB) error) (err error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
