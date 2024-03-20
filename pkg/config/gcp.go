package config

type GCPConfig struct {
	PubsubAccountPath string `mapstructure:"pubsub_account_path"`
	ProjectID         string `mapstructure:"project_id"`
}
