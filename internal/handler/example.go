package handler

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	otel "github.com/alextavella/go-opentelemetry/pkg/otel"
	"github.com/gofiber/fiber/v3"
)

func HandleRequest(c fiber.Ctx) error {
	meter := otel.GetMeter()
	requestCounter, err := meter.Int64Counter("request.counter")
	if err != nil {
		return fmt.Errorf("error creating counter: %v\n", err)
	}

	tracer := otel.GetTracer()
	ctx, span := tracer.Start(c.Context(), "handle-request")
	defer span.End()

	requestCounter.Add(ctx, 1)
	log.Println("Recebida requisição")

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	return c.SendString("Hello, OpenTelemetry!")
}
