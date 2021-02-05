package selectors

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/hash"
)

type selector struct {
	hash          hash.Hash
	decryptionKey encryption.PrivateKey
	content       Content
}

func createSelector(
	hash hash.Hash,
	decryptionKey encryption.PrivateKey,
	content Content,
) Selector {
	out := selector{
		hash:          hash,
		decryptionKey: decryptionKey,
		content:       content,
	}

	return &out
}

// Hash returns the hash
func (obj *selector) Hash() hash.Hash {
	return obj.hash
}

// DecryptionKey returns the decryption key
func (obj *selector) DecryptionKey() encryption.PrivateKey {
	return obj.decryptionKey
}

// Content returns the content
func (obj *selector) Content() Content {
	return obj.content
}
