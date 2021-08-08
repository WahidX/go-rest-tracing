package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func NewRouter(ctx context.Context) http.Handler {
	sleepSpan(ctx, 20)	// dummy trace span
	sleepSpan(ctx, 10)

	// http routers
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", pingHandler)

	return r
}

func sleepSpan(ctx context.Context, val int) {
	// Use the global TracerProvider.
	tr := otel.Tracer("bar-tracer"+fmt.Sprint(val))
	_, span := tr.Start(ctx, "bar-span"+fmt.Sprint(val))
	span.SetAttributes(attribute.Key("sleep-duration").String(fmt.Sprint(val)))	// testset: value => this key-val pair will be visible in jaeger under tags
	span.SetName("span-name")
	defer span.End()

	sleepMs(val)
}

func sleepMs(val int) {
	time.Sleep(time.Duration(val) * time.Millisecond)
	fmt.Println("done")
}