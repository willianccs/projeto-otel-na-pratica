package memory

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

type inMemoryPayment struct {
	store map[string]model.Payment
}

func NewPaymentStore() store.Payment {
	return &inMemoryPayment{
		store: make(map[string]model.Payment),
	}
}

func (u *inMemoryPayment) Get(_ context.Context, id string) (model.Payment, error) {
	return u.store[id], nil
}

func (u *inMemoryPayment) Create(_ context.Context, user model.Payment) (model.Payment, error) {
	u.store[user.ID] = user
	return user, nil
}

func (u *inMemoryPayment) Update(_ context.Context, user model.Payment) (model.Payment, error) {
	u.store[user.ID] = user
	return user, nil
}

func (u *inMemoryPayment) Delete(_ context.Context, id string) error {
	delete(u.store, id)
	return nil
}

func (u *inMemoryPayment) List(_ context.Context) ([]model.Payment, error) {
	users := make([]model.Payment, 0, len(u.store))
	for _, user := range u.store {
		users = append(users, user)
	}
	return users, nil
}
