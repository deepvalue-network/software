package outgoings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an outgoings builder
type Builder interface {
	Create() Builder
	WithOutgoings(outgoings []Outgoing) Builder
	Now() (Outgoings, error)
}

// Outgoings represents an outgoings
type Outgoings interface {
	Hash() hash.Hash
	All() []Outgoing
}

// OutgoingBuilder represents an outgoing builder
type OutgoingBuilder interface {
	Create() OutgoingBuilder
	WithTransfer(transfer views.Transfer) OutgoingBuilder
	WithNote(note string) OutgoingBuilder
	CreatedOn(createdOn time.Time) OutgoingBuilder
	Now() (Outgoing, error)
}

// Outgoing represents an outgoing transfer
type Outgoing interface {
	Hash() hash.Hash
	Transfer() views.Transfer
	Note() string
	CreatedOn() time.Time
}

// Service represents an outgoing service
type Service interface {
	Insert(ins Outgoing) error
}
