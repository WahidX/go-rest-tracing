package tracing

import (
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

// src:
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go

var (
	lemonsKey = attribute.Key("ex.com/lemons")
)

func initMeter() {
	config := prometheus.Config{}
	c := controller.New(
		processor.New(
			selector.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(config.DefaultHistogramBoundaries),
			),
			export.CumulativeExportKindSelector(),
			processor.WithMemory(true),
		),
	)
	exporter, err := prometheus.New(config, c)
	if err != nil {
		log.Panicf("failed to initialize prometheus exporter %v", err)
	}
	global.SetMeterProvider(exporter.MeterProvider())

	http.HandleFunc("/", exporter.ServeHTTP)
	go func() {
		_ = http.ListenAndServe(":2222", nil)
	}()

	fmt.Println("Prometheus server running on :2222")
}

// func main() {
// 	initMeter()

// 	meter := global.Meter("ex.com/basic")
// 	observerLock := new(sync.RWMutex)
// 	observerValueToReport := new(float64)
// 	observerLabelsToReport := new([]attribute.KeyValue)
// 	cb := func(_ context.Context, result metric.Float64ObserverResult) {
// 		(*observerLock).RLock()
// 		value := *observerValueToReport
// 		labels := *observerLabelsToReport
// 		(*observerLock).RUnlock()
// 		result.Observe(value, labels...)
// 	}
// 	_ = metric.Must(meter).NewFloat64ValueObserver("ex.com.one", cb,
// 		metric.WithDescription("A ValueObserver set to 1.0"),
// 	)

// 	valuerecorder := metric.Must(meter).NewFloat64ValueRecorder("ex.com.two")
// 	counter := metric.Must(meter).NewFloat64Counter("ex.com.three")

// 	commonLabels := []attribute.KeyValue{lemonsKey.Int(10), attribute.String("A", "1"), attribute.String("B", "2"), attribute.String("C", "3")}
// 	notSoCommonLabels := []attribute.KeyValue{lemonsKey.Int(13)}

// 	ctx := context.Background()

// 	(*observerLock).Lock()
// 	*observerValueToReport = 1.0
// 	*observerLabelsToReport = commonLabels
// 	(*observerLock).Unlock()
// 	meter.RecordBatch(
// 		ctx,
// 		commonLabels,
// 		valuerecorder.Measurement(2.0),
// 		counter.Measurement(12.0),
// 	)

// 	time.Sleep(5 * time.Second)

// 	(*observerLock).Lock()
// 	*observerValueToReport = 1.0
// 	*observerLabelsToReport = notSoCommonLabels
// 	(*observerLock).Unlock()
// 	meter.RecordBatch(
// 		ctx,
// 		notSoCommonLabels,
// 		valuerecorder.Measurement(2.0),
// 		counter.Measurement(22.0),
// 	)

// 	time.Sleep(5 * time.Second)

// 	(*observerLock).Lock()
// 	*observerValueToReport = 13.0
// 	*observerLabelsToReport = commonLabels
// 	(*observerLock).Unlock()
// 	meter.RecordBatch(
// 		ctx,
// 		commonLabels,
// 		valuerecorder.Measurement(12.0),
// 		counter.Measurement(13.0),
// 	)

// 	fmt.Println("Example finished updating, please visit :2222")

// 	select {}
// }