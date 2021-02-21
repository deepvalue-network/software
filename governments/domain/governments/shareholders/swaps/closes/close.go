package closes

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type close struct {
	hash    hash.Hash
	content Content
	sig     signature.RingSignature
}

func createClose(
	hash hash.Hash,
	content Content,
	sig signature.RingSignature,
) Close {
	out := close{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *close) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *close) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *close) Signature() signature.RingSignature {
	return obj.sig
}
