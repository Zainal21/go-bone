package config

type CacheConfig struct {
	RedisHost     string `mapstructure:"redis_host"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}
