package authenticated

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/governments/domain/identities"
	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// NewUpdateIdentityBuilder creates a new updateIdentity builder
func NewUpdateIdentityBuilder() UpdateIdentityBuilder {
	return createUpdateIdentityBuilder()
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithPassword(pass string) Builder
	Now() (Application, error)
}

// Application represents an authenticated shareholder application
type Application interface {
	Connection() Connection
	Identity() Identity
	Proposition() Proposition
	Transaction() Transaction
	Swap() Swap
}

// Identity represents the identity application
type Identity interface {
	Retrieve() (identities.Identity, error)
	Update(update UpdateIdentity) error
	Delete() error
}

// Proposition represents an authenticated proposition application
type Proposition interface {
	New(content propositions.Content, sigs []signature.RingSignature) error
	Approve(propositionHash hash.Hash) error
	Cancel(propositionHash hash.Hash) error
	Disapprove(propositionHash hash.Hash) error
}

// Transaction represents a transaction application
type Transaction interface {
	Payment(govID *uuid.UUID, amount uint, note string) error
	Transfer(govID *uuid.UUID, amount uint, seed string, to []hash.Hash, note string) error
	View(govID *uuid.UUID, amount uint, seed string) (views.Section, error)
	ViewTransfer(section views.Section, govID *uuid.UUID, to []hash.Hash) (views.Transfer, error)
	Receive(view views.Section, pk signature.PrivateKey, note string) error
}

// Swap represents a swap application
type Swap interface {
	Request(fromGovID *uuid.UUID, amount uint, seed string, to []hash.Hash, forGov *uuid.UUID, expireOn time.Time) error
	Trade(requestHash hash.Hash, expireOn time.Time) error
	Close(tradeHash hash.Hash) error
}

// Connection represents a connection application
type Connection interface {
	Request(host string, port uint) error
	Accept(requestHash hash.Hash, name string) error
	Deny(requestHash hash.Hash) error
	Block(connID *uuid.UUID) error
	UpdateProfile(connID *uuid.UUID, profile UpdateProfile) error
}

// UpdateIdentityBuilder represents an update identity builder
type UpdateIdentityBuilder interface {
	Create() UpdateIdentityBuilder
	WithName(name string) UpdateIdentityBuilder
	WithSeed(seed string) UpdateIdentityBuilder
	WithPassword(password string) UpdateIdentityBuilder
	Now() (UpdateIdentity, error)
}

// UpdateIdentity represents an update identity
type UpdateIdentity interface {
	HasName() bool
	Name() string
	HasSeed() bool
	Seed() string
	HasPassword() bool
	Password() string
}

// UpdateProfileBuilder represents an update profile builder
type UpdateProfileBuilder interface {
	Create() UpdateProfileBuilder
	WithName(name string) UpdateProfileBuilder
	WithRank(rank uint) UpdateProfileBuilder
	Now() (UpdateProfile, error)
}

// UpdateProfile represents an update profile
type UpdateProfile interface {
	HasName() bool
	Name() string
	HasRank() bool
	Rank() uint
}
