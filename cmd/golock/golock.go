package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mes1234/golock/dto"
	"github.com/mes1234/golock/endpoints"
	"github.com/mes1234/golock/service"
)

func main() {
	svc := service.AccessService{}

	lockerHandler := httptransport.NewServer(
		endpoints.MakeAddLockerEndpoint(svc),
		dto.DecodeHttpAddLockerRequest,
		dto.EncodeHttpAddLockerResponse,
	)

	http.Handle("/addlocker", lockerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
