package locker

import (
	"crypto/aes"
	"crypto/cipher"
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

	cp, _ := aes.NewCipher(prepareKey(key))
	gcm, _ := cipher.NewGCM(cp)
	nonce := make([]byte, gcm.NonceSize())

	encrypted := gcm.Seal(nonce, nonce, plainContent, nil)
	return Secret{
		Id:      uuid.New(),
		Content: encrypted,
	}
}

func prepareKey(key keys.Value) []byte {
	keyBytes := []byte(key.Key)
	keyLen := len(keyBytes)
	filler := make([]byte, 32-keyLen)
	longKey := append(keyBytes, filler...)
	return longKey
}

// decrypt is a function allowing to decrypt message using client identity and key
func (c crypter) decrypt(key keys.Value, s Secret) []byte {

	cp, _ := aes.NewCipher(prepareKey(key))
	gcm, _ := cipher.NewGCM(cp)
	nonce := make([]byte, gcm.NonceSize())
	nonce, ciphertext := s.Content[:gcm.NonceSize()], s.Content[gcm.NonceSize():]

	decrypted, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return decrypted
}
