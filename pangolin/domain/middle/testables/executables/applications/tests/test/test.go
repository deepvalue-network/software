package test

import "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions"

type test struct {
	name string
	ins  instructions.Instructions
}

func createTest(
	name string,
	ins instructions.Instructions,
) Test {
	out := test{
		name: name,
		ins:  ins,
	}

	return &out
}

// Name returns the name
func (obj *test) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *test) Instructions() instructions.Instructions {
	return obj.ins
}
