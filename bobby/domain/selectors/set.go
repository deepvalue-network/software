package selectors

import (
	"github.com/steve-care-software/products/bobby/domain/selectors/specifiers"
	"github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
	"github.com/steve-care-software/products/libs/hash"
)

type set struct {
	hash      hash.Hash
	graphbase specifiers.Specifier
	db        specifiers.Specifier
	schema    schemas.Schema
	content   SetContent
	rank      SetRank
}

func createSet(
	hash hash.Hash,
	graphbase specifiers.Specifier,
	db specifiers.Specifier,
	schema schemas.Schema,
	content SetContent,
) Set {
	return createSetInternally(hash, graphbase, db, schema, content, nil)
}

func createSetWithRank(
	hash hash.Hash,
	graphbase specifiers.Specifier,
	db specifiers.Specifier,
	schema schemas.Schema,
	content SetContent,
	rank SetRank,
) Set {
	return createSetInternally(hash, graphbase, db, schema, content, rank)
}

func createSetInternally(
	hash hash.Hash,
	graphbase specifiers.Specifier,
	db specifiers.Specifier,
	schema schemas.Schema,
	content SetContent,
	rank SetRank,
) Set {
	out := set{
		hash:      hash,
		graphbase: graphbase,
		db:        db,
		schema:    schema,
		content:   content,
		rank:      rank,
	}

	return &out
}

// Hash returns the hash
func (obj *set) Hash() hash.Hash {
	return obj.hash
}

// Graphbase returns the graphbase
func (obj *set) Graphbase() specifiers.Specifier {
	return obj.graphbase
}

// Database returns the database
func (obj *set) Database() specifiers.Specifier {
	return obj.db
}

// Schema returns the schema
func (obj *set) Schema() schemas.Schema {
	return obj.schema
}

// Content returns the content
func (obj *set) Content() SetContent {
	return obj.content
}

// HasRank returns true if there is a rank, false otherwise
func (obj *set) HasRank() bool {
	return obj.rank != nil
}

// Rank returns the rank, if any
func (obj *set) Rank() SetRank {
	return obj.rank
}
