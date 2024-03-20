package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"github.com/Zainal21/go-bone/app/consts"
	"github.com/Zainal21/go-bone/pkg/util"
)

type BrokerConfig struct {
	RabbitEventName string `mapstructure:"rabbitmq_event_name"`
	RabbitQueueName string `mapstructure:"rabbitmq_queue_name"`
	RabbitHost      string `mapstructure:"rabbitmq_host"`
	RabbitPort      int    `mapstructure:"rabbitmq_port"`
	RabbitUsername  string `mapstructure:"rabbitmq_user"`
	RabbitPassword  string `mapstructure:"rabbitmq_pass"`
	TLS             bool   `mapstructure:"rabbitmq_tls"`
	CAPath          string `mapstructure:"rabbitmq_ca_cert"`
	ClientCertPath  string `mapstructure:"rabbitmq_client_cert"`
	ClientKeyPath   string `mapstructure:"rabbitmq_client_key"`
}

func (config *BrokerConfig) TlsConfig(env string) (*tls.Config, error) {
	if !config.TLS {
		return nil, nil
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: util.EnvironmentTransform(env) != consts.AppProduction,
	}

	pool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(config.CAPath)
	if err != nil {
		return nil, err
	}

	if ok := pool.AppendCertsFromPEM(pem); !ok {
		return nil, errors.New("unable to append root cert to pool")
	}

	cert, err := tls.LoadX509KeyPair(config.ClientCertPath, config.ClientKeyPath)
	if err != nil {
		return nil, err
	}

	tlsConfig.Certificates = []tls.Certificate{cert}

	return tlsConfig, nil
}
