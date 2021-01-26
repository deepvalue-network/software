package overviews

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/bobby/domain/errors"
	"github.com/steve-care-software/products/bobby/domain/structures"
	"github.com/steve-care-software/products/bobby/domain/transactions"
)

// OverviewBuilder represents the overview builder
type OverviewBuilder interface {
	Create() OverviewBuilder
	WithValid(valid []ValidTransaction) OverviewBuilder
	WithInvalid(invalid []InvalidTransaction) OverviewBuilder
	CanBeSaved() OverviewBuilder
	Now() (Overview, error)
}

// Overview represents an overview instance
type Overview interface {
	Valid() []ValidTransaction
	Invalid() []InvalidTransaction
	CanBeSaved() bool
}

// InvalidTransactionBuilder represents an invalid transaction builder
type InvalidTransactionBuilder interface {
	Create() InvalidTransactionBuilder
	WithTransaction(trx transactions.Transaction) InvalidTransactionBuilder
	WithError(err errors.Error) InvalidTransactionBuilder
	Now() (InvalidTransaction, error)
}

// InvalidTransaction represents an invalid transaction
type InvalidTransaction interface {
	Transaction() transactions.Transaction
	Error() errors.Error
}

// ValidTransactionBuilder represents a valid transaction builder
type ValidTransactionBuilder interface {
	Create() ValidTransactionBuilder
	WithTransaction(trx transactions.Transaction) ValidTransactionBuilder
	WithStructures(structures []structures.Structure) ValidTransactionBuilder
	WithChain(chain chains.Chain) ValidTransactionBuilder
	WithBlock(block blocks.Block) ValidTransactionBuilder
	Now() (ValidTransaction, error)
}

// ValidTransaction represents a valid transaction
type ValidTransaction interface {
	Transaction() transactions.Transaction
	Structures() []structures.Structure
	Chain() chains.Chain
	Block() blocks.Block
}
