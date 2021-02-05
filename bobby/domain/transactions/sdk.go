package transactions

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/bobby/domain/structures"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewTransactionBuilder creates a new transaction builder
func NewTransactionBuilder() TransactionBuilder {
	hashAdapter := hash.NewAdapter()
	return createTransactionBuilder(hashAdapter)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithTransactions(list []Transaction) Builder
	IsAtomic() Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	All() []Transaction
	IsAtomic() bool
}

// TransactionProcessor represents a transaction processor
type TransactionProcessor interface {
	Execute(trx Transaction) ([]structures.Structure, error)
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithBody(body bodies.Body) TransactionBuilder
	WithSignature(sig signature.RingSignature) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Body() bodies.Body
	Signature() signature.RingSignature
}

// Repository represents a transaction repository
type Repository interface {
	Retrieve(chain *uuid.UUID, link hash.Hash, hash hash.Hash) (Transaction, error)
	RetrieveList(chain *uuid.UUID, link hash.Hash) (Transactions, error)
}

// Service represents a transaction service
type Service interface {
	Save(trx Transaction) error
	SaveAll(atomicTrx Transactions) error
}
