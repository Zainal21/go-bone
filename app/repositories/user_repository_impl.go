package repositories

import (
	"context"
	"database/sql"

	"github.com/Zainal21/go-bone/app/entity"
	"github.com/Zainal21/go-bone/app/utils/query"
	"github.com/Zainal21/go-bone/pkg/database/mysql"
)

type userRepositoryImpl struct {
	db mysql.Adapter
}

// FindById implements UserRepository.
func (r userRepositoryImpl) FindById(ctx context.Context, id string) (*entity.User, error) {
	_query := query.SelectQuery(
		"users",
		[]string{
			"id",
			"name",
			"email",
			"password",
			"status",
			"role",
			"created_at",
			"updated_at",
		},
		"id = ?",
		1,
		0,
	)

	var result entity.User

	row := r.db.QueryRowX(ctx, _query, id)

	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Email,
		&result.Password,
		&result.Status,
		&result.Role,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

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
