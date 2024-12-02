package subscriptions

import (
	"net/http"

	subscriptionhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

// App is a subscriptions application, holding a store.Subscription and a http.SubscriptionHandler
type App struct {
	SubscriptionHandler *subscriptionhttp.SubscriptionHandler
	SubscriptionStore   store.Subscription
}

// NewApp returns a new App
func NewApp() *App {
	store := memory.NewSubscriptionStore()
	return &App{
		SubscriptionHandler: subscriptionhttp.NewSubscriptionHandler(store),
		SubscriptionStore:   store,
	}
}

// Start starts the application, listening on port 8083
func (a *App) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/subscriptions", a.SubscriptionHandler.Handle)
	http.ListenAndServe(":8083", mux)
}
