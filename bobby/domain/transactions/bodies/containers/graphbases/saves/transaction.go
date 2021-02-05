package saves

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/tables"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash      hash.Hash
	metaData  tables.Transaction
	parent    selectors.Selector
	graphbase selectors.Selector
}

func createTransaction(
	hash hash.Hash,
	metaData tables.Transaction,
) Transaction {
	return createTransactionInternally(hash, metaData, nil, nil)
}

func createTransactionWithParent(
	hash hash.Hash,
	metaData tables.Transaction,
	parent selectors.Selector,
) Transaction {
	return createTransactionInternally(hash, metaData, parent, nil)
}

func createTransactionWithGraphbase(
	hash hash.Hash,
	metaData tables.Transaction,
	graphbase selectors.Selector,
) Transaction {
	return createTransactionInternally(hash, metaData, nil, graphbase)
}

func createTransactionWithParentAndGraphbase(
	hash hash.Hash,
	metaData tables.Transaction,
	parent selectors.Selector,
	graphbase selectors.Selector,
) Transaction {
	return createTransactionInternally(hash, metaData, parent, graphbase)
}

func createTransactionInternally(
	hash hash.Hash,
	metaData tables.Transaction,
	parent selectors.Selector,
	graphbase selectors.Selector,
) Transaction {
	out := transaction{
		hash:      hash,
		metaData:  metaData,
		parent:    parent,
		graphbase: graphbase,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// MetaData returns the metaData
func (obj *transaction) MetaData() tables.Transaction {
	return obj.metaData
}

// HasParent returns true if there is a parent, false otherwise
func (obj *transaction) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *transaction) Parent() selectors.Selector {
	return obj.parent
}

// HasGraphbase returns true if there is a graphbase, false otherwise
func (obj *transaction) HasGraphbase() bool {
	return obj.graphbase != nil
}

// Graphbase returns the graphbase, if any
func (obj *transaction) Graphbase() selectors.Selector {
	return obj.graphbase
}
