package contents

import "github.com/deepvalue-network/software/libs/hash"

type transaction struct {
	hash  hash.Hash
	table Table
	set   Set
}

func createTransactionWithTable(
	hash hash.Hash,
	table Table,
) Transaction {
	return createTableInternally(hash, table, nil)
}

func createTransactionWithSet(
	hash hash.Hash,
	set Set,
) Transaction {
	return createTableInternally(hash, nil, set)
}

func createTableInternally(
	hash hash.Hash,
	table Table,
	set Set,
) Transaction {
	out := transaction{
		hash:  hash,
		table: table,
		set:   set,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// IsTable returns true if there is a table, false otherwise
func (obj *transaction) IsTable() bool {
	return obj.table != nil
}

// Table returns the table, if any
func (obj *transaction) Table() Table {
	return obj.table
}

// IsSet returns true if there is a set, false otherwise
func (obj *transaction) IsSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *transaction) Set() Set {
	return obj.set
}
