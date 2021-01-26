package contents

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/bobby/domain/structures/sets"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/rows"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithTable(table selectors.Selector) Builder
	WithTableRows(rows rows.Rows) Builder
	WithSet(set selectors.Selector) Builder
	WithSetElements(elements sets.Elements) Builder
	Now() (Transaction, error)
}

// Transaction represents a save transaction
type Transaction interface {
	Hash() hash.Hash
	IsTable() bool
	Table() Table
	IsSet() bool
	Set() Set
}

// Table represents a table save
type Table interface {
	Table() selectors.Selector
	Rows() rows.Rows
}

// Set represents a set save
type Set interface {
	Set() selectors.Selector
	Elements() sets.Elements
}
