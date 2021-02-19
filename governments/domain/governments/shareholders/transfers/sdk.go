package transfers

import (
	"time"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewContentBuilder creates a new content builder instance
func NewContentBuilder(minPubKeysPerRingSig uint) ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter, minPubKeysPerRingSig)
}

// Builder represents a transfer builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Transfer, error)
}

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithOrigin(origin hash.Hash) ContentBuilder
	WithAmount(amount hash.Hash) ContentBuilder
	WithSeed(seed hash.Hash) ContentBuilder
	WithOwner(owner []hash.Hash) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a transfer content
type Content interface {
	Hash() hash.Hash
	Origin() hash.Hash
	Amount() hash.Hash
	Seed() hash.Hash
	Owner() []hash.Hash
	CreatedOn() time.Time
}
