package values

import "github.com/deepvalue-network/software/bobby/domain/resources"

type value struct {
	resource resources.Resource
	content  ValueContent
}

func createValue(
	resource resources.Resource,
	content ValueContent,
) Value {
	out := value{
		resource: resource,
		content:  content,
	}

	return &out
}

// Resource returns the resource
func (obj *value) Resource() resources.Resource {
	return obj.resource
}

// Content returns the content
func (obj *value) Content() ValueContent {
	return obj.content
}
