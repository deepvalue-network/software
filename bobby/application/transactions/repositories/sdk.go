package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/bobby/domain/transactions"
	"github.com/steve-care-software/products/libs/hash"
)

// Application represents the transaction repository
type Application interface {
	Retrieve(chain *uuid.UUID, link hash.Hash, hash hash.Hash) (transactions.Transaction, error)
	RetrieveList(chain *uuid.UUID, link hash.Hash) (transactions.Transactions, error)
}
