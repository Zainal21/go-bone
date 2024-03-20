package broker

import (
	"fmt"

	"github.com/Zainal21/go-bone/app/controller/contract"
	amqp "github.com/rabbitmq/amqp091-go"
)

type subscriber struct {
	conn         *amqp.Connection
	queueName    string
	exchangeName string
	controller   contract.MessageController
}

func NewSubscriber(conn *amqp.Connection, name string, controller contract.MessageController) Subscriber {
	return &subscriber{
		conn:         conn,
		queueName:    fmt.Sprintf("queue_%s", name),
		exchangeName: fmt.Sprintf("exchange_%s", name),
		controller:   controller,
	}
}

func (s *subscriber) Listen(topics []string) error {
	channel, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return s.listenQueue(channel, topics)
}

func (s *subscriber) listenQueue(channel *amqp.Channel, topics []string) error {
	queue, err := declareQueue(channel, s.queueName)
	if err != nil {
		return err
	}

	if err = declareExchange(channel, s.exchangeName); err != nil {
		return err
	}

	for _, topic := range topics {
		if err = channel.QueueBind(
			queue.Name,
			topic,
			s.exchangeName,
			false,
			nil,
		); err != nil {
			return err
		}

		if err = channel.ExchangeBind(
			s.exchangeName,
			topic,
			"exchange_event",
			false,
			nil,
		); err != nil {
			return err
		}
	}

	msg1, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,
	)

	if err != nil {
		return err
	}

	s.consumeMessage(msg1)

	return nil
}

func (s *subscriber) consumeMessage(msgs <-chan amqp.Delivery) {
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			_ = s.controller.Serve(d)
			// Do acknowledgement or something with the payload
		}
	}()
	fmt.Printf("Waiting for message [Exchange, Queue] [%v, %v]\n", s.exchangeName, s.queueName)
	<-forever
}
