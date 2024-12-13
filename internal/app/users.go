// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	userhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
)

type User struct {
	Handler *userhttp.UserHandler
	Store   store.User
}

func NewUser(*config.Users) *User {
	store := memory.NewUserStore()
	return &User{
		Handler: userhttp.NewUserHandler(store),
		Store:   store,
	}
}

func (a *User) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", a.Handler.List)
	mux.HandleFunc("POST /users", a.Handler.Create)
	mux.HandleFunc("GET /users/{id}", a.Handler.Get)
	mux.HandleFunc("PUT /users/{id}", a.Handler.Update)
	mux.HandleFunc("DELETE /users/{id}", a.Handler.Delete)
}
