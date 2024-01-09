package router

import (
	"github.com/Zainal21/go-bone/internal/appctx"
	"github.com/Zainal21/go-bone/internal/controller/contract"
	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type httpHandlerFunc func(xCtx *fiber.Ctx, svc contract.Controller, conf *config.Config) appctx.Response

type Router interface {
	Route()
}
