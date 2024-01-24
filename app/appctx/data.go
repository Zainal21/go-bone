package appctx

import (
	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type Data struct {
	FiberCtx *fiber.Ctx
	Cfg      *config.Config
}
