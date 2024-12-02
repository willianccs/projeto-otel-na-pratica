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

// Handle handles the HTTP request
func (h *SubscriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.store.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user model.Subscription
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		createdSubscription, err := h.store.Create(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdSubscription)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
