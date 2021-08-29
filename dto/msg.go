package dto

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
)

// Domain internal add locker request
type AddLockerRequest struct {
	Client client.Credentials
}

// Domain internal add locker response
type AddLockerResponse struct {
	LockerId uuid.UUID
	Err      error
}
