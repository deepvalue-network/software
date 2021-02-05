package databases

import (
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/databases/deletes"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/databases/saves"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	del  deletes.Transaction
	save saves.Transaction
}

func createTransactionWithDelete(
	del deletes.Transaction,
) Transaction {
	return createTransactionInternally(del, nil)
}

func createTransactionWithSave(
	save saves.Transaction,
) Transaction {
	return createTransactionInternally(nil, save)
}

func createTransactionInternally(
	del deletes.Transaction,
	save saves.Transaction,
) Transaction {
	out := transaction{
		del:  del,
		save: save,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	if obj.IsDelete() {
		return obj.Delete().Hash()
	}

	return obj.Save().Hash()
}

// IsDelete returns true if there is a delete transaction, false otherwise
func (obj *transaction) IsDelete() bool {
	return obj.del != nil
}

// Delete returns the delete transaction, if any
func (obj *transaction) Delete() deletes.Transaction {
	return obj.del
}

// IsSave returns true if there is a save transaction, false otherwise
func (obj *transaction) IsSave() bool {
	return obj.save != nil
}

// Save returns the save transaction, if any
func (obj *transaction) Save() saves.Transaction {
	return obj.save
}
