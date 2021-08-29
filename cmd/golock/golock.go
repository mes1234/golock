package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/endpoints"
	"github.com/mes1234/golock/internal/middlewares"
	"github.com/mes1234/golock/service"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.AccessService{}

	addLockerEndpoint := endpoints.MakeAddLockerEndpoint(svc)
	addLockerEndpoint = middlewares.LoggingMiddleware(log.With(logger, "method", "addLockerEndpoint"))(addLockerEndpoint)

	lockerHandler := httptransport.NewServer(
		addLockerEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)

	http.Handle("/addlocker", lockerHandler)
	http.ListenAndServe(":8080", nil)
}
