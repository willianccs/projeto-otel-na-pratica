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

func NewSubscription(*config.Subscriptions) *Subscription {
	store := memory.NewSubscriptionStore()
	return &Subscription{
		Handler: subscriptionhttp.NewSubscriptionHandler(store),
		Store:   store,
	}
}

func (a *Subscription) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /subscriptions", a.Handler.ListSubscriptions)
	mux.HandleFunc("POST /subscriptions", a.Handler.CreateSubscription)
	mux.HandleFunc("GET /subscriptions/{id}", a.Handler.GetSubscription)
	mux.HandleFunc("PUT /subscriptions/{id}", a.Handler.UpdateSubscription)
	mux.HandleFunc("DELETE /subscriptions/{id}", a.Handler.DeleteSubscription)
}
