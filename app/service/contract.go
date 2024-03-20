package service

import (
	"context"

	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/entity"
)

type UserService interface {
	ListUser(ctx context.Context) (*[]entity.User, error)
	StoreUser(ctx context.Context) appctx.Response
}
