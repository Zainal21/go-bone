package app

import (
	"fmt"
	"strings"

	"github.com/Zainal21/go-bone/app/router"

	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type App struct {
	*fiber.App
	Cfg *config.Config
}

var appServer *App

func InitializeApp(cfg *config.Config) {
	// boostrap run and initialize package dependency
	f := fiber.New(cfg.FiberConfig())
	f.Use(
		cors.New(cors.Config{
			MaxAge: 300,
			AllowOrigins: strings.Join([]string{
				"http://*",
				"https://*",
			}, ","),
			AllowHeaders: strings.Join([]string{
				"Origin",
				"Content-Type",
				"Accept",
			}, ","),
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodPut,
				fiber.MethodDelete,
				fiber.MethodHead,
			}, ","),
		}),
		requestid.New(requestid.Config{
			ContextKey: "refid",
			Header:     "X-Reference-Id",
			Generator: func() string {
				return uuid.New().String()
			},
		}),
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}),
	)

	rtr := router.NewRouter(cfg, f)
	rtr.Route()

	appServer = &App{
		App: f,
		Cfg: cfg,
	}
}

func (app *App) StartServer() (err error) {
	return app.Listen(fmt.Sprintf("%v:%v", app.Cfg.AppHost, app.Cfg.AppPort))
}

func GetServer() *App {
	return appServer
}
