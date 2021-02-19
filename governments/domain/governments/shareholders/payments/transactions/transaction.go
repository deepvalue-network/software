package transactions

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash    hash.Hash
	payment payments.Payment
	note    string
}

func createTransaction(
	hash hash.Hash,
	payment payments.Payment,
	note string,
) Transaction {
	out := transaction{
		hash:    hash,
		payment: payment,
		note:    note,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Payment returns the payment
func (obj *transaction) Payment() payments.Payment {
	return obj.payment
}

// Note returns the note
func (obj *transaction) Note() string {
	return obj.note
}
