package bootstrap

import (
	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/pkg/broker"
	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/logger"
)

func RegistryRabbitMQSubscriber(name string, cfg *config.Config, mController contract.MessageController) broker.Subscriber {
	conn, err := broker.ConnectRabbitMQ(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	return broker.NewSubscriber(conn, name, mController)
}

func RegistryRabbitMQPublisher(name string, cfg *config.Config) broker.Publisher {
	conn, err := broker.ConnectRabbitMQ(cfg)
	if err != nil {
		logger.Fatal("dial rabbit mq failed")
	}

	return broker.NewPublisher(conn)
}
