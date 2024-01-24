package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"github.com/Zainal21/go-bone/app/consts"
	"github.com/Zainal21/go-bone/pkg/util"
)

// DatabaseConfig holds the DatabaseConfig configuration
type DatabaseConfig struct {
	DBHost          string `mapstructure:"db_host"`
	DBPort          int    `mapstructure:"db_port"`
	DBName          string `mapstructure:"db_name"`
	DBUser          string `mapstructure:"db_user"`
	DBPassword      string `mapstructure:"db_password"`
	MaxOpenConn     int    `mapstructure:"db_max_open_conn"`
	MaxIdleConn     int    `mapstructure:"db_max_idle_conn"`
	MaxConnLifetime int    `mapstructure:"db_conn_lifetime"`
	MaxIdleTime     int    `mapstructure:"db_idle_time"`
	TLS             bool   `mapstructure:"db_tls"`
	CAPath          string `mapstructure:"db_ca_cert"`
	ClientCertPath  string `mapstructure:"db_client_cert"`
	ClientKeyPath   string `mapstructure:"db_client_key"`
}

func (config *DatabaseConfig) TlsConfig(env string) (*tls.Config, error) {
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
