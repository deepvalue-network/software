package payments

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewPaymentBuilder creates a new payment builder instance
func NewPaymentBuilder() PaymentBuilder {
	hashAdapter := hash.NewAdapter()
	return createPaymentBuilder(hashAdapter)
}

// Builder creates a new payments instance
type Builder interface {
	Create() Builder
	WithPayments(payments []Payment) Builder
	Now() (Payments, error)
}

// Payments retruns the payments
type Payments interface {
	Hash() hash.Hash
	All() []Payment
}

// PaymentBuilder represents a payment builder
type PaymentBuilder interface {
	Create() PaymentBuilder
	WithPayment(payment payments.Payment) PaymentBuilder
	WithNote(note string) PaymentBuilder
	Now() (Payment, error)
}

// Payment represents a payment
type Payment interface {
	Hash() hash.Hash
	Payment() payments.Payment
	Note() string
}
