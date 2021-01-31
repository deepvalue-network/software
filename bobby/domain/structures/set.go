package structures

import (
	"github.com/steve-care-software/products/bobby/domain/structures/sets"
	set_schemas "github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
)

type set struct {
	schema set_schemas.Schema
	set    sets.Set
}

func createSetWithSchema(
	schema set_schemas.Schema,
) Set {
	return createSetInternally(schema, nil)
}

func createSetWithSet(
	set sets.Set,
) Set {
	return createSetInternally(nil, set)
}

func createSetInternally(
	schema set_schemas.Schema,
	st sets.Set,
) Set {
	out := set{
		schema: schema,
		set:    st,
	}

	return &out
}

// IsSchema returns true if there is a schema, false otherwise
func (obj *set) IsSchema() bool {
	return obj.schema != nil
}

// Schema returns the schema, if any
func (obj *set) Schema() set_schemas.Schema {
	return obj.schema
}

// IsSet returns true if there is a set, false otherwise
func (obj *set) IsSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *set) Set() sets.Set {
	return obj.set
}
