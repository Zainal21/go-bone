package user

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/app/service"
	"github.com/gofiber/fiber/v2"
)

type getAllUser struct {
	service service.UserService
}

func (g *getAllUser) Serve(xCtx appctx.Data) appctx.Response {
	ctx := xCtx.FiberCtx.Context()
	users, err := g.service.ListUser(ctx)
	if err != nil {
		return *appctx.NewResponse().WithError([]appctx.ErrorResp{
			{
				Key:      "PROVIDER_ERR",
				Messages: []string{err.Error()},
			},
		}).WithMessage(err.Error()).WithCode(fiber.StatusBadRequest)
	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("Success").WithData(users)
}

func NewGetAllUser(svc service.UserService) contract.Controller {
	return &getAllUser{service: svc}
}
