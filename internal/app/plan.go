package app

import (
	"net/http"

	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

type Plan struct {
	Handler *planhttp.PlanHandler
	Store   store.Plan
}

func NewPlan() *Plan {
	store := memory.NewPlanStore()
	return &Plan{
		Handler: planhttp.NewPlanHandler(store),
		Store:   store,
	}
}

func (a *Plan) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/plans", a.Handler.Handle)
}
