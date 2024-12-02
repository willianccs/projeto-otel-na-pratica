package plan

import (
	"net/http"

	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

// App is a plan application, holding a store.Plan and a http.PlanHandler
type App struct {
	PlanHandler *planhttp.PlanHandler
	PlanStore   store.Plan
}

// NewApp returns a new App
func NewApp() *App {
	store := memory.NewPlanStore()
	return &App{
		PlanHandler: planhttp.NewPlanHandler(store),
		PlanStore:   store,
	}
}

// Start starts the application, listening on port 8083
func (a *App) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/plans", a.PlanHandler.Handle)
	http.ListenAndServe(":8082", mux)
}
