package tests

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/tests/test"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	testAdapter test.Adapter
	builder     Builder
}

func createAdapter(
	testAdapter test.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		testAdapter: testAdapter,
		builder:     builder,
	}

	return &out
}

// ToTests converts a testSection to tests
func (app *adapter) ToTests(section parsers.TestSection) (Tests, error) {
	lst := []test.Test{}
	declarations := section.Declarations()
	for _, oneDeclaration := range declarations {
		test, err := app.testAdapter.ToTest(oneDeclaration)
		if err != nil {
			return nil, err
		}

		lst = append(lst, test)
	}

	return app.builder.Create().WithList(lst).Now()
}
