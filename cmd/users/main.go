// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
)

func main() {
	// Set up logging
	logger := log.Default()

	// Initialize OpenTelemetry
	shutdown := initTracer()
	defer shutdown()

	// Set up flag for config file path
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()

	// Load configuration
	c, err := config.LoadConfig(*configFlag)
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize application
	a := app.NewUser(&c.Users)

	// Register routes
	a.RegisterRoutes(http.DefaultServeMux)

	// Log the starting server message with the configured port
	port := extractPort(c.Server.Endpoint.HTTP)
	logger.Printf("Starting server on port %s...", port)

	// Start the server
	if err := http.ListenAndServe(c.Server.Endpoint.HTTP, http.DefaultServeMux); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func initTracer() func() {
	// Create a stdout trace exporter to output spans to the console
	exporter, _ := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)

	// Create the trace provider with a batch span processor
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)

	// Set the global tracer provider
	otel.SetTracerProvider(tp)

	// Return a shutdown function to clean up resources
	return func() {
		_ = tp.Shutdown(context.Background())
	}
}

func extractPort(endpoint string) string {
	// Extract the port from the endpoint string. Assumes the endpoint is in the format "host:port"
	hostPort := strings.Split(endpoint, ":")
	if len(hostPort) > 1 {
		return hostPort[1]
	}
	return "unknown"
}
