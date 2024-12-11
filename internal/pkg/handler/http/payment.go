// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// PaymentHandler is an HTTP handler that performs CRUD operations for model.Payment using a store.Payment
type PaymentHandler struct {
	store     store.Payment
	js        jetstream.JetStream
	jsSubject string
}

// NewPaymentHandler returns a new PaymentHandler
func NewPaymentHandler(store store.Payment, js jetstream.JetStream, jsSubject string) *PaymentHandler {
	return &PaymentHandler{
		store:     store,
		js:        js,
		jsSubject: jsSubject,
	}
}

// Handle handles the HTTP request
func (h *PaymentHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.store.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var payment model.Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		payload, err := json.Marshal(payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = h.js.PublishMsgAsync(&nats.Msg{
			Subject: h.jsSubject,
			Data:    payload,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(payment)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PaymentHandler) OnMessage(msg jetstream.Msg) {
	var payment model.Payment
	err := json.Unmarshal(msg.Data(), &payment)
	if err != nil {
		return
	}

	_, err = h.store.Create(context.Background(), payment)
	if err != nil {
		return
	}

	_ = msg.Ack()
}
