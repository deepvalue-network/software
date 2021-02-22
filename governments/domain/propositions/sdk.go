package propositions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// NewSectionBuilder creates a new section builder instance
func NewSectionBuilder() SectionBuilder {
	return createSectionBuilder()
}

// Builder represents a proposition builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignatures(sigs []signature.RingSignature) Builder
	Now() (Proposition, error)
}

// Proposition represents a proposition
type Proposition interface {
	Hash() hash.Hash
	Content() Content
	Signatures() []signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithGovernment(gov governments.Government) ContentBuilder
	WithSection(section Section) ContentBuilder
	WithDeadline(deadline time.Time) ContentBuilder
	ActiveOn(activeOn time.Time) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a proposition content
type Content interface {
	Hash() hash.Hash
	Government() governments.Government
	Section() Section
	ActiveOn() time.Time
	Deadline() time.Time
	CreatedOn() time.Time
}

// SectionBuilder represents a section builder
type SectionBuilder interface {
	Create() SectionBuilder
	WithGovernment(government governments.Content) SectionBuilder
	WithShareHolders(shareHolders shareholders.ShareHolders) SectionBuilder
	WithCustom(custom hash.Hash) SectionBuilder
	Now() (Section, error)
}

// Section represents a proposition section
type Section interface {
	Hash() hash.Hash
	IsGovernment() bool
	Government() governments.Content
	IsShareHolders() bool
	ShareHolders() shareholders.ShareHolders
	IsCustom() bool
	Custom() *hash.Hash
}

// Repository represents a proposition repository
type Repository interface {
	Retrieve(hash hash.Hash) (Proposition, error)
}

// Service represents a proposition service
type Service interface {
	Insert(proposition Proposition) error
}
