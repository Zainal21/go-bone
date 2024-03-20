package router

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/bootstrap"
	"github.com/Zainal21/go-bone/app/controller"
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/app/controller/todo"
	"github.com/Zainal21/go-bone/app/controller/user"
	"github.com/Zainal21/go-bone/app/handler"
	"github.com/Zainal21/go-bone/app/middleware"
	"github.com/Zainal21/go-bone/app/provider"
	"github.com/Zainal21/go-bone/app/repositories"
	"github.com/Zainal21/go-bone/app/service"
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

	//define services
	userSvc := service.NewUserServiceImpl(userRepo)

	//define middleware
	basicMiddleware := middleware.NewAuthMiddleware()

	//define provider
	example := provider.NewExampleProvider(rtr.cfg)

	//define controller
	getAllUser := user.NewGetAllUser(userSvc)
	storeUser := user.NewStoreUser(userSvc)
	getTodos := todo.NewGetTodo(example)

	health := controller.NewGetHealth()
	internalV1 := rtr.fiber.Group("/api/internal/v1")

	rtr.fiber.Get("/ping", rtr.handle(
		handler.HttpRequest,
		health,
	))

	internalV1.Get("/users", rtr.handle(
		handler.HttpRequest,
		getAllUser,
		//middleware
		basicMiddleware.Authenticate,
	))

	internalV1.Post("/users", rtr.handle(
		handler.HttpRequest,
		storeUser,
		//middleware
		// basicMiddleware.Authenticate,
	))

	internalV1.Get("/todos", rtr.handle(
		handler.HttpRequest,
		getTodos,
		//middleware
		// basicMiddleware.Authenticate,
	))

}

func NewRouter(cfg *config.Config, fiber fiber.Router) Router {
	return &router{
		cfg:   cfg,
		fiber: fiber,
	}
}
