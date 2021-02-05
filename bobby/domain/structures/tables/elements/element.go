package elements

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/values"
)

type element struct {
	resource resources.Resource
	property schemas.Property
	value    values.Value
}

func createElement(
	resource resources.Resource,
	property schemas.Property,
	value values.Value,
) Element {
	out := element{
		resource: resource,
		property: property,
		value:    value,
	}

	return &out
}

// Resource returns the resource
func (obj *element) Resource() resources.Resource {
	return obj.resource
}

// Property returns the property
func (obj *element) Property() schemas.Property {
	return obj.property
}

// Value returns the value
func (obj *element) Value() values.Value {
	return obj.value
}
