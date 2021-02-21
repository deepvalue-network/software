package incomings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an incoming builder
type Builder interface {
	Create() Builder
	WithPK(pk signature.PrivateKey) Builder
	WithTransfer(transfer views.Section) Builder
	WithNote(note string) Builder
	CreatedOn(createdOn time.Time) Builder
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
