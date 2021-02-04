package owners

import (
	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Signature() signature.PrivateKey
	Encryption() encryption.PrivateKey
}
