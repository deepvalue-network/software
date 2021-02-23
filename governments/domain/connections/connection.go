package connections

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type connection struct {
	hash    hash.Hash
	content Content
	sig     signature.Signature
}

func createConnection(
	hash hash.Hash,
	content Content,
	sig signature.Signature,
) Connection {
	out := connection{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *connection) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *connection) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *connection) Signature() signature.Signature {
	return obj.sig
}
