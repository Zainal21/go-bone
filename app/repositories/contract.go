package repositories

import (
	"context"
	"database/sql"

	"github.com/Zainal21/go-bone/app/entity"
)

type UserRepository interface {
	ListUser(ctx context.Context) (*[]entity.User, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Store(ctx context.Context, payload entity.User, opts ...Option) (int, error)
}
