package schemas

import (
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/tables"
)

// NewBuilder creates a new schema builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the schema builder
type Builder interface {
	Create() Builder
	WithResource(res resources.Accessible) Builder
	WithName(name string) Builder
	WithTable(table tables.Table) Builder
	IsUniqueElements() Builder
	Now() (Schema, error)
}

// Schema represents a set schema
type Schema interface {
	Resource() resources.Accessible
	Name() string
	Table() tables.Table
	IsUniqueElements() bool
}
