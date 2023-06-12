package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

func createTracingProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-resto"),
			attribute.String("environment", "staging"),
		)),
	)

	return tp, nil
}

func Init(url string) error {
	tp, err := createTracingProvider(url)
	if err != nil {
		return err
	}

	otel.SetTracerProvider(tp)

	return nil
}

func CreateSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}

	tr := otel.Tracer(name)

	ctx, span := tr.Start(ctx, name)

	return ctx, span
}
