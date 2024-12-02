package memory

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

type inMemoryPlan struct {
	store map[string]model.Plan
}

func NewPlanStore() store.Plan {
	return &inMemoryPlan{
		store: make(map[string]model.Plan),
	}
}

func (u *inMemoryPlan) Get(_ context.Context, id string) (model.Plan, error) {
	return u.store[id], nil
}

func (u *inMemoryPlan) Create(_ context.Context, user model.Plan) (model.Plan, error) {
	u.store[user.ID] = user
	return user, nil
}

func (u *inMemoryPlan) Update(_ context.Context, user model.Plan) (model.Plan, error) {
	u.store[user.ID] = user
	return user, nil
}

func (u *inMemoryPlan) Delete(_ context.Context, id string) error {
	delete(u.store, id)
	return nil
}

func (u *inMemoryPlan) List(_ context.Context) ([]model.Plan, error) {
	users := make([]model.Plan, 0, len(u.store))
	for _, user := range u.store {
		users = append(users, user)
	}
	return users, nil
}
