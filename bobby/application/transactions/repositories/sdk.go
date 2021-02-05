package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/bobby/domain/transactions"
	"github.com/deepvalue-network/software/libs/hash"
)

// Application represents the transaction repository
type Application interface {
	Retrieve(chain *uuid.UUID, link hash.Hash, hash hash.Hash) (transactions.Transaction, error)
	RetrieveList(chain *uuid.UUID, link hash.Hash) (transactions.Transactions, error)
}
