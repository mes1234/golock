package adapters

import "github.com/google/uuid"

type ClientAssigner interface {
	AssignClientId(uuid.UUID) interface{}
}

func (msg AddLockerRequest) AssignClientId(id uuid.UUID) interface{} {
	msg.ClientId = id
	return msg
}

// Client represent all data requried to perist and manage users
type Client struct {
	ClientName string
	ClientId   uuid.UUID
	Password   string
}
