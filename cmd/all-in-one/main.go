package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"google.golang.org/grpc"
)

func main() {
	mux := http.NewServeMux()

	// starts the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	{
		a := app.NewUser()
		a.RegisterRoutes(mux)
	}

	{
		a := app.NewPlan()
		a.RegisterRoutes(mux, grpcServer)
	}

	{
		a, err := app.NewPayment()
		if err != nil {
			panic(err)
		}
		a.RegisterRoutes(mux)
	}

	{
		a := app.NewSubscription()
		a.RegisterRoutes(mux)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.ListenAndServe(":8080", mux)
}
