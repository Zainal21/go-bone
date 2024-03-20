package broker

import (
	"fmt"
	"testing"
	"time"

	"github.com/Zainal21/go-bone/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TestPublish(t *testing.T) {
	logger.SetJSONFormatter()
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatal(fmt.Sprintf("failed dial %v", err))
	}

	publisher := NewPublisher(connection)

	for i := 0; i < 10; i++ {
		err = publisher.Publish("info", MessagePayload{
			Type:    "test",
			Message: "test",
			Data:    "ok",
		})
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < 5; i++ {
		err = publisher.Publish("debug", MessagePayload{
			Type:    "log",
			Message: "log",
			Data:    "log",
		})
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		t.Fatal(err)
	}
}
