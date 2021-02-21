package trades

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type trade struct {
	hash    hash.Hash
	content Content
	sig     signature.RingSignature
}

func createTrade(
	hash hash.Hash,
	content Content,
	sig signature.RingSignature,
) Trade {
	out := trade{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *trade) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *trade) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *trade) Signature() signature.RingSignature {
	return obj.sig
}
