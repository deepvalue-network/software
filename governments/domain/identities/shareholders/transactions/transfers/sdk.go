package transfers

import (
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/transfers/incomings"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/transfers/outgoings"
)

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithIncomings() incomings.Incomings
	WithOutgoings() outgoings.Outgoings
	Now() (Transfers, error)
}

// Transfers represents the transfers
type Transfers interface {
	Incomings() incomings.Incomings
	Outgoings() outgoings.Outgoings
}
