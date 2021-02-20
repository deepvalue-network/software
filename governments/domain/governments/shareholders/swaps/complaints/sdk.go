package complaints

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a complaint builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Complaint, error)
}

// Complaint represents a trade complaint
type Complaint interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithTrade(trade trades.Trade) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a complaint content
type Content interface {
	Hash() hash.Hash
	Trade() trades.Trade
	CreatedOn() time.Time
}
