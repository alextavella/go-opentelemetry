package main

import (
	"context"
	"log"
	"net/http"
	"os"

	handler "github.com/alextavella/go-opentelemetry/internal/handler"
	otel "github.com/alextavella/go-opentelemetry/pkg/otel"
)

func main() {
	ctx := context.Background()

	//  OpenTelemetry configuration
	otelShutdown, err := otel.SetupOtel(ctx, "example-app")
	if err != nil {
		log.Fatalf("Erro ao configurar OpenTelemetry: %v", err)
	}

	// Shutdown gracefully on interrupt signal
	shuCh := make(chan os.Signal, 1)
	go func() {
		<-shuCh
		log.Println("Recebido sinal de desligamento, encerrando...")
		if err := otelShutdown(); err != nil {
			log.Fatalf("Erro ao desligar OpenTelemetry: %v", err)
		}
		log.Println("OpenTelemetry desligado com sucesso")
		os.Exit(0)
	}()

	// HTTP Server
	http.HandleFunc("/", handler.HandleRequest)
	log.Println("Servidor ouvindo em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
