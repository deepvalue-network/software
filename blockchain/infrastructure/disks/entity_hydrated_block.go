package disks

import (
	"github.com/deepvalue-network/software/libs/hashtree"
)

// EntityHydratedBlock represents an entity hydrated block
type EntityHydratedBlock struct {
	Tree *hashtree.JSONCompact `json:"tree" hydro:"0"`
}

func blockOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if tree, ok := ins.(hashtree.HashTree); ok {
		compact := tree.Compact()
		return hashtree.ToJSON(compact), nil
	}

	return nil, nil
}

func blockOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if js, ok := ins.(*hashtree.JSONCompact); ok {
		return hashtree.ToCompact(js)
	}

	return nil, nil
}
