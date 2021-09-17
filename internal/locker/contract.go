package locker

import (
	"github.com/google/uuid"
)

type LockerId = uuid.UUID // Identification of Locker

// PlainContent is contend which client requested to be encrypted
type PlainContent struct {
	Value []byte
}

type SecretId string // identifier of Secret
// Secret is single secret instance
type Secret struct {
	Id      uuid.UUID // Identifier of secret
	Content []byte    //encrypted content of secret
}
