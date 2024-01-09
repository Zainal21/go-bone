package service

import (
	"context"

	"github.com/Zainal21/go-bone/internal/appctx"
	"github.com/Zainal21/go-bone/internal/entity"
)

type UserService interface {
	ListUser(ctx context.Context) (*[]entity.User, error)
	StoreUser(ctx context.Context) appctx.Response
}
