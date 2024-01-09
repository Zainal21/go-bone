package config

type DistributeTraceConfig struct {
	JaegerHost string `mapstructure:"jaeger_host"`
	TempoHost  string `mapstructure:"tempo_host"`
}
