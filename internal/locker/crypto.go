package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/client"
	"github.com/mes1234/golock/internal/key"
)

// Crypter allows to en/de crypt information using clients credentials
type crypter struct{}

func NewCrypter() Crypter {
	return &crypter{}
}

// encrypt is a function allowing to encrypt message using client identity and key
func (c crypter) encrypt(clientId client.ClientId, key key.Value, plainContent PlainContent) Secret {
	return Secret{
		Id:       uuid.New(),
		Revision: 0,
		Content:  plainContent.Value,
	}
}

// decrypt is a function allowing to decrypt message using client identity and key
func (c crypter) decrypt(client.ClientId, key.Value, Secret) PlainContent {
	return PlainContent{}
}
