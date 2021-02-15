package propositions

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

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
	WithContent(content Content) ContentBuilder
	WithDeadline(deadline time.Time) ContentBuilder
	ActiveOn(activeOn time.Time) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a proposition content
type Content interface {
	Hash() hash.Hash
	Government() governments.Government
	Content() Content
	ActiveOn() time.Time
	Deadline() time.Time
	CreatedOn() time.Time
}

// SectionBuilder represents a section builder
type SectionBuilder interface {
	Create() SectionBuilder
	WithGovernment(government Government) SectionBuilder
	WithShareHolders(shareHolders ShareHolders) SectionBuilder
	WithCustom(custom hash.Hash) SectionBuilder
	Now() (Section, error)
}

// Section represents a proposition section
type Section interface {
	Hash() hash.Hash
	IsGovernment() bool
	Government() Government
	IsShareHolders() bool
	ShareHolders() ShareHolders
	IsCustom() bool
	Custom() *hash.Hash
}

// GovernmentBuilder represents a government builder
type GovernmentBuilder interface {
	Create() GovernmentBuilder
	WithID(id *uuid.UUID) GovernmentBuilder
	WithChain(chain chains.Chain) GovernmentBuilder
	WithMinPowerToPassResolution(minPowerToPassRes uint) GovernmentBuilder
	WithMinPowerToPropose(minPowerToPropose uint) GovernmentBuilder
	WithSharesVelocity(sharesVelocity uint) GovernmentBuilder
	WithSharesCap(sharesCap uint) GovernmentBuilder
	CanCancelVote() GovernmentBuilder
	BurnOnReceive() GovernmentBuilder
	CreatedOn(createdOn time.Time) GovernmentBuilder
	Now() (Government, error)
}

// Government represents a government
type Government interface {
	Hash() hash.Hash
	ID() *uuid.UUID
	HasChain() bool
	Chain() chains.Chain
	HasMinPowerToPassResolution() bool
	MinPowerToPassResolution() *uint
	HasMinPowerToPropose() bool
	MinPowerToPropose() *uint
	HasCanCancelVote() bool
	CanCancelVote() *bool
	HasBurnOnReceive() bool
	BurnOnReceive() *bool
	HasSharesVelocity() bool
	SharesVelocity() *uint
	HasSharesCap() bool
	SharesCap() *uint
}

// ShareHoldersBuilder represents a shareholders builder
type ShareHoldersBuilder interface {
	Create() ShareHoldersBuilder
	WithShareHolders(shareHolders []ShareHolder) ShareHoldersBuilder
	Now() (ShareHolders, error)
}

// ShareHolders represents shareholders
type ShareHolders interface {
	Hash() hash.Hash
	All() []ShareHolder
}

// ShareHolderBuilder represents a shareholder builder
type ShareHolderBuilder interface {
	Create() ShareHolderBuilder
	WithKeys(keys []hash.Hash) ShareHolderBuilder
	WithNewPower(newPower uint) ShareHolderBuilder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Hash() hash.Hash
	Keys() []hash.Hash
	NewPower() uint
}
