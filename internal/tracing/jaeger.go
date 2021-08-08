package tracing

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// src:
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go

// Environment variable				Option			Default value
// OTEL_EXPORTER_JAEGER_AGENT_HOST	WithAgentHost	localhost
// OTEL_EXPORTER_JAEGER_AGENT_PORT	WithAgentPort	6831
// OTEL_EXPORTER_JAEGER_ENDPOINT	WithEndpoint	http://localhost:14268/api/traces
// OTEL_EXPORTER_JAEGER_USER		WithUsername
// OTEL_EXPORTER_JAEGER_PASSWORD	WithPassword

const (
	service     = "const-srv-name"
	environment = "production"
	id          = 1
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.

// We can export through WithCollectorEndpoint or WithAgentEndpoint
func TracerProvider() (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	url := "http://localhost:14268/api/traces"	// Should be in env var

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

