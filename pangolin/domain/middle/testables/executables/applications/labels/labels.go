package labels

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label"
)

type labels struct {
	mp  map[string]label.Label
	lst []label.Label
}

func createLabels(
	mp map[string]label.Label,
	lst []label.Label,
) Labels {
	out := labels{
		mp:  mp,
		lst: lst,
	}

	return &out
}

// All return all the labels
func (obj *labels) All() []label.Label {
	return obj.lst
}

// Fetch fetches a label by name, if any
func (obj *labels) Fetch(name string) (label.Label, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the label (name: %s) does not exists", name)
	return nil, errors.New(str)
}
