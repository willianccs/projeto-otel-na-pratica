package model

type Payment struct {
	ID             string  `json:"id"`
	SubscriptionID string  `json:"subscription_id"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
}
