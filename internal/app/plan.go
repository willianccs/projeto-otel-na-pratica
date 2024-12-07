package app

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/api"
	grpchandler "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/grpc"
	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
	"google.golang.org/grpc"
)

type Plan struct {
	Handler     *planhttp.PlanHandler
	GRPCHandler api.PlanServiceServer
	Store       store.Plan
}

func NewPlan() *Plan {
	store := memory.NewPlanStore()
	return &Plan{
		Handler:     planhttp.NewPlanHandler(store),
		GRPCHandler: grpchandler.NewPlanServer(store),
		Store:       store,
	}
}

func (a *Plan) RegisterRoutes(mux *http.ServeMux, grpcSrv *grpc.Server) {
	mux.HandleFunc("/plans", a.Handler.Handle)
	api.RegisterPlanServiceServer(grpcSrv, a.GRPCHandler)
}
