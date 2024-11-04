package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Subscription struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PlanID string `json:"plan_id"`
}

var subscriptions = make(map[string]Subscription)

func main() {
	http.HandleFunc("/subscriptions", handleSubscriptions)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func handleSubscriptions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(subscriptions)
	case http.MethodPost:
		var subscription Subscription
		if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		subscriptions[subscription.ID] = subscription
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
