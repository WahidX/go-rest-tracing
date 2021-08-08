package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wahidx/go-rest-sample/internal/tracing"
	"github.com/wahidx/go-rest-sample/web/rest"
	"go.opentelemetry.io/otel"
)

func main() {
	//Tracer init
	tp, err := tracing.TracerProvider()
	if err != nil {
		log.Fatalln(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()		// Check readme for this

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tracer := tp.Tracer("main-tracer")
	
	ctx, span := tracer.Start(ctx, "foo")
	defer span.End()

	router := rest.NewRouter(ctx)

	//check for port
	port := "8000"

	fmt.Println("Listening to port:", port)
	http.ListenAndServe(":"+port, router)
}
