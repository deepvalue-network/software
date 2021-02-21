package requests

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
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

// Builder represents a request builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Request, error)
}

// Request represents a swap request
type Request interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	From(from governments.Government) ContentBuilder
	WithStake(stake views.Section) ContentBuilder
	For(forGov governments.Government) ContentBuilder
	To(to []hash.Hash) ContentBuilder
	WithAmount(amount uint) ContentBuilder
	ExpiresOn(expiresOn time.Time) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a request content
type Content interface {
	Hash() hash.Hash
	From() governments.Government
	Stake() views.Section
	For() governments.Government
	To() []hash.Hash
	Amount() uint
	ExpiresOn() time.Time
	CreatedOn() time.Time
}
