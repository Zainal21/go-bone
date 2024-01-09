package broker

import (
	"fmt"

	"github.com/Zainal21/go-bone/pkg/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func connect(cnf *config.Config) (*amqp.Connection, error) {
	var (
		amqpConfig   = amqp.Config{}
		err          error
		brokerConfig = cnf.BrokerConfig
		amqpConnMode = AMQPMode
	)

	tlsConfig, err := brokerConfig.TlsConfig(cnf.AppEnv)

	if err != nil {
		return nil, err
	}

	if tlsConfig != nil {
		amqpConnMode = AMQPSMode
		amqpConfig.TLSClientConfig = tlsConfig
	}

	dsn := fmt.Sprintf(
		UrlPattern,
		amqpConnMode,
		brokerConfig.RabbitUsername,
		brokerConfig.RabbitPassword,
		brokerConfig.RabbitHost,
		brokerConfig.RabbitPort,
	)
	return amqp.DialConfig(dsn, amqpConfig)
}

func ConnectRabbitMQ(cnf *config.Config) (*amqp.Connection, error) {
	conn, err := connect(cnf)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
