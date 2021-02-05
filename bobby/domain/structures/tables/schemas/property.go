package schemas

import "github.com/deepvalue-network/software/bobby/domain/resources"

type property struct {
	resource resources.Accessible
	content  PropertyContent
	name     string
}

func createProperty(
	resource resources.Accessible,
	content PropertyContent,
	name string,
) Property {
	out := property{
		resource: resource,
		content:  content,
		name:     name,
	}

	return &out
}

// Resource returns the resource
func (obj *property) Resource() resources.Accessible {
	return obj.resource
}

// Content returns the content
func (obj *property) Content() PropertyContent {
	return obj.content
}

// Name returns the name
func (obj *property) Name() string {
	return obj.name
}
