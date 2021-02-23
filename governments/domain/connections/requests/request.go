package requests

import (
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type request struct {
	hash    hash.Hash
	content Content
	sig     signature.Signature
}

func createRequest(
	hash hash.Hash,
	content Content,
	sig signature.Signature,
) Request {
	out := request{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *request) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *request) Content() Content {
	return obj.content
}

// Signature returns the signature
func (obj *request) Signature() signature.Signature {
	return obj.sig
}
