package service

import (
	"context"

	"github.com/mes1234/golock/adapters"
)

//TokenService interface describes a service that can get access tokens
type TokenService interface {

	// Add item to locker
	GetToken(
		ctx context.Context,
		request adapters.TokenRequest,
	) (adapters.TokenResponse, error)
}
