package bills

import (
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

type bill struct {
	hash    hash.Hash
	content Content
	sig     signature.RingSignature
}

func createBill(
	hash hash.Hash,
	content Content,
	sig signature.RingSignature,
) Bill {
	out := bill{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *bill) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *bill) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *bill) Signature() signature.RingSignature {
	return obj.sig
}
