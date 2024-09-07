package tracer

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// JaegerTraceProvider creates a new Jaeger trace provider
// with the given application name, environment, and Jaeger endpoint.
func JaegerTraceProvider(appName, appEnv, jaegerEndpoint string) (*sdktrace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)))
	if err != nil {
		return nil, fmt.Errorf("failed to create Jaeger exporter: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
			semconv.DeploymentEnvironmentKey.String(appEnv),
		)),
	)
	return tp, nil
}
