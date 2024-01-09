package provider

import (
	"context"
	"time"

	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/httpclient"
	"github.com/Zainal21/go-bone/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/sony/gobreaker"
)

type example struct {
	cb  CircuitBreaker
	cfg *config.Config
}

func (e *example) GetTodos(ctx context.Context) ([]byte, error) {
	reqOption := httpclient.RequestOptions{
		Payload:       nil,
		URL:           "http://localhost:8080/api",
		Header:        nil,
		Method:        fiber.MethodGet,
		TimeoutSecond: 4,
		Context:       ctx,
	}
	result, err := e.cb.Execute(ctx, SendRequest, reqOption)
	if err != nil {
		return nil, err
	}
	return result.(httpclient.Response).RawByte(), nil
}

func NewExampleProvider(cnf *config.Config) Example {
	var (
		goBreakerName       = "example-provider"
		goBreakerMaxRequest = 100
		goBreakerInterval   = 20 * time.Second
		goBreakerTimeout    = 5 * time.Second
	)

	settings := gobreaker.Settings{
		Name: goBreakerName,
		// MAXREQUESTS IS THE MAXIMUM NUMBER OF REQUESTS ALLOWED TO PASS THROUGH WHEN THE CIRCUITBREAKER IS HALF-OPEN
		MaxRequests: uint32(goBreakerMaxRequest),

		// INTERVAL IS THE CYCLIC PERIOD OF THE CLOSED STATE FOR CIRCUITBREAKER TO CLEAR THE INTERNAL COUNTS
		Interval: goBreakerInterval,

		// TIMEOUT IS THE PERIOD OF THE OPEN STATE, AFTER WHICH THE STATE OF CIRCUITBREAKER BECOMES HALF-OPEN.
		Timeout: goBreakerTimeout,
		OnStateChange: func(name string, from, to gobreaker.State) {
			var (
				lf = logger.NewFields(
					logger.EventName("ExampleProviderGoBreaker"),
				)
			)
			lf.Append(logger.Any("gobreaker.name", goBreakerName))
			lf.Append(logger.Any("gobreaker.maxRequest", goBreakerMaxRequest))
			lf.Append(logger.Any("gobreaker.timeout", goBreakerInterval))
			lf.Append(logger.Any("gobreaker.timeout", goBreakerTimeout))

			// TO DO CALLBACK GOBREAKER WITH CURRENT STATE
			switch to {
			case gobreaker.StateOpen:
				lf.Append(logger.String("gobreaker.state", "OPEN"))
				logger.ErrorWithContext(context.Background(), "success trigger gobreaker with OPEN state", lf...)
			case gobreaker.StateHalfOpen:
				lf.Append(logger.String("gobreaker.state", "HALF-OPEN"))
				logger.WarnWithContext(context.Background(), "success trigger gobreaker with HALF-OPEN state", lf...)
			case gobreaker.StateClosed:
				lf.Append(logger.String("gobreaker.state", "CLOSED"))
				logger.InfoWithContext(context.Background(), "success trigger gobreaker with CLOSED state", lf...)
			}

		},
		// READYTOTRIP IS CALCULATION MOVE TO OPEN STATE
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests > 2 && failureRatio >= 0.3
		},
	}
	cb := NewCircuitBreaker(cnf, settings)
	return &example{
		cb:  cb,
		cfg: cnf,
	}
}
