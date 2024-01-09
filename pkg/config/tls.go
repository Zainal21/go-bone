package config

import "crypto/tls"

type TLS interface {
	TlsConfig(env string) (*tls.Config, error)
}
