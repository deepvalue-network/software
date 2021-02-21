package swaps

import "github.com/deepvalue-network/software/governments/domain/identities/transfers"

type incoming struct {
	complete Complete
	incoming transfers.Transfer
}

func createIncoming(
	complete Complete,
	in transfers.Transfer,
) Incoming {
	out := incoming{
		complete: complete,
		incoming: in,
	}

	return &out
}

// Complete returns the complete instance
func (obj *incoming) Complete() Complete {
	return obj.complete
}

// Incoming returns the incomijng transfer instance
func (obj *incoming) Incoming() transfers.Transfer {
	return obj.incoming
}
