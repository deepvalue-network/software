package label

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions"
)

type label struct {
	name string
	ins  instructions.Instructions
}

func createLabel(name string, ins instructions.Instructions) Label {
	out := label{
		name: name,
		ins:  ins,
	}

	return &out
}

// Name returns the name
func (obj *label) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *label) Instructions() instructions.Instructions {
	return obj.ins
}
