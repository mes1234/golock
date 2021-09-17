package adapters

import (
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/locker"
)

// Domain internal add locker request
type AddLockerRequest struct {
	ClientId client.ClientId
}

// Domain internal add locker response
type AddLockerResponse struct {
	LockerId locker.LockerId
}

// Domain internal add item request
type AddItemRequest struct {
	ClientId client.ClientId
	LockerId locker.LockerId     // Identification of locker to insert into
	SecretId locker.SecretId     // Identification of secret to get
	Content  locker.PlainContent // Content which shall be injected
}

// Domain internal add item response
type AddItemResponse struct {
	Status bool // Operation status
}

type RemoveItemRequest struct {
	ClientId client.ClientId
	LockerId locker.LockerId // Identification of locker to insert into
	SecretId locker.SecretId // Identification of secret to get
}

type RemoveItemResponse struct {
	Status bool // Operation status
}

type GetItemRequest struct {
	ClientId client.ClientId
	LockerId locker.LockerId // Identification of locker to insert into
	SecretId locker.SecretId // Identification of secret to get
}

type GetItemResponse struct {
	Content locker.PlainContent // Content which shall responded
}

type TokenRequest struct {
	Username string
	Password string
}

type TokenResponse struct {
	Token string
}
