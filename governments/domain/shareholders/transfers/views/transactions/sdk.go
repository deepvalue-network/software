package transactions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// OutgoingBuilder represents an outgoing transaction builder
type OutgoingBuilder interface {
	Create() OutgoingBuilder
	WithPK(pk signature.PrivateKey) OutgoingBuilder
	WithTransfer(transfer views.Section) OutgoingBuilder
	WithNote(note string) OutgoingBuilder
	CreatedOn(createdOn time.Time) OutgoingBuilder
	Now() (Outgoing, error)
}

// Outgoing represents an outgoing transaction
type Outgoing interface {
	Hash() hash.Hash
	Transfer() views.Section
	Note() string
	CreatedOn() time.Time
}

// IncomingBuilder represents an incoming transaction builder
type IncomingBuilder interface {
	Create() IncomingBuilder
	WithPK(pk signature.PrivateKey) IncomingBuilder
	WithTransfer(transfer views.Section) IncomingBuilder
	WithNote(note string) IncomingBuilder
	CreatedOn(createdOn time.Time) IncomingBuilder
	Now() (Incoming, error)
}

// Incoming represents a incoming transaction
type Incoming interface {
	Hash() hash.Hash
	PK() signature.PrivateKey
	Transfer() views.Section
	Note() string
	CreatedOn() time.Time
}
