package controller

import (
	"encoding/json"
	"fmt"

	"github.com/Zainal21/go-bone/internal/controller/contract"
	"github.com/Zainal21/go-bone/pkg/broker"
	"github.com/Zainal21/go-bone/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type logController struct {
}

func (l logController) Serve(data amqp.Delivery) error {
	var (
		lf = logger.NewFields(logger.EventName("logProcessor"))
	)
	var payload broker.MessagePayload
	_ = json.Unmarshal(data.Body, &payload)

	logger.Info(fmt.Sprintf("Payload Data %+v)", payload), lf...)

	return nil
}

func NewLogController() contract.MessageController {
	return &logController{}
}
