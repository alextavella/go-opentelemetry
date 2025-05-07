package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func setupOtelMetrics(
	ctx context.Context,
	res *resource.Resource,
) (func(ctx context.Context) error, error) {
	metricOpts := []otlpmetricgrpc.Option{otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint("otel-collector:4317")}
	metricExporter, err := otlpmetricgrpc.New(ctx, metricOpts...)
	if err != nil {
		return nil, err
	}
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter, sdkmetric.WithInterval(1))),
	)
	otel.SetMeterProvider(mp)
	return mp.Shutdown, nil
}
