package repositories

import (
	"database/sql"
)

type option struct {
	tx *sql.Tx
}

type Option func(*option)

func WithTransaction(tx *sql.Tx) Option {
	return func(o *option) {
		o.tx = tx
	}
}
