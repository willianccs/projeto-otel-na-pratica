package app

import (
	"net/http"

	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

type Payment struct {
	Handler *planhttp.PaymentHandler
	Store   store.Payment
}

func NewPayment() *Payment {
	store := memory.NewPaymentStore()
	return &Payment{
		Handler: planhttp.NewPaymentHandler(store),
		Store:   store,
	}
}

func (a *Payment) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/payments", a.Handler.Handle)
}
