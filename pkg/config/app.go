package config

// AppConfig holds the AppConfig configuration
type AppConfig struct {
	AppName string `mapstructure:"app_name"`
	AppEnv  string `mapstructure:"app_env"`
	AppHost string `mapstructure:"app_host"`
	AppPort int    `mapstructure:"app_port"`

	AppLoggerDebug bool   `mapstructure:"app_logger_debug"`
	AppLoggerLevel string `mapstructure:"app_logger_level"`
}
