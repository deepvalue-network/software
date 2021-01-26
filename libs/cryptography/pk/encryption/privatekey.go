package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/steve-care-software/products/libs/cryptography/pk/encryption/public"
)

type privateKey struct {
	key    rsa.PrivateKey
	pubKey public.Key
}

func createPrivateKey(key rsa.PrivateKey, pubKey public.Key) PrivateKey {
	out := privateKey{
		key:    key,
		pubKey: pubKey,
	}

	return &out
}

// Key returns the key
func (obj *privateKey) Key() rsa.PrivateKey {
	return obj.key
}

// Public returns the public key
func (obj *privateKey) Public() public.Key {
	return obj.pubKey
}

// Decrypt decrypts a cipher
func (obj *privateKey) Decrypt(cipher []byte) ([]byte, error) {
	h := sha256.New()
	decrypted, err := rsa.DecryptOAEP(h, rand.Reader, &obj.key, cipher, []byte(""))
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
