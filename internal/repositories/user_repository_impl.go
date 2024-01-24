package repositories

import (
	"context"
	"database/sql"

	"github.com/Zainal21/go-bone/pkg/database/mysql"

	"github.com/Zainal21/go-bone/internal/entity"
)

type userRepositoryImpl struct {
	db mysql.Adapter
}

func (u *userRepositoryImpl) ListUser(ctx context.Context) (*[]entity.User, error) {
	query := `SELECT * FROM users;`
	var result []entity.User

	err := u.db.Query(ctx, &result, query)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r userRepositoryImpl) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, opts)
}

func NewUserRepositoryImpl(db mysql.Adapter) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}
