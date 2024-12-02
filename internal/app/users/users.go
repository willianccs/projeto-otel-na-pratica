package users

import (
	"net/http"

	userhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

// App is a users application, holding a store.User and a http.UserHandler
type App struct {
	UserHandler *userhttp.UserHandler
	UserStore   store.User
}

// NewApp returns a new App
func NewApp() *App {
	store := memory.NewUserStore()
	return &App{
		UserHandler: userhttp.NewUserHandler(store),
		UserStore:   store,
	}
}

// Start starts the application, listening on port 8081
func (a *App) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", a.UserHandler.Handle)
	http.ListenAndServe(":8081", mux)
}
