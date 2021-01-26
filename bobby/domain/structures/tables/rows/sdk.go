package rows

import (
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/tables"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/elements"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewRowBuilder creates a new row builder instance
func NewRowBuilder() RowBuilder {
	hashAdapter := hash.NewAdapter()
	resourceBuilder := resources.NewBuilder()
	elementsBuilder := elements.NewBuilder()
	return createRowBuilder(hashAdapter, resourceBuilder, elementsBuilder)
}

// Builder represents the rows builder
type Builder interface {
	Create() Builder
	WithRows(rows []Row) Builder
	Now() (Rows, error)
}

// Rows represents rows
type Rows interface {
	Hash() hash.Hash
	All() []Row
	IsEmpty() bool
}

// RowBuilder represents the row builder
type RowBuilder interface {
	Create() RowBuilder
	WithElements(elements []elements.Element) RowBuilder
	OnTable(table tables.Table) RowBuilder
	Now() (Row, error)
}

// Row represents a table row
type Row interface {
	Resource() resources.Resource
	Elements() elements.Elements
	OnTable() tables.Table
}
