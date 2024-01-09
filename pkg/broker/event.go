package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareQueue(channel *amqp.Channel, name string) (amqp.Queue, error) {
	return channel.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		amqp.Table{
			"x-queue-type": "quorum",
		},
	)
}

func declareExchange(channel *amqp.Channel, exchangeName string) error {
	return channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		true,
		false,
		false,
		nil,
	)
}
