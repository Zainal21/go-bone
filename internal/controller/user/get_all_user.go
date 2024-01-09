package user

import (
	"github.com/Zainal21/go-bone/internal/appctx"
	"github.com/Zainal21/go-bone/internal/controller/contract"
	"github.com/Zainal21/go-bone/internal/service"
	"github.com/gofiber/fiber/v2"
)

type getAllUser struct {
	service service.UserService
}

func (g *getAllUser) Serve(xCtx appctx.Data) appctx.Response {
	ctx := xCtx.FiberCtx.Context()
	users, err := g.service.ListUser(ctx)
	if err != nil {
		return *appctx.NewResponse().WithError(map[string]interface{}{
			"Message": "PROVIDER_ERR",
			"Error":   []string{err.Error()},
		}).WithMessage(err.Error()).WithCode(fiber.StatusBadRequest)
	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("Success").WithData(users)
}

func NewGetAllUser(svc service.UserService) contract.Controller {
	return &getAllUser{service: svc}
}
