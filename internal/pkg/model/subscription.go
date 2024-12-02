package model

type Subscription struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PlanID string `json:"plan_id"`
}
