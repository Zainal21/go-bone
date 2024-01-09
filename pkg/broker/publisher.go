package broker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Zainal21/go-bone/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type publisher struct {
	connection   *amqp.Connection
	exchangeName string
}

func NewPublisher(conn *amqp.Connection) Publisher {
	p := &publisher{
		connection:   conn,
		exchangeName: "exchange_event", // when publishing we always pub to this exchange
	}

	if err := p.setup(); err != nil {
		logger.Fatal("publisher setup not initialized")
	}

	return p
}

func (p *publisher) setup() error {
	ch, err := p.connection.Channel()
	if err != nil {
		return err
	}

	return declareExchange(ch, p.exchangeName)
}

func (p *publisher) Publish(route string, payload MessagePayload) error {
	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	dataPublish, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = channel.PublishWithContext(
		context.Background(),
		p.exchangeName,
		route,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         dataPublish,
			DeliveryMode: 2,
			Timestamp:    time.Now(),
			AppId:        "testing-app-id",
		},
	)
	if err != nil {
		return err
	}

	return nil
}
