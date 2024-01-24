package router

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/bootstrap"
	"github.com/Zainal21/go-bone/app/controller"
	"github.com/Zainal21/go-bone/app/controller/auth"
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/app/controller/user"
	"github.com/Zainal21/go-bone/app/handler"
	"github.com/Zainal21/go-bone/app/middleware"
	"github.com/Zainal21/go-bone/app/repositories"
	"github.com/Zainal21/go-bone/app/service"
	cryptoservice "github.com/Zainal21/go-bone/app/utils/crypto"
	"github.com/Zainal21/go-bone/app/utils/sanctum"
	"github.com/Zainal21/go-bone/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	cfg   *config.Config
	fiber fiber.Router
}

func (rtr *router) handle(hfn httpHandlerFunc, svc contract.Controller, mdws ...middleware.MiddlewareFunc) fiber.Handler {
	return func(xCtx *fiber.Ctx) error {

		//check registered middleware functions
		if rm := middleware.FilterFunc(rtr.cfg, xCtx, mdws); rm.Code != fiber.StatusOK {
			// return response base on middleware
			res := *appctx.NewResponse().
				WithCode(rm.Code).
				WithError(rm.Errors).
				WithMessage(rm.Message)
			return rtr.response(xCtx, res)
		}

		//send to controller
		resp := hfn(xCtx, svc, rtr.cfg)
		return rtr.response(xCtx, resp)
	}
}

func (rtr *router) response(fiberCtx *fiber.Ctx, resp appctx.Response) error {
	fiberCtx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return fiberCtx.Status(resp.Code).Send(resp.Byte())
}

func (rtr *router) Route() {
	//init db
	db := bootstrap.RegistryDatabase(rtr.cfg)

	//define repositories
	userRepo := repositories.NewUserRepositoryImpl(db)
	tokenRepo := repositories.NewPersonalToken(db, &sanctum.Token{
		Crypto: &cryptoservice.Crypto{},
	}, userRepo)

	//define services
	userSvc := service.NewUserServiceImpl(userRepo)

	//define middleware
	basicMiddleware := middleware.NewAuthMiddleware()

	//define provider

	//define controller
	getAllUser := user.NewGetAllUser(userSvc)
	signIn := auth.NewSignIn(userSvc, tokenRepo, rtr.cfg)

	health := controller.NewGetHealth()
	privateV1 := rtr.fiber.Group("/api/private/v1")

	rtr.fiber.Get("/ping", rtr.handle(
		handler.HttpRequest,
		health,
	))

	privateV1.Post("/sign-in", rtr.handle(
		handler.HttpRequest,
		signIn,
	))

	privateV1.Get("/users", rtr.handle(
		handler.HttpRequest,
		getAllUser,
		//middleware
		basicMiddleware.Authenticate,
	))
}

func NewRouter(cfg *config.Config, fiber fiber.Router) Router {
	return &router{
		cfg:   cfg,
		fiber: fiber,
	}
}
