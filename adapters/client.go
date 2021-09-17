package adapters

import "github.com/google/uuid"

type ClientAssigner interface {
	AssignClient(uuid.UUID) interface{}
}

func (msg AddLockerRequest) AssignClient(clientId uuid.UUID) interface{} {
	msg.ClientId = clientId
	return msg
}

func (msg AddItemRequest) AssignClient(clientId uuid.UUID) interface{} {
	msg.ClientId = clientId
	return msg
}

func (msg GetItemRequest) AssignClient(clientId uuid.UUID) interface{} {
	msg.ClientId = clientId
	return msg
}

func (msg RemoveItemRequest) AssignClient(clientId uuid.UUID) interface{} {
	msg.ClientId = clientId
	return msg
}

// Client represent all data requried to perist and manage users
type Client struct {
	ClientName string
	ClientId   uuid.UUID
	Password   string
}

type ClientId struct {
	ClientId uuid.UUID
}

type ClientName struct {
	ClientName string
}
