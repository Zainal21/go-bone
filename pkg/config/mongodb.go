package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"github.com/Zainal21/go-bone/app/consts"
	"github.com/Zainal21/go-bone/pkg/util"
)

type MongoDBConfig struct {
	DBHost          string `mapstructure:"mongodb_host"`
	DBPort          int    `mapstructure:"mongodb_port"`
	DBName          string `mapstructure:"mongodb_name"`
	DBUser          string `mapstructure:"mongodb_user"`
	DBPassword      string `mapstructure:"mongodb_password"`
	MaxOpenConn     int    `mapstructure:"mongodb_max_open_conn"`
	MaxIdleConn     int    `mapstructure:"mongodb_max_idle_conn"`
	MaxConnLifetime int    `mapstructure:"mongodb_conn_lifetime"`
	MaxIdleTime     int    `mapstructure:"mongodb_idle_time"`
	TLS             bool   `mapstructure:"mongodb_tls"`
	CAPath          string `mapstructure:"mongodb_ca_cert"`
	ClientCertPath  string `mapstructure:"mongodb_client_cert"`
	ClientKeyPath   string `mapstructure:"mongodb_client_key"`
}

func (config *MongoDBConfig) TlsConfig(env string) (*tls.Config, error) {
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
