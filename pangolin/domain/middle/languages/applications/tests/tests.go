package tests

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test"

type tests struct {
	list []test.Test
}

func createTests(
	list []test.Test,
) Tests {
	out := tests{
		list: list,
	}

	return &out
}

// All returns the tests
func (obj *tests) All() []test.Test {
	return obj.list
}
