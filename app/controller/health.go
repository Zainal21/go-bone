package controller

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/gofiber/fiber/v2"
)

type getHealth struct {
}

func (g *getHealth) Serve(xCtx appctx.Data) appctx.Response {
	// Ping Endpoint
	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("ok").WithData(struct {
		Message string `json:"message"`
	}{
		Message: "Waras!",
	})
}

func NewGetHealth() contract.Controller {
	return &getHealth{}
}
