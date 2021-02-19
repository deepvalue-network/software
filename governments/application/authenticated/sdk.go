package authenticated

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/governments/domain/identities"
	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

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
	Identity() Identity
	Proposition() Proposition
	Transactions() Transactions
}

// Identity represents the identity application
type Identity interface {
	Retrieve() (identities.Identity, error)
	Update(update UpdateIdentity, password string) error
	Delete() error
}

// Proposition represents an authenticated proposition application
type Proposition interface {
	New(content propositions.Proposition, sigs []signature.RingSignature) error
	Approve(propositionHash hash.Hash) error
	Cancel(propositionHash hash.Hash) error
	Disapprove(propositionHash hash.Hash) error
}

// Transactions represents a transactions application
type Transactions interface {
	Payment(amount uint, note string) error
	Transfer(amount uint, seed string, to []hash.Hash, note string) error
	View(amount uint, seed string, to []hash.Hash) (views.Section, error)
	Receive(view views.Section, pk signature.PrivateKey, note string) error
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
