package propositions

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type proposition struct {
	hash    hash.Hash
	content Content
	sigs    []signature.RingSignature
}

func createProposition(
	hash hash.Hash,
	content Content,
	sigs []signature.RingSignature,
) Proposition {
	out := proposition{
		hash:    hash,
		content: content,
		sigs:    sigs,
	}

	return &out
}

// Hash returns the hash
func (obj *proposition) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *proposition) Content() Content {
	return obj.content
}

// Signatures returns the signatures
func (obj *proposition) Signatures() []signature.RingSignature {
	return obj.sigs
}
