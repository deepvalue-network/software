package selectors

import (
	"github.com/steve-care-software/products/bobby/domain/selectors/specifiers"
	"github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
	table_schemas "github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	pkAdapter := encryption.NewAdapter()
	return createBuilder(
		hashAdapter,
		pkAdapter,
	)
}

// NewGraphbaseBuilder creates a new graphbase builder instance
func NewGraphbaseBuilder() GraphbaseBuilder {
	hashAdapter := hash.NewAdapter()
	return createGraphbaseBuilder(hashAdapter)
}

// NewDatabaseBuilder creates a new database builder instance
func NewDatabaseBuilder() DatabaseBuilder {
	hashAdapter := hash.NewAdapter()
	return createDatabaseBuilder(hashAdapter)
}

// NewSetBuilder creates a new set builder instance
func NewSetBuilder() SetBuilder {
	hashAdapter := hash.NewAdapter()
	return createSetBuilder(hashAdapter)
}

// NewTableBuilder creates anew table builder instance
func NewTableBuilder() TableBuilder {
	hashAdapter := hash.NewAdapter()
	return createTableBuilder(hashAdapter)
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithDecryptionKey(decryptionKey encryption.PrivateKey) Builder
	WithGraphbase(graphbase Graphbase) Builder
	WithDatabase(db Database) Builder
	WithTable(table Table) Builder
	WithSet(set Set) Builder
	Now() (Selector, error)
}

// Selector retrieves data from the database
type Selector interface {
	Hash() hash.Hash
	DecryptionKey() encryption.PrivateKey
	Content() Content
}

// Content represents a selector content
type Content interface {
	IsGraphbase() bool
	Graphbase() Graphbase
	IsDatabase() bool
	Database() Database
	IsTable() bool
	Table() Table
	IsSet() bool
	Set() Set
}

// GraphbaseBuilder represents a graphbase builder
type GraphbaseBuilder interface {
	Create() GraphbaseBuilder
	WithParent(parent specifiers.Specifier) GraphbaseBuilder
	WithSpecifier(specifier specifiers.Specifier) GraphbaseBuilder
	WithMetaData(metaData Table) GraphbaseBuilder
	Now() (Graphbase, error)
}

// Graphbase represents a graphbase
type Graphbase interface {
	Hash() hash.Hash
	Content() GraphbaseContent
	HasParent() bool
	Parent() specifiers.Specifier
}

// GraphbaseContent represents a graphbase content
type GraphbaseContent interface {
	IsSpecifier() bool
	Specifier() specifiers.Specifier
	IsMetaData() bool
	MetaData() Table
}

// DatabaseBuilder represents a database builder
type DatabaseBuilder interface {
	Create() DatabaseBuilder
	WithGraphbase(graphbase specifiers.Specifier) DatabaseBuilder
	WithSpecifier(specifier specifiers.Specifier) DatabaseBuilder
	WithName(name string) DatabaseBuilder
	WithNames(names []string) DatabaseBuilder
	Now() (Database, error)
}

// Database represents a database selector
type Database interface {
	Hash() hash.Hash
	Graphbase() specifiers.Specifier
	Content() DatabaseContent
}

// DatabaseContent represents a database content
type DatabaseContent interface {
	IsSpecifier() bool
	Specifier() specifiers.Specifier
	IsName() bool
	Name() string
	IsNames() bool
	Names() []string
}

// SetBuilder represents a set builder
type SetBuilder interface {
	Create() SetBuilder
	WithGraphbase(graphbase specifiers.Specifier) SetBuilder
	WithDatabase(db specifiers.Specifier) SetBuilder
	WithSchema(schema schemas.Schema) SetBuilder
	WithSpecifier(specifier specifiers.Specifier) SetBuilder
	From(from uint) SetBuilder
	To(to uint) SetBuilder
	Now() (Set, error)
}

// Set represents a set
type Set interface {
	Hash() hash.Hash
	Graphbase() specifiers.Specifier
	Database() specifiers.Specifier
	Schema() schemas.Schema
	Content() SetContent
	HasRank() bool
	Rank() SetRank
}

// SetContent represents a set content
type SetContent interface {
	IsSpecifier() bool
	Specifier() specifiers.Specifier
}

// SetRank represents a set rank
type SetRank interface {
	HasFrom() bool
	From() *uint
	HasTo() bool
	To() *uint
}

// TableBuilder represents a table builder
type TableBuilder interface {
	Create() TableBuilder
	WithGraphbase(graphbase specifiers.Specifier) TableBuilder
	WithDatabase(db specifiers.Specifier) TableBuilder
	WithSchema(schema table_schemas.Schema) TableBuilder
	WithSpecifier(specifier specifiers.Specifier) TableBuilder
	Now() (Table, error)
}

// Table represents a table content
type Table interface {
	Hash() hash.Hash
	Graphbase() specifiers.Specifier
	Database() specifiers.Specifier
	Schema() table_schemas.Schema
	Content() TableContent
}

// TableContent represents a table content
type TableContent interface {
	Hash() hash.Hash
	IsSpecifier() bool
	Specifier() specifiers.Specifier
}
