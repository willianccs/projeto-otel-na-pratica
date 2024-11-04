package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payment struct {
	ID            string  `json:"id"`
	SubscriptionID string  `json:"subscription_id"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
}

var payments = make(map[string]Payment)

func main() {
	http.HandleFunc("/payments", handlePayments)
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func handlePayments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(payments)
	case http.MethodPost:
		var payment Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		payments[payment.ID] = payment
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
