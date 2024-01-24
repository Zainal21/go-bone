package repositories

import (
	"context"
	"database/sql"

	"github.com/Zainal21/go-bone/internal/entity"
)

type UserRepository interface {
	ListUser(ctx context.Context) (*[]entity.User, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
