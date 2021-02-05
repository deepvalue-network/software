package labels

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/labels/label"
)

type labels struct {
	lst []label.Label
}

func createLabels(lst []label.Label) Labels {
	out := labels{
		lst: lst,
	}

	return &out
}

// All return all the labels
func (obj *labels) All() []label.Label {
	return obj.lst
}
