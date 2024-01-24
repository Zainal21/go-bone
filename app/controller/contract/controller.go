package contract

import (
	"github.com/Zainal21/go-bone/app/appctx"
	"github.com/rabbitmq/amqp091-go"
)

type MessageController interface {
	Serve(data amqp091.Delivery) error
}

type Controller interface {
	Serve(xCtx appctx.Data) appctx.Response
}
