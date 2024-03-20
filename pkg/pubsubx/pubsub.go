package pubsubx

import "context"

type Subscriberer interface {
	Subscribe(ctx context.Context, handler func(context.Context, *Message), opts ...Option) error
}

type Publisher interface {
	Publish(ctx context.Context, msg *Message) error
}

type Message struct {
	ID        string
	Topic     string
	Data      []byte
	Attribute map[string]string
}
