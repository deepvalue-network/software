package defenses

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/complaints"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a defense builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Defense, error)
}

// Defense represents a complaint defense
type Defense interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithComplaint(complaint complaints.Complaint) ContentBuilder
	WithTransfer(transfer transfers.Transfer) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a defense content
type Content interface {
	Hash() hash.Hash
	Complaint() complaints.Complaint
	Transfer() views.Transfer
	CreatedOn() time.Time
}
