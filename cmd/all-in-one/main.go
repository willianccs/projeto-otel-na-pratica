// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/telemetry"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/log/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"net"
	"net/http"
	"os"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"google.golang.org/grpc"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	otelConfigFlag := flag.String("otel", "otel.yml", "path to the OTel config file")
	flag.Parse()

	closer, err := telemetry.Setup(context.Background(), *otelConfigFlag)
	if err != nil {
		fmt.Printf("failed to setup telemetry: %v\n", err)
	}
	defer closer(context.Background())

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		otelzap.NewCore("all-in-one", otelzap.WithLoggerProvider(global.GetLoggerProvider())),
	)
	logger := zap.New(core)

	logger.Info("starting the server")
	c, err := config.LoadConfig(*configFlag)
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	mux := http.NewServeMux()

	// starts the gRPC server
	lis, err := net.Listen("tcp", c.Server.Endpoint.GRPC)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	{
		logger.Info("registering the user service")
		a := app.NewUser(&c.Users)
		a.RegisterRoutes(mux)
	}

	{
		logger.Info("registering the plans service")
		a := app.NewPlan(&c.Plans)
		a.RegisterRoutes(mux, grpcServer)
	}

	{
		logger.Info("registering the payments service")
		a, err := app.NewPayment(&c.Payments)
		if err != nil {
			panic(err)
		}
		a.RegisterRoutes(mux)
		defer func() {
			_ = a.Shutdown()
		}()
	}

	{
		logger.Info("registering the subscriptions service")
		a := app.NewSubscription(&c.Subscriptions)
		a.RegisterRoutes(mux)
	}

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			logger.Fatal("failed to serve gRPC", zap.Error(err))
		}
	}()

	err = http.ListenAndServe(c.Server.Endpoint.HTTP, mux)
	if err != nil {
		logger.Fatal("failed to serve HTTP", zap.Error(err))
	}

	logger.Info("stopping the all-in-one service")
}
