package payments

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

type payment struct {
	hash    hash.Hash
	payment payments.Payment
	note    string
}

func createPayment(
	hash hash.Hash,
	pay payments.Payment,
	note string,
) Payment {
	out := payment{
		hash:    hash,
		payment: pay,
		note:    note,
	}

	return &out
}

// Hash returns the hash
func (obj *payment) Hash() hash.Hash {
	return obj.hash
}

// Payment returns the payment
func (obj *payment) Payment() payments.Payment {
	return obj.payment
}

// Note returns the note
func (obj *payment) Note() string {
	return obj.note
}
