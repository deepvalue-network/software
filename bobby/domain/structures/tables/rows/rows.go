package rows

import "github.com/deepvalue-network/software/libs/hash"

type rows struct {
	hash hash.Hash
	list []Row
}

func createRows(
	hash hash.Hash,
	list []Row,
) Rows {
	out := rows{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *rows) Hash() hash.Hash {
	return obj.hash
}

// All returns the rows
func (obj *rows) All() []Row {
	return obj.list
}

// IsEmpty returns true if the list is empty, false otherwise
func (obj *rows) IsEmpty() bool {
	return len(obj.list) <= 0
}
