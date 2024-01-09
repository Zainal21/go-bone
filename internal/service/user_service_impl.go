package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Zainal21/go-bone/internal/appctx"
	"github.com/Zainal21/go-bone/internal/entity"
	"github.com/Zainal21/go-bone/internal/repositories"
	"github.com/Zainal21/go-bone/pkg/logger"
	"github.com/Zainal21/go-bone/pkg/tracer"
)

type userServiceImpl struct {
	repo repositories.UserRepository
}

func (u userServiceImpl) ListUser(ctx context.Context) (*[]entity.User, error) {
	return u.repo.ListUser(ctx)
}

func (u userServiceImpl) StoreUser(ctx context.Context) appctx.Response {
	var (
		lf = logger.NewFields(
			logger.EventName("ServiceStoreUser"),
		)
	)
	ctx, span := tracer.NewSpan(ctx, "Service.StoreUser", nil)
	defer span.End()

	payload := entity.User{
		ID:   1,
		Name: "John Doe",
		Age:  12,
	}

	// value custom logger
	lf.Append(logger.Any("payload.id", payload.ID))
	lf.Append(logger.Any("payload.name", payload.Name))
	lf.Append(logger.Any("payload.age", payload.Age))

	// start db transaction
	tx, err := u.repo.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		tracer.AddSpanError(span, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("start db transaction got error: %v", err), lf...)
		return *appctx.NewResponse().WithCode(http.StatusInternalServerError).WithMessage("Somethin went wrong")
	}

	txOpt := repositories.WithTransaction(tx)

	_, err = u.repo.Store(ctx, entity.User{
		Name: "John Doe",
		Age:  12,
	}, txOpt)
	if err != nil {
		tracer.AddSpanError(span, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("store user got error: %v", err), lf...)

		// rollback transaction
		tx.Rollback()
		return *appctx.NewResponse().WithCode(http.StatusInternalServerError).WithMessage("Somethin went wrong")
	}

	logger.InfoWithContext(ctx, "success store user", lf...)

	// commit transaction
	tx.Commit()
	return *appctx.NewResponse().WithCode(http.StatusCreated).WithMessage("Success created user").WithData(
		map[string]interface{}{
			"user_name":  payload.Name,
			"created_at": time.Now(),
		},
	)

}

func NewUserServiceImpl(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
