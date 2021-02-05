package tests

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/tests/test"
)

type tests struct {
	lst []test.Test
}

func createTests(lst []test.Test) Tests {
	out := tests{
		lst: lst,
	}

	return &out
}

// All return all the tests
func (obj *tests) All() []test.Test {
	return obj.lst
}
