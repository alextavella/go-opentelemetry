package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alextavella/go-opentelemetry/internal/handler"
	otel "github.com/alextavella/go-opentelemetry/pkg/otel"
	"github.com/gofiber/fiber/v3"
)

func main() {
	ctx := context.Background()

	//  OpenTelemetry configuration
	otelShutdown, err := otel.SetupOtel(ctx, "example-app")
	if err != nil {
		log.Fatalf("Erro ao configurar OpenTelemetry: %v", err)
	}

	// HTTP Server
	app := fiber.New()
	app.Get("/", handler.HandleRequest)

	// Shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Received shutdown signal, shutting down...")
		if err := otelShutdown(); err != nil {
			log.Fatalf("Error shutting down OpenTelemetry: %v", err)
		}
		log.Println("OpenTelemetry shut down successfully")
		if err := app.Shutdown(); err != nil {
			log.Fatalf("Error shutting down: %v", err)
		}
		log.Println("Server shut down gracefully")
	}()

	// HTTP Server
	log.Println("Server listening on :8080")
	log.Fatal(app.Listen(":8080"))
}
