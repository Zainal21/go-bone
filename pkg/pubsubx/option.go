package pubsubx

type Option func(cfg *config)

type config struct {
	MaxConcurrent  int
	SubscribeAsync bool
	Topic          string
}

func defaults() *config {
	return &config{
		SubscribeAsync: false,
	}
}

func WithTopic(v string) Option {
	return func(c *config) {
		c.Topic = v
	}
}

func WithMaxConcurrent(v int) Option {
	return func(c *config) {
		c.MaxConcurrent = v
	}
}

func WithSubscribeAsync(v bool) Option {
	return func(c *config) {
		c.SubscribeAsync = v
	}
}
