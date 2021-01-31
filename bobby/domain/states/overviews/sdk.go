package overviews

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	domain_errors "github.com/steve-care-software/products/bobby/domain/errors"
	"github.com/steve-care-software/products/bobby/domain/structures"
	"github.com/steve-care-software/products/bobby/domain/transactions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewInvalidTransactionBuilder creates a new invalid transation builder
func NewInvalidTransactionBuilder() InvalidTransactionBuilder {
	return createInvalidTransactionBuilder()
}

// NewValidTransactionBuilder creates a new valid transaction builder
func NewValidTransactionBuilder() ValidTransactionBuilder {
	return createValidTransactionBuilder()
}

// Builder represents the overview builder
type Builder interface {
	Create() Builder
	WithValid(valid []ValidTransaction) Builder
	WithInvalid(invalid []InvalidTransaction) Builder
	CanBeSaved() Builder
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
	WithError(err domain_errors.Error) InvalidTransactionBuilder
	Now() (InvalidTransaction, error)
}

// InvalidTransaction represents an invalid transaction
type InvalidTransaction interface {
	Transaction() transactions.Transaction
	Error() domain_errors.Error
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
