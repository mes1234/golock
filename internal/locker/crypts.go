package locker

import (
	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/keys"
)

// Crypter allows to en/de crypt information using clients credentials
type Crypter interface {
	encrypt(keys.Value, []byte) Secret // encrypt is a function allowing to encrypt message using client identity and key
	decrypt(keys.Value, Secret) []byte // decrypt is a function allowing to decrypt message using client identity and key
}

// Crypter allows to en/de crypt information using clients credentials
type crypter struct{}

func NewCrypter() Crypter {
	return &crypter{}
}

// encrypt is a function allowing to encrypt message using client identity and key
func (c crypter) encrypt(key keys.Value, plainContent []byte) Secret {
	return Secret{
		Id:      uuid.New(),
		Content: plainContent,
	}
}

// decrypt is a function allowing to decrypt message using client identity and key
func (c crypter) decrypt(keys.Value, Secret) []byte {
	return make([]byte, 0)
}
