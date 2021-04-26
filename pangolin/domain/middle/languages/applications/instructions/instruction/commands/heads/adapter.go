package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/externals"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	externalAdapter externals.Adapter
	builder         Builder
	valueBuilder    ValueBuilder
}

func createAdapter(
	externalAdapter externals.Adapter,
	builder Builder,
	valueBuilder ValueBuilder,
) Adapter {
	out := adapter{
		externalAdapter: externalAdapter,
		builder:         builder,
		valueBuilder:    valueBuilder,
	}

	return &out
}

// ToHead converts a parsed head command to a head instance
func (app *adapter) ToHead(parsed parsers.HeadCommand) (Head, error) {
	values := []Value{}
	headValues := parsed.Values()
	for _, oneHeadValue := range headValues {
		valueBuilder := app.valueBuilder.Create()
		if oneHeadValue.IsName() {
			name := oneHeadValue.Name()
			valueBuilder.WithName(name)
		}

		if oneHeadValue.IsVersion() {
			version := oneHeadValue.Version()
			valueBuilder.WithVersion(version)
		}

		if oneHeadValue.IsImport() {
			imports := oneHeadValue.Import()
			externals, err := app.externalAdapter.ToExternals(imports)
			if err != nil {
				return nil, err
			}

			valueBuilder.WithImports(externals)
		}

		value, err := valueBuilder.Now()
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	variable := parsed.Variable()
	return app.builder.Create().
		WithVariable(variable).
		WithValues(values).
		Now()
}
