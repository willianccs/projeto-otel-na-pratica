// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"encoding/json"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

// SubscriptionHandler is an HTTP handler that performs CRUD operations for model.Subscription using a store.Subscription
type SubscriptionHandler struct {
	store store.Subscription
}

// NewSubscriptionHandler returns a new SubscriptionHandler
func NewSubscriptionHandler(store store.Subscription) *SubscriptionHandler {
	return &SubscriptionHandler{
		store: store,
	}
}

func (h *SubscriptionHandler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subscriptions, err := h.store.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	subscription := &model.Subscription{}
	if err := json.NewDecoder(r.Body).Decode(subscription); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdSubscription, err := h.store.Create(r.Context(), subscription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(createdSubscription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (h *SubscriptionHandler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	subscription, err := h.store.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if subscription == nil {
		http.Error(w, "Subscription not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(subscription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SubscriptionHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
}

func (h *SubscriptionHandler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
}
