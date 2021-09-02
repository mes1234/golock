package dto

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/locker"
)

// Domain internal add locker request
type AddLockerRequest struct {
	Client client.ClientId
}

// Domain internal add locker response
type AddLockerResponse struct {
	LockerId uuid.UUID
}

// Domain internal add item request
type AddItemRequest struct {
	Client   client.ClientId     // Identification of client
	LockerId locker.LockerId     // Identification of locker to insert into
	SecretId locker.SecretId     // Identification of secret to get
	Content  locker.PlainContent // Content which shall be injected
}

// Domain internal add item response
type AddItemResponse struct {
	Status bool // Operation status
}

type RemoveItemRequest struct {
	Client   client.ClientId // Identification of client
	LockerId locker.LockerId // Identification of locker to insert into
	SecretId locker.SecretId // Identification of secret to get
}

type RemoveItemResponse struct {
	Status bool // Operation status
}

type GetItemRequest struct {
	Client   client.ClientId // Identification of client
	LockerId locker.LockerId // Identification of locker to insert into
	SecretId locker.SecretId // Identification of secret to get
}

type GetItemResponse struct {
	Content locker.PlainContent // Content which shall responded
}
