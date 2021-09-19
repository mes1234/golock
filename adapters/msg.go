package adapters

import (
	"github.com/google/uuid"
)

// Domain internal add locker request
type AddLockerRequest struct {
	ClientId uuid.UUID
}

// Domain internal add locker response
type AddLockerResponse struct {
	LockerId uuid.UUID
}

// Domain internal add item request
type AddItemRequest struct {
	ClientId uuid.UUID
	LockerId uuid.UUID // Identification of locker to insert into
	SecretId string    // Identification of secret to get
	Content  []byte    // Content which shall be injected
}

// Domain internal add item response
type AddItemResponse struct {
	Status bool // Operation status
}

type RemoveItemRequest struct {
	ClientId uuid.UUID
	LockerId uuid.UUID // Identification of locker to insert into
	SecretId string    // Identification of secret to get
}

type RemoveItemResponse struct {
	Status bool // Operation status
}

type GetItemRequest struct {
	ClientId uuid.UUID
	LockerId uuid.UUID // Identification of locker to insert into
	SecretId string    // Identification of secret to get
}

type GetItemResponse struct {
	Content []byte // Content which shall responded
}

type TokenRequest struct {
	Username string
	Password string
}

type TokenResponse struct {
	Token string
}
