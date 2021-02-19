package payments

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type payment struct {
	hash    hash.Hash
	content Content
	sig     signature.Signature
}

func createPayment(
	hash hash.Hash,
	content Content,
	sig signature.Signature,
) Payment {
	out := payment{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *payment) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *payment) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *payment) Signature() signature.Signature {
	return obj.sig
}
