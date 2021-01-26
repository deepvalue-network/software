package public

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type key struct {
	ky rsa.PublicKey
}

func createKey(ky rsa.PublicKey) Key {
	out := key{
		ky: ky,
	}

	return &out
}

// Key returns the public key
func (obj *key) Key() rsa.PublicKey {
	return obj.ky
}

// Encrypt encrypts a message using the public key
func (obj *key) Encrypt(msg []byte) ([]byte, error) {
	h := sha256.New()
	return rsa.EncryptOAEP(h, rand.Reader, &obj.ky, msg, []byte(""))
}
