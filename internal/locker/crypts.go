package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/keys"
)

// Crypter allows to en/de crypt information using clients credentials
type Crypter interface {
	encrypt(keys.Value, PlainContent) Secret // encrypt is a function allowing to encrypt message using client identity and key
	decrypt(keys.Value, Secret) PlainContent // decrypt is a function allowing to decrypt message using client identity and key
}

// Crypter allows to en/de crypt information using clients credentials
type crypter struct{}

func NewCrypter() Crypter {
	return &crypter{}
}

// encrypt is a function allowing to encrypt message using client identity and key
func (c crypter) encrypt(key keys.Value, plainContent PlainContent) Secret {
	return Secret{
		Id:      uuid.New(),
		Content: plainContent.Value,
	}
}

// decrypt is a function allowing to decrypt message using client identity and key
func (c crypter) decrypt(keys.Value, Secret) PlainContent {
	return PlainContent{}
}
