package payments

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Builder represents a payment builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.Signature) Builder
	Now() (Payment, error)
}

// Payment represents a payment from a shareholder to its government
type Payment interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.Signature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithShareHolder(shareHolder shareholders.ShareHolder) ContentBuilder
	WithAmount(amount uint) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a payment content
type Content interface {
	Hash() hash.Hash
	ShareHolder() shareholders.ShareHolder
	Amount() uint
	CreatedOn() time.Time
}
