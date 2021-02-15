package views

import (
	"github.com/deepvalue-network/software/governments/domain/shareholders"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a transfer builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Transfer, error)
}

// Transfer represents a view transfer
type Transfer interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithSection(section Section) ContentBuilder
	WithShareHolder(shareHolder []hash.Hash) ContentBuilder
	Now() (Content, error)
}

// Content represents a view transfer content
type Content interface {
	Hash() hash.Hash
	Section() Section
	ShareHolder() []hash.Hash
}

// SectionBuilder represents a section builder
type SectionBuilder interface {
	Create() SectionBuilder
	WithShareHolder(shareHolder shareholders.ShareHolder) SectionBuilder
	WithTransfer(transfer transfers.Transfer) SectionBuilder
	WithSeed(seed string) SectionBuilder
	WithAmount(amount uint) SectionBuilder
	Now() (Section, error)
}

// Section represents a view transfer section
type Section interface {
	Hash() hash.Hash
	ShareHolder() shareholders.ShareHolder
	Transfer() transfers.Transfer
	Seed() string
	Amount() uint
}
