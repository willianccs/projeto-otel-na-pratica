// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	subscriptionhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

type Subscription struct {
	Handler *subscriptionhttp.SubscriptionHandler
	Store   store.Subscription
}

func NewSubscription(cfg *config.Subscriptions) *Subscription {
	store := memory.NewSubscriptionStore()
	return &Subscription{
		Handler: subscriptionhttp.NewSubscriptionHandler(store, cfg.UsersEndpoint, cfg.PlansEndpoint),
		Store:   store,
	}
}

func (a *Subscription) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /subscriptions", a.Handler.List)
	mux.HandleFunc("POST /subscriptions", a.Handler.Create)
	mux.HandleFunc("GET /subscriptions/{id}", a.Handler.Get)
	mux.HandleFunc("PUT /subscriptions/{id}", a.Handler.Update)
	mux.HandleFunc("DELETE /subscriptions/{id}", a.Handler.Delete)
}
