package labels

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label"

type labels struct {
	list []label.Label
}

func createLabels(
	list []label.Label,
) Labels {
	out := labels{
		list: list,
	}

	return &out
}

// All returns the labels
func (obj *labels) All() []label.Label {
	return obj.list
}
