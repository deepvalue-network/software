package transfers

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
	Now() (Transfer, error)
}

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	Transfer() views.Transfer
	Note() string
	CreatedOn() time.Time
}

// Service represents a transfer service
type Service interface {
	Insert(ins Transfer) error
}
