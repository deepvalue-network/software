package payments

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new payment builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a payment builder
type Builder interface {
	Create() Builder
	WithPayment(payment payments.Payment) Builder
	WithNote(note string) Builder
	Now() (Payment, error)
}

// Payment represents a payment
type Payment interface {
	Hash() hash.Hash
	Payment() payments.Payment
	Note() string
}

// Service represents a payment service
type Service interface {
	Insert(payment Payment) error
}
