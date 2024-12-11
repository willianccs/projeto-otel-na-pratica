// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/cfg"
	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	storegorm "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Payment struct {
	Handler *planhttp.PaymentHandler
	Store   store.Payment
}

func NewPayment(cfg *cfg.Payments) (*Payment, error) {
	db, err := gorm.Open(sqlite.Open(cfg.SQLLite.DSN))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Payment{})

	store := storegorm.NewPaymentStore(db)
	return &Payment{
		Handler: planhttp.NewPaymentHandler(store),
		Store:   store,
	}, nil
}

func (a *Payment) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/payments", a.Handler.Handle)
}
