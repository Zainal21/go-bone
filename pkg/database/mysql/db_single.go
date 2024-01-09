package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	// check in runtime implement Databaser
	_ Adapter = (*DB)(nil)
)

type DB struct {
	db *sqlx.DB
	//instanceID string
	tx   *sqlx.Tx
	conn *sqlx.Conn // the Conn of the Tx, when tx != nil
	//opts       sql.TxOptions // valid when tx != nil
	reaMode bool
	dbName  string
}

func New(db *sqlx.DB, readMode bool, dbName string) *DB {
	return &DB{
		db:      db,
		reaMode: readMode,
		dbName:  dbName,
	}
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) InTransaction() bool {
	return db.tx != nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.db.Close()
}

// Exec executes a SQL statement and returns the number of rows it affected.
func (db *DB) Exec(ctx context.Context, query string, args ...any) (_ int64, err error) {
	if db.reaMode {
		return 0, fmt.Errorf("database mode read only")
	}

	res, err := db.execResult(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("rows affected: %v", err)
	}

	return n, nil
}

// execResult executes a SQL statement and returns a sql.Result.
func (db *DB) execResult(ctx context.Context, query string, args ...any) (res sql.Result, err error) {
	if db.tx != nil {
		return db.tx.ExecContext(ctx, query, args...)
	}

	return db.db.ExecContext(ctx, query, args...)
}

// Query runs the DB query.
func (db *DB) Query(ctx context.Context, dst any, query string, args ...any) error {
	if db.tx != nil {
		return db.tx.SelectContext(ctx, dst, query, args...)
	}

	return db.db.SelectContext(ctx, dst, query, args...)
}

// QueryRow runs the query and returns a single row.
func (db *DB) QueryRow(ctx context.Context, dst interface{}, query string, args ...any) error {

	if db.tx != nil {
		return db.tx.GetContext(ctx, dst, query, args...)
	}

	return db.db.GetContext(ctx, dst, query, args...)
}

// QueryX runs the DB query.
func (db *DB) QueryX(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if db.tx != nil {
		return db.tx.QueryContext(ctx, query, args...)
	}

	return db.db.QueryContext(ctx, query, args...)
}

// QueryRowX runs the query and returns a single row.
func (db *DB) QueryRowX(ctx context.Context, query string, args ...any) *sql.Row {
	if db.tx != nil {
		return db.tx.QueryRowContext(ctx, query, args...)
	}

	return db.db.QueryRowContext(ctx, query, args...)
}

// Transact executes the given function in the context of a SQL transaction at
// the given isolation level
func (db *DB) Transact(ctx context.Context, iso sql.IsolationLevel, txFunc func(*DB) error) (err error) {
	if db.reaMode {
		return fmt.Errorf("database mode read only")
	}

	// For the levels which require retry, see
	// https://www.postgresql.org/docs/11/transaction-iso.html.
	opts := &sql.TxOptions{Isolation: iso}

	return db.transact(ctx, opts, txFunc)
}

func (db *DB) transact(ctx context.Context, opts *sql.TxOptions, txFunc func(*DB) error) (err error) {
	if db.InTransaction() {
		return errors.New("db transact function was called on a DB already in a transaction")
	}

	conn, err := db.db.Connx(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	tx, err := conn.BeginTxx(ctx, opts)
	if err != nil {
		return fmt.Errorf("tx begin: %w", err)
	}

	//defer func() {
	//	if p := recover(); p != nil {
	//		tx.Rollback()
	//	} else if err != nil {
	//		tx.Rollback()
	//	} else {
	//		if txErr := tx.Commit(); txErr != nil {
	//			err = fmt.Errorf("tx commit: %w", txErr)
	//		}
	//	}
	//}()

	dbtx := New(db.db, false, db.dbName)
	dbtx.tx = tx
	dbtx.conn = conn
	//dbtx.opts = *opts

	if err := txFunc(dbtx); err != nil {
		tx.Rollback()
		return fmt.Errorf("fn(tx): %w", err)
	}

	return tx.Commit()
}

// BeginTx start new transaction session
func (d *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.db.BeginTx(ctx, opts)
}
