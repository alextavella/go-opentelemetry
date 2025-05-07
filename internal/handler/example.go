package handler

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	infra_otel "github.com/alextavella/go-opentelemetry/internal/infra/otel"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	meter := infra_otel.GetMeter()
	requestCounter, err := meter.Int64Counter("request.counter")
	if err != nil {
		fmt.Printf("error creating counter: %v\n", err)
	}

	tracer := infra_otel.GetTracer()
	ctx, span := tracer.Start(r.Context(), "handle-request")
	defer span.End()

	requestCounter.Add(ctx, 1)
	log.Println("Recebida requisição")

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	fmt.Fprintln(w, "Hello, OpenTelemetry!")
}
