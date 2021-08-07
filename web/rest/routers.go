package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/wahidx/go-rest-sample/internal/tracing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func NewRouter() http.Handler {
	tp, err := tracing.TracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tracer := tp.Tracer("global-tracer")

	ctx, span := tracer.Start(ctx, "foo")
	defer span.End()

	bar(ctx)	// dummy trace span

	// http routers
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", pingHandler)

	return r
}

func bar(ctx context.Context) {
	// Use the global TracerProvider.
	tr := otel.Tracer("global-tracer")
	_, span := tr.Start(ctx, "bar")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	// Do bar...
}
