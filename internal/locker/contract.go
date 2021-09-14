package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/keys"
)

type LockerId = uuid.UUID // Identification of Locker

// Crypter allows to en/de crypt information using clients credentials
type Crypter interface {
	encrypt(keys.Value, PlainContent) Secret // encrypt is a function allowing to encrypt message using client identity and key
	decrypt(keys.Value, Secret) PlainContent // decrypt is a function allowing to decrypt message using client identity and key
}

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
