package transfers

import (
	"github.com/deepvalue-network/software/governments/domain/shareholders/transactions/transfers/incomings"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transactions/transfers/outgoings"
)

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithIncomings() []incomings.Incoming
	WithOutgoings() []outgoings.Outgoing
	Now() (Transfers, error)
}

// Transfers represents the transfers
type Transfers interface {
	Incomings() []incomings.Incoming
	Outgoings() []outgoings.Outgoing
}
