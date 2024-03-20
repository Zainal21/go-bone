package broker

const (
	AMQPMode  = "amqp"
	AMQPSMode = "amqps"

	UrlPattern = "%s://%s:%s@%s:%d/"
)

type Subscriber interface {
	Listen(topics []string) error
}

type Publisher interface {
	Publish(route string, payload MessagePayload) error
}
