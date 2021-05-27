package tests

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test"
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

// ToTests converts a parsed language test section to tests
func (app *adapter) ToTests(parsed parsers.LanguageTestSection) (Tests, error) {
	list := []test.Test{}
	declarations := parsed.Declarations()
	for _, oneDeclaration := range declarations {
		test, err := app.testAdapter.ToTest(oneDeclaration)
		if err != nil {
			return nil, err
		}

		list = append(list, test)
	}

	return app.builder.Create().WithList(list).Now()
}
