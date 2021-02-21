package transfers

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type transfer struct {
	hash     hash.Hash
	transfer views.Transfer
	note     string
}

func createTransfer(
	hash hash.Hash,
	trsf views.Transfer,
	note string,
) Transfer {
	out := transfer{
		hash:     hash,
		transfer: trsf,
		note:     note,
	}

	return &out
}

// Hash returns the hash
func (obj *transfer) Hash() hash.Hash {
	return obj.hash
}

// Transfer returns the transfer
func (obj *transfer) Transfer() views.Transfer {
	return obj.transfer
}

// Note returns the note
func (obj *transfer) Note() string {
	return obj.note
}
