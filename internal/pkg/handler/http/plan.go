// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"encoding/json"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

// PlanHandler is an HTTP handler that performs CRUD operations for model.Plan using a store.Plan
type PlanHandler struct {
	store store.Plan
}

// NewPlanHandler returns a new PlanHandler
func NewPlanHandler(store store.Plan) *PlanHandler {
	return &PlanHandler{
		store: store,
	}
}

// Handle handles the HTTP request
func (h *PlanHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		plans, err := h.store.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(plans)
	case http.MethodPost:
		var plan model.Plan
		if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		createdPlan, err := h.store.Create(r.Context(), &plan)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdPlan)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
