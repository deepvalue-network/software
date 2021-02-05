package schemas

import "github.com/deepvalue-network/software/bobby/domain/resources"

type properties struct {
	resource resources.Accessible
	list     []Property
}

func createProperties(
	resource resources.Accessible,
	list []Property,
) Properties {
	out := properties{
		resource: resource,
		list:     list,
	}

	return &out
}

// Resource returns the resource
func (obj *properties) Resource() resources.Accessible {
	return obj.resource
}

// All returns the properties list
func (obj *properties) All() []Property {
	return obj.list
}

// First retruns the first property
func (obj *properties) First() Property {
	if obj.IsEmpty() {
		return nil
	}

	return obj.list[0]
}

// IsEmpty returns true if the properties list is empty, false otherwise
func (obj *properties) IsEmpty() bool {
	return len(obj.list) <= 0
}
