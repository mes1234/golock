package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/endpoints"
	"github.com/mes1234/golock/service"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	svc := service.NewAccessService(logger, requestCount)

	addLockerEndpoint := endpoints.MakeAddLockerEndpoint(svc)

	lockerHandler := httptransport.NewServer(
		addLockerEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)

	http.Handle("/addlocker", lockerHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
