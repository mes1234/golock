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
	"github.com/mes1234/golock/service"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.NewAccessService(logger)
	tokenService := service.NewTokenService(logger)

	addLockerEndpoint := endpoints.MakeEndpoint(svc, "addlocker")
	addItemEndpoint := endpoints.MakeEndpoint(svc, "additem")
	getItemEndpoint := endpoints.MakeEndpoint(svc, "getitem")
	removeItemEndpoint := endpoints.MakeEndpoint(svc, "removeitem")

	tokenEndpoint := endpoints.MakeTokenEndpoint(tokenService)

	// Attach  Authorization

	addLockerEndpoint = middlewares.AuthorizationMiddleware(logger)(addLockerEndpoint)
	addItemEndpoint = middlewares.AuthorizationMiddleware(logger)(addItemEndpoint)
	getItemEndpoint = middlewares.AuthorizationMiddleware(logger)(getItemEndpoint)
	removeItemEndpoint = middlewares.AuthorizationMiddleware(logger)(removeItemEndpoint)

	// Attach Authentication

	addLockerEndpoint = gokitjwt.NewParser(auth.Keys, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(addLockerEndpoint)
	addItemEndpoint = gokitjwt.NewParser(auth.Keys, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(addItemEndpoint)
	getItemEndpoint = gokitjwt.NewParser(auth.Keys, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(getItemEndpoint)
	removeItemEndpoint = gokitjwt.NewParser(auth.Keys, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(removeItemEndpoint)

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
		jwtOptions...,
	)
	getItemFromLockerHandler := httptransport.NewServer(
		getItemEndpoint,
		dto.DecodeHttpGetItemRequest,
		dto.EncodeHttpGetItemResponse,
		jwtOptions...,
	)
	removeFromLockerHandler := httptransport.NewServer(
		removeItemEndpoint,
		dto.DecodeHttpRemoveItemRequest,
		dto.EncodeHttpRemoveItemResponse,
		jwtOptions...,
	)

	tokenHandler := httptransport.NewServer(
		tokenEndpoint,
		dto.DecodeHttpGetTokenRequest,
		dto.EncodeHttpGetTokenResponse,
	)

	// Expose HTTP endpoints
	{
		http.Handle("/addlocker", middlewares.AccessControl(addLockerHandler))
		http.Handle("/add", middlewares.AccessControl(addItemToLockerHandler))
		http.Handle("/get", middlewares.AccessControl(getItemFromLockerHandler))
		http.Handle("/remove", middlewares.AccessControl(removeFromLockerHandler))
		http.Handle("/token", middlewares.AccessControl(tokenHandler))
	}

	// Start server
	http.ListenAndServe(":8080", nil)
}
