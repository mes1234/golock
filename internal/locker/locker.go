package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/key"
)

type LockerId = uuid.UUID // Identification of Locker

// Locker is container for all secrect
type Locker struct {
	Id      LockerId                            // Identifier of locker
	Clients map[client.ClientId]client.Identity // Identifiers of all clients with access
	Secrets map[SecretId]Secret                 //Content of Locker

}

// Crypter allows to en/de crypt information using clients credentials
type Crypter interface {
	encrypt(client.Identity, key.Value, PlainContent)        // encrypt is a function allowing to encrypt message using client identity and key
	decrypt(client.Identity, key.Value, Secret) PlainContent // decrypt is a function allowing to decrypt message using client identity and key
}

// PlainContent is contend which client requested to be encrypted
type PlainContent = []byte

type SecretId string // identifier of Secret
// Secret is single secret instance
type Secret struct {
	Id       uuid.UUID // Identifier of secret
	Revision int16     // Revision for version control
	Content  []byte    //encrypted content of secret
}
