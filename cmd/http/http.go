package http

import (
	"context"
	"fmt"

	"github.com/Zainal21/go-bone/app/bootstrap"
	"github.com/Zainal21/go-bone/app/controller"
	"github.com/Zainal21/go-bone/pkg/pubsubx"

	"github.com/Zainal21/go-bone/pkg/app"
	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/logger"
)

func Start() {
	logger.SetJSONFormatter()
	cnf, err := config.LoadAllConfigs()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load configuration file: %v", err))
	}

	app.InitializeApp(cnf)
	application := app.GetServer()

	//work on pubsub
	messageController := controller.NewPubsubController()
	go func() {
		subscriberer := bootstrap.RegistryPubSubConsumer(cnf)
		err := subscriberer.Subscribe(context.Background(), messageController.Serve, pubsubx.WithTopic("test-pub-provider-sub"), pubsubx.WithSubscribeAsync(true), pubsubx.WithMaxConcurrent(1))
		if err != nil {
			logger.Fatal(err)
		}
	}()

	if err := application.StartServer(); err != nil {
		logger.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}
