package views

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type transfer struct {
	hash    hash.Hash
	content Content
	sig     signature.RingSignature
}

func createTransfer(
	hash hash.Hash,
	content Content,
	sig signature.RingSignature,
) Transfer {
	out := transfer{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *transfer) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *transfer) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *transfer) Signature() signature.RingSignature {
	return obj.sig
}
