package selectors

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type table struct {
	hash      hash.Hash
	graphbase specifiers.Specifier
	database  specifiers.Specifier
	schema    schemas.Schema
	content   TableContent
}

func createTable(
	hash hash.Hash,
	graphbase specifiers.Specifier,
	database specifiers.Specifier,
	schema schemas.Schema,
	content TableContent,
) Table {
	out := table{
		hash:      hash,
		graphbase: graphbase,
		database:  database,
		schema:    schema,
		content:   content,
	}

	return &out
}

// Hash returns the hash
func (obj *table) Hash() hash.Hash {
	return obj.hash
}

// Graphbase returns the graphbase
func (obj *table) Graphbase() specifiers.Specifier {
	return obj.graphbase
}

// Database returns the database
func (obj *table) Database() specifiers.Specifier {
	return obj.database
}

// Schema returns the schema
func (obj *table) Schema() schemas.Schema {
	return obj.schema
}

// Content returns the content
func (obj *table) Content() TableContent {
	return obj.content
}
