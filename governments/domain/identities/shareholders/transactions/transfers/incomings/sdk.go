package incomings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents the incomings builder
type Builder interface {
	Create() Builder
	WithIncomings(incomings []Incoming) Builder
	Now() (Incomings, error)
}

// Incomings represents an incomings
type Incomings interface {
	Hash() hash.Hash
	All() []Incoming
}

// IncomingBuilder represents an incoming builder
type IncomingBuilder interface {
	Create() IncomingBuilder
	WithPK(pk signature.PrivateKey) IncomingBuilder
	WithTransfer(transfer views.Section) IncomingBuilder
	WithNote(note string) IncomingBuilder
	CreatedOn(createdOn time.Time) IncomingBuilder
	Now() (Incoming, error)
}

// Incoming represents an incoming transfer
type Incoming interface {
	Hash() hash.Hash
	PK() signature.PrivateKey
	Transfer() views.Section
	Note() string
	CreatedOn() time.Time
}

// Service represents an incoming service
type Service interface {
	Insert(ins Incoming) error
}
