package persistance

import (
	"github.com/google/uuid"
)

type SecretPersisted struct {
	Active     bool
	Revision   int
	Content    []byte
	SecretName string
	LockerId   uuid.UUID
	ClientId   uuid.UUID
}
