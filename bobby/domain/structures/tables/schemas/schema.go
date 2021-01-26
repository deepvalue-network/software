package schemas

import "github.com/steve-care-software/products/bobby/domain/resources"

type schema struct {
	resource   resources.Accessible
	name       string
	properties Properties
}

func createSchema(
	resource resources.Accessible,
	name string,
	properties Properties,
) Schema {
	out := schema{
		resource:   resource,
		name:       name,
		properties: properties,
	}

	return &out
}

// Resource returns the resource
func (obj *schema) Resource() resources.Accessible {
	return obj.resource
}

// Name returns the name
func (obj *schema) Name() string {
	return obj.name
}

// Properties returns the properties
func (obj *schema) Properties() Properties {
	return obj.properties
}
