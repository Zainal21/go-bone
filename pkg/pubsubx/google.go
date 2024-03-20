package pubsubx

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Zainal21/go-bone/pkg/logger"

	"cloud.google.com/go/pubsub"
)

type gPublisher struct {
	client *pubsub.Client
}

// NewGPublisher create new instance of pubsubx publisher
func NewGPublisher(client *pubsub.Client) *gPublisher {
	return &gPublisher{client: client}
}

// Publish publish message to the topic
func (p *gPublisher) Publish(ctx context.Context, msg *Message) error {
	tp := p.client.Topic(msg.Topic)

	payload := &pubsub.Message{
		Data:        msg.Data,
		Attributes:  msg.Attribute,
		PublishTime: time.Now(),
	}
	result := tp.Publish(ctx, payload)

	//_ = result
	_, err := result.Get(ctx)
	return err
}

type gSubscriber struct {
	client *pubsub.Client
}

// NewGSubscriber create new instance of pubsubx subscriber
func NewGSubscriber(client *pubsub.Client) *gSubscriber {
	return &gSubscriber{client: client}
}

// Subscribe publish message to the topic
func (p *gSubscriber) Subscribe(ctx context.Context, handler func(context.Context, *Message), opts ...Option) error {
	defer p.client.Close()

	cfg := defaults()
	for _, fn := range opts {
		fn(cfg)
	}

	sub := p.client.Subscription(cfg.Topic)
	sub.ReceiveSettings.Synchronous = cfg.SubscribeAsync
	sub.ReceiveSettings.MaxOutstandingMessages = cfg.MaxConcurrent

	var received int32
	err := sub.Receive(ctx, func(xCtx context.Context, msg *pubsub.Message) {
		handler(xCtx, &Message{
			ID:   msg.ID,
			Data: msg.Data,
		})
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}

	logger.Info(fmt.Sprintf("%s received %d messages\n", cfg.Topic, received))
	return nil

}
