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
	deleteItemEndpoint := endpoints.MakeEndpoint(svc, "deleteitem")

	// Attach Metrics
	requstTimer := middlewares.NewPrometheusTimer()
	addLockerEndpoint = requstTimer.TimingMetricMiddleware()(addLockerEndpoint)
	addItemEndpoint = requstTimer.TimingMetricMiddleware()(addItemEndpoint)
	getItemEndpoint = requstTimer.TimingMetricMiddleware()(getItemEndpoint)
	deleteItemEndpoint = requstTimer.TimingMetricMiddleware()(deleteItemEndpoint)

	// Create Handlers
	addLockerHandler := httptransport.NewServer(
		addLockerEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)
	addItemToLockerHandler := httptransport.NewServer(
		addItemEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)
	getItemFromLockerHandler := httptransport.NewServer(
		getItemEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)
	deleteFromLockerHandler := httptransport.NewServer(
		deleteItemEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)

	// Expose HTTP endpoints
	{
		http.Handle("/addlocker", addLockerHandler)
		http.Handle("/add", addItemToLockerHandler)
		http.Handle("/get", getItemFromLockerHandler)
		http.Handle("delete", deleteFromLockerHandler)
	}

	// Expose metrics
	http.Handle("/metrics", promhttp.Handler())

	// Start server
	http.ListenAndServe(":8080", nil)
}
