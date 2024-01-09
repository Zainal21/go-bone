package user

import (
	"github.com/Zainal21/go-bone/internal/appctx"
	"github.com/Zainal21/go-bone/internal/controller/contract"
	"github.com/Zainal21/go-bone/internal/service"
	"github.com/Zainal21/go-bone/pkg/tracer"
)

type storeUser struct {
	service service.UserService
}

func (g *storeUser) Serve(xCtx appctx.Data) appctx.Response {
	ctx, span := tracer.NewSpan(xCtx.FiberCtx.Context(), "Controller.CreateUser", nil)
	defer span.End()

	res := g.service.StoreUser(ctx)
	return res
}

func NewStoreUser(svc service.UserService) contract.Controller {
	return &storeUser{service: svc}
}
