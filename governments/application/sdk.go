package application

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/resolutions"
	"github.com/deepvalue-network/software/governments/domain/shareholders"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers/views"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers/views/transactions"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// Application represents the government application
type Application interface {
	Retrieve(id *uuid.UUID) (governments.Government, error)
	Proposition(governmentID *uuid.UUID) Proposition
	Resolution(governmentID *uuid.UUID) Resolution
	ShareHolder(governmentID *uuid.UUID) ShareHolder
}

// Proposition represents the proposition application
type Proposition interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) propositions.Proposition
}

// Resolution represents a resolution
type Resolution interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (resolutions.Resolution, error)
}

// ShareHolder represents the shareholder application
type ShareHolder interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (shareholders.ShareHolder, error)
	Authenticate(hash hash.Hash, pk signature.PrivateKey) AuthenticatedShareHolder
}

// AuthenticatedShareHolder represents an authenticated shareholder
type AuthenticatedShareHolder interface {
	Propose() AuthenticatedProposition
	Transfer(amount uint, seed string, to []hash.Hash, note string) error
	View(amount uint, seed string, to []hash.Hash) (views.Section, error)
	Receive(view views.Section, pk signature.PrivateKey, note string) error
	Transaction(filter DateFilter) Transaction
}

// AuthenticatedProposition represents an authenticated proposition
type AuthenticatedProposition interface {
	New(content propositions.Proposition, sigs []signature.RingSignature) error
	Approve(propositionHash hash.Hash) error
	Cancel(propositionHash hash.Hash) error
	Disapprove(propositionHash hash.Hash) error
}

// Transaction represents a transaction application
type Transaction interface {
	Incoming() IncomingTransaction
	Outgoing() OutgoingTransaction
}

// IncomingTransaction represents an incoming transaction application
type IncomingTransaction interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) ([]transactions.Incoming, error)
}

// OutgoingTransaction represents an outgoing transaction application
type OutgoingTransaction interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) ([]transactions.Outgoing, error)
}

// DateFilterBuilder represents a date filter builder
type DateFilterBuilder interface {
	Create() DateFilterBuilder
	From(from time.Time) DateFilterBuilder
	To(to time.Time) DateFilterBuilder
	Now() (DateFilter, error)
}

// DateFilter represents a date filter
type DateFilter interface {
	HasFrom() bool
	From() *time.Time
	HasTo() bool
	To() *time.Time
}
