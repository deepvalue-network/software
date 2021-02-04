package disks

import (
	"github.com/steve-care-software/products/libs/hashtree"
)

type entityHydratedBlock struct {
	Tree hashtree.Compact `json:"tree" hydro:"0"`
}

func blockOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if tree, ok := ins.(hashtree.HashTree); ok {
		return tree.Compact(), nil
	}

	return nil, nil
}
