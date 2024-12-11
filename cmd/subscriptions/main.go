// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/cfg"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()

	c, err := cfg.LoadConfig(*configFlag)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	a := app.NewSubscription(&c.Subscriptions)
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8083", http.DefaultServeMux)
}
