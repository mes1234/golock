package client

import "github.com/google/uuid"

//  Identity is a unique identity which identifes user
type Identity struct {
	Id ClientId
}

// Client password
type Password struct {
	Value string
}

type ClientId = uuid.UUID // Identification of client

// All data needed to identify user
type Credentials struct {
	Identity Identity
	Password Password
}
