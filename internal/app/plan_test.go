package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/api"
	grpchandler "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/grpc"
	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestPlan_RegisterRoutes(t *testing.T) {
	// prepare
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		require.NoError(t, err)
	}
	grpcServer := grpc.NewServer()
	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	mux := http.NewServeMux()
	plan := NewPlan()
	expected := &model.Plan{
		ID:          "123",
		Name:        "Test Plan",
		Description: "This is a test plan",
		Price:       10,
	}
	plan.Store.Create(context.Background(), expected)

	// test
	plan.RegisterRoutes(mux, grpcServer)

	// verify
	{ // http
		req, err := http.NewRequest("GET", "/plans", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var plans []*model.Plan
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &plans))
		assert.Equal(t, expected, plans[0])
	}

	{ // grpc
		conn, err := grpc.NewClient("localhost:8081", grpc.WithInsecure())
		defer conn.Close()

		require.NoError(t, err)
		cl := api.NewPlanServiceClient(conn)
		resp, err := cl.List(context.Background(), &api.ListRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expected.ID, resp.Plans[0].Id)
	}
}

func TestNewPlan(t *testing.T) {
	plan := NewPlan()
	assert.NotNil(t, plan.Handler)
	assert.NotNil(t, plan.GRPCHandler)
	assert.NotNil(t, plan.Store)
}

func TestPlanHandler_Handle(t *testing.T) {
	store := memory.NewPlanStore()
	plan := NewPlan()
	plan.Handler = planhttp.NewPlanHandler(store)

	req, err := http.NewRequest("GET", "/plans", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	plan.Handler.Handle(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGRPCHandler(t *testing.T) {
	store := memory.NewPlanStore()
	plan := NewPlan()
	plan.GRPCHandler = grpchandler.NewPlanServer(store)

	req := &api.ListRequest{}
	resp, err := plan.GRPCHandler.List(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Len(t, resp.Plans, 0)
}
