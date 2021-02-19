package incomings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an incoming builder
type Builder interface {
	Create() Builder
	WithPK(pk signature.PrivateKey) Builder
	WithTransfer(transfer transfers.Section) Builder
	WithNote(note string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Incoming, error)
}

// Incoming represents an incoming transfer
type Incoming interface {
	Hash() hash.Hash
	PK() signature.PrivateKey
	Transfer() transfers.Section
	Note() string
	CreatedOn() time.Time
}
