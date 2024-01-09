package repositories

import (
	"context"
	"database/sql"

	"github.com/Zainal21/go-bone/pkg/database/mysql"

	"github.com/Zainal21/go-bone/internal/entity"
	"github.com/Zainal21/go-bone/pkg/helper"
	"github.com/Zainal21/go-bone/pkg/tracer"
	"github.com/pkg/errors"
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

func (r userRepositoryImpl) Store(ctx context.Context, payload entity.User, opts ...Option) (int, error) {
	var (
		id  int
		err error
		tx  *sql.Tx
	)

	ctx, span := tracer.NewSpan(ctx, "Repo.StoreUser", nil)
	defer span.End()

	opt := &option{}
	for _, f := range opts {
		f(opt)
	}

	if opt.tx != nil {
		tx = opt.tx
	} else {
		tx, err = r.db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			tracer.AddSpanError(span, err)
			return 0, err
		}

		defer func() {
			err = tx.Commit()
			if err != nil {
				tracer.AddSpanError(span, err)
				err = errors.Wrap(err, "failed to commit")
			}
		}()
	}

	query, val, err := helper.StructQueryInsert(payload, "users", "db", false)

	rows, err := tx.QueryContext(
		ctx,
		query,
		val...,
	)
	if err != nil {
		tracer.AddSpanError(span, err)
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			tracer.AddSpanError(span, err)
			return 0, err
		}
	}

	return id, err
}

func NewUserRepositoryImpl(db mysql.Adapter) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}
