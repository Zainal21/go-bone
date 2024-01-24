package user

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/app/service"
	"github.com/gofiber/fiber/v2"
)

type storeUser struct {
	service service.UserService
}

func (g *storeUser) Serve(xCtx appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(fiber.StatusNotFound).WithMessage("Resource Not Found")
}

func NewStoreUser(svc service.UserService) contract.Controller {
	return &storeUser{service: svc}
}
