package trades

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/requests"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder(minPubKeys uint) ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter, minPubKeys)
}

// Builder represents a trade builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Trade, error)
}

// Trade represents a trade
type Trade interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithRequest(request requests.Request) ContentBuilder
	WithTransfer(transfer views.Transfer) ContentBuilder
	To(to []hash.Hash) ContentBuilder
	ExpiresOn(expiresOn time.Time) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represets a trade complaint
type Content interface {
	Hash() hash.Hash
	Request() requests.Request
	Transfer() views.Transfer
	To() []hash.Hash
	ExpiresOn() time.Time
	CreatedOn() time.Time
}
