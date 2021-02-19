package outgoings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an outgoing builder
type Builder interface {
	Create() Builder
	WithTransfer(transfer views.Section) Builder
	WithNote(note string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Outgoing, error)
}

// Outgoing represents an outgoing transfer
type Outgoing interface {
	Hash() hash.Hash
	Transfer() views.Section
	Note() string
	CreatedOn() time.Time
}
