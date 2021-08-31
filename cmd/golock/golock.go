package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/endpoints"
	"github.com/mes1234/golock/middlewares"
	"github.com/mes1234/golock/service"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requstTimer := kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: "golock",
		Subsystem: "access",
		Name:      "request_timer",
		Help:      "execution time of request",
	}, fieldKeys)

	svc := service.NewAccessService(logger, requstTimer)

	addLockerEndpoint := endpoints.MakeAddLockerEndpoint(svc)
	addLockerEndpoint = middlewares.TimingMetricMiddleware(*requstTimer)(addLockerEndpoint)

	lockerHandler := httptransport.NewServer(
		addLockerEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)

	http.Handle("/addlocker", lockerHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
