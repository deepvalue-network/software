package votes

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type vote struct {
	hash    hash.Hash
	content Content
	sig     signature.RingSignature
}

func createVote(
	hash hash.Hash,
	content Content,
	sig signature.RingSignature,
) Vote {
	out := vote{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *vote) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *vote) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *vote) Signature() signature.RingSignature {
	return obj.sig
}
