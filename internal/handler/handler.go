package handler

import (
	"github.com/Zainal21/go-bone/internal/appctx"
	"github.com/Zainal21/go-bone/internal/controller/contract"
	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func HttpRequest(xCtx *fiber.Ctx, svc contract.Controller, conf *config.Config) appctx.Response {
	data := appctx.Data{
		FiberCtx: xCtx,
		Cfg:      conf,
	}
	return svc.Serve(data)
}
