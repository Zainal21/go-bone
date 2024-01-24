package todo

import (
	"encoding/json"

	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/app/provider"
	"github.com/gofiber/fiber/v2"
)

type todo struct {
	provider provider.Example
}

func (t todo) Serve(xCtx appctx.Data) appctx.Response {
	todos, err := t.provider.GetTodos(xCtx.FiberCtx.Context())
	if err != nil {
		if err != nil {
			return *appctx.NewResponse().WithError(map[string]interface{}{
				"Message": "PROVIDER_ERR",
				"Error":   []string{err.Error()},
			}).WithMessage("Get Data Failed").WithCode(fiber.StatusBadRequest)
		}
	}

	var result = make([]map[string]interface{}, 0)

	json.Unmarshal(todos, &result)

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithData(result)
}

func NewGetTodo(provider provider.Example) contract.Controller {
	return &todo{provider: provider}
}
