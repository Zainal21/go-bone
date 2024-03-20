package controller

import (
	"context"
	"fmt"

	"github.com/Zainal21/go-bone/app/controller/contract"
	"github.com/Zainal21/go-bone/pkg/logger"
	"github.com/Zainal21/go-bone/pkg/pubsubx"
)

type pubsubController struct {
}

func (p *pubsubController) Serve(ctx context.Context, message *pubsubx.Message) {
	logger.InfoWithContext(ctx, fmt.Sprintf("Received Data %s", string(message.Data)))
}

func NewPubsubController() contract.PubSubMessageController {
	return &pubsubController{}
}
