package outgoings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an outgoing builder
type Builder interface {
	Create() Builder
	WithTransfer(transfer views.Transfer) Builder
	WithNote(note string) Builder
	CreatedOn(createdOn time.Time) Builder
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
