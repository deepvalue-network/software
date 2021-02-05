package transactions

import "github.com/deepvalue-network/software/libs/hash"

type trans struct {
	hash         hash.Hash
	transactions []Transaction
	isAtomic     bool
}

func createTransactions(
	hash hash.Hash,
	transactions []Transaction,
) Transactions {
	return createTransactionsInternally(hash, transactions, false)
}

func createTransactionsWithAtomic(
	hash hash.Hash,
	transactions []Transaction,
) Transactions {
	return createTransactionsInternally(hash, transactions, true)
}

func createTransactionsInternally(
	hash hash.Hash,
	transactions []Transaction,
	isAtomic bool,
) Transactions {
	out := trans{
		hash:         hash,
		transactions: transactions,
		isAtomic:     isAtomic,
	}

	return &out
}

// Hash returns the hash
func (obj *trans) Hash() hash.Hash {
	return obj.hash
}

// All returns the transactions
func (obj *trans) All() []Transaction {
	return obj.transactions
}

// IsAtomic returns true if the transactions are atomic, false otherwise
func (obj *trans) IsAtomic() bool {
	return obj.isAtomic
}
