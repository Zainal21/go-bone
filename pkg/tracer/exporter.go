package tracer

import (
	"context"
	"errors"
	"fmt"

	"github.com/Zainal21/go-bone/internal/consts"
	"github.com/Zainal21/go-bone/pkg/config"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewExporter(cfg *config.Config) (trace.SpanExporter, error) {
	switch cfg.AppOtelExporter {
	case consts.JaegerExporter:
		return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(fmt.Sprintf("%s/api/traces", cfg.JaegerHost))))
	case consts.TempoExporter:
		return tempoExporter(cfg)
	default:
		return nil, errors.New("unknown otel driver")
	}
}

func tempoExporter(cfg *config.Config) (trace.SpanExporter, error) {
	insecureOpt := otlptracehttp.WithInsecure()
	endpointOpt := otlptracehttp.WithEndpoint(fmt.Sprintf("%s", cfg.TempoHost))

	return otlptracehttp.New(context.Background(), insecureOpt, endpointOpt)
}
