package main

import (
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mes1234/golock/auth"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/endpoints"
	"github.com/mes1234/golock/middlewares"
	"github.com/mes1234/golock/persistance"
	"github.com/mes1234/golock/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.NewAccessService(logger)
	tokenService := service.NewTokenService(logger)

	go persistance.Run()

	addLockerEndpoint := endpoints.MakeEndpoint(svc, "addlocker")
	addItemEndpoint := endpoints.MakeEndpoint(svc, "additem")
	getItemEndpoint := endpoints.MakeEndpoint(svc, "getitem")
	removeItemEndpoint := endpoints.MakeEndpoint(svc, "removeitem")

	tokenEndpoint := endpoints.MakeTokenEndpoint(tokenService)

	// Attach Metrics
	requstTimer := middlewares.NewPrometheusTimer()

	addLockerEndpoint = requstTimer.TimingMetricMiddleware()(addLockerEndpoint)
	addItemEndpoint = requstTimer.TimingMetricMiddleware()(addItemEndpoint)
	getItemEndpoint = requstTimer.TimingMetricMiddleware()(getItemEndpoint)
	removeItemEndpoint = requstTimer.TimingMetricMiddleware()(removeItemEndpoint)
	tokenEndpoint = requstTimer.TimingMetricMiddleware()(tokenEndpoint)

	// Attach  Authorization

	addLockerEndpoint = middlewares.AuthorizationMiddleware(logger)(addLockerEndpoint)
	// Attach Authentication

	addLockerEndpoint = gokitjwt.NewParser(auth.Keys, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(addLockerEndpoint)
	jwtOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(gokitjwt.HTTPToContext()),
	}

	// Create Handlers
	addLockerHandler := httptransport.NewServer(
		addLockerEndpoint,
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
		jwtOptions...,
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

	tokenHandler := httptransport.NewServer(
		tokenEndpoint,
		dto.DecodeHttpGetTokenRequest,
		dto.EncodeHttpGetTokenResponse,
	)

	// Expose HTTP endpoints
	{
		http.Handle("/addlocker", addLockerHandler)
		http.Handle("/add", addItemToLockerHandler)
		http.Handle("/get", getItemFromLockerHandler)
		http.Handle("/remove", removeFromLockerHandler)
		http.Handle("/token", tokenHandler)
	}

	// Expose metrics
	http.Handle("/metrics", promhttp.Handler())

	// Start server
	http.ListenAndServe(":8080", nil)
}
