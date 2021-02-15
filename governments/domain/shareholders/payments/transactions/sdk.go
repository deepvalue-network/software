package transactions

import (
	"github.com/deepvalue-network/software/governments/domain/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithPayment(payment payments.Payment) Builder
	WithNote(note string) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Payment() payments.Payment
	Note() string
}
