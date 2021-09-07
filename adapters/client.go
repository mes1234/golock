package adapters

import "github.com/google/uuid"

type ClientAssigner interface {
	AssignClient(string) interface{}
}

func (msg AddLockerRequest) AssignClient(clientName string) interface{} {
	msg.ClientName = clientName
	return msg
}

func (msg AddItemRequest) AssignClient(clientName string) interface{} {
	msg.ClientName = clientName
	return msg
}

func (msg GetItemRequest) AssignClient(clientName string) interface{} {
	msg.ClientName = clientName
	return msg
}

func (msg RemoveItemRequest) AssignClient(clientName string) interface{} {
	msg.ClientName = clientName
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
