package owners

import (
	"time"

	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

type owner struct {
	hash      hash.Hash
	seed      string
	sigPK     signature.PrivateKey
	encPK     encryption.PrivateKey
	createdOn time.Time
}

func createOwner(
	hash hash.Hash,
	seed string,
	sigPK signature.PrivateKey,
	encPK encryption.PrivateKey,
	createdOn time.Time,
) Owner {
	out := owner{
		hash:      hash,
		seed:      seed,
		sigPK:     sigPK,
		encPK:     encPK,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *owner) Hash() hash.Hash {
	return obj.hash
}

// Seed returns the seed
func (obj *owner) Seed() string {
	return obj.seed
}

// Signature returns the signature PK
func (obj *owner) Signature() signature.PrivateKey {
	return obj.sigPK
}

// Encryption returns the encryption PK
func (obj *owner) Encryption() encryption.PrivateKey {
	return obj.encPK
}

// CreatedOn returns the creation time
func (obj *owner) CreatedOn() time.Time {
	return obj.createdOn
}
