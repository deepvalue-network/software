package outgoings

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents an outgoing builder
type Builder interface {
	Create() Builder
	WithPK(pk signature.PrivateKey) Builder
	WithTransfer(transfer transfers.Section) Builder
	WithNote(note string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Outgoing, error)
}

// Outgoing represents an outgoing transfer
type Outgoing interface {
	Hash() hash.Hash
	Transfer() transfers.Section
	Note() string
	CreatedOn() time.Time
}
