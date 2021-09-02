package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/endpoints"
	"github.com/mes1234/golock/middlewares"
	"github.com/mes1234/golock/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.NewAccessService(logger)

	addLockerEndpoint := endpoints.MakeEndpoint(svc, "addlocker")
	addItemEndpoint := endpoints.MakeEndpoint(svc, "additem")
	getItemEndpoint := endpoints.MakeEndpoint(svc, "getitem")
	removeItemEndpoint := endpoints.MakeEndpoint(svc, "removeitem")

	// Attach Metrics
	requstTimer := middlewares.NewPrometheusTimer()
	addLockerEndpoint = requstTimer.TimingMetricMiddleware()(addLockerEndpoint)
	addItemEndpoint = requstTimer.TimingMetricMiddleware()(addItemEndpoint)
	getItemEndpoint = requstTimer.TimingMetricMiddleware()(getItemEndpoint)
	removeItemEndpoint = requstTimer.TimingMetricMiddleware()(removeItemEndpoint)

	// Create Handlers
	addLockerHandler := httptransport.NewServer(
		addLockerEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)
	addItemToLockerHandler := httptransport.NewServer(
		addItemEndpoint,
		dto.DecodeHttpAddItemRequest,
		dto.EncodeHttpAddItemResponse,
	)
	getItemFromLockerHandler := httptransport.NewServer(
		getItemEndpoint,
		dto.DecodeHttpGetItemRequest,
		dto.EncodeHttpGetItemResponse,
	)
	removeFromLockerHandler := httptransport.NewServer(
		removeItemEndpoint,
		dto.DecodeHttpRemoveItemRequest,
		dto.EncodeHttpRemoveItemResponse,
	)

	// Expose HTTP endpoints
	{
		http.Handle("/addlocker", addLockerHandler)
		http.Handle("/add", addItemToLockerHandler)
		http.Handle("/get", getItemFromLockerHandler)
		http.Handle("/remove", removeFromLockerHandler)
	}

	// Expose metrics
	http.Handle("/metrics", promhttp.Handler())

	// Start server
	http.ListenAndServe(":8080", nil)
}
