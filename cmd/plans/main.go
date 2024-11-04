// plans_service/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Plan struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var plans = make(map[string]Plan)

func main() {
	http.HandleFunc("/plans", handlePlans)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handlePlans(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(plans)
	case http.MethodPost:
		var plan Plan
		if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		plans[plan.ID] = plan
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
