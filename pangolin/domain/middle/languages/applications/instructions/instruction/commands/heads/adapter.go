package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/heads"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	builder           Builder
	valueBuilder      ValueBuilder
	loadSingleBuilder heads.LoadSingleBuilder
}

func createAdapter(
	builder Builder,
	valueBuilder ValueBuilder,
	loadSingleBuilder heads.LoadSingleBuilder,
) Adapter {
	out := adapter{
		builder:           builder,
		valueBuilder:      valueBuilder,
		loadSingleBuilder: loadSingleBuilder,
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
			valueBuilder.WithImports(imports)
		}

		if oneHeadValue.IsLoad() {
			loads := []heads.LoadSingle{}
			parsedLoads := oneHeadValue.Load()
			for _, oneParsedLoad := range parsedLoads {
				internal := oneParsedLoad.Internal()
				external := oneParsedLoad.External()
				load, err := app.loadSingleBuilder.Create().WithInternal(internal).WithExternal(external).Now()
				if err != nil {
					return nil, err
				}

				loads = append(loads, load)
			}

			valueBuilder.WithLoads(loads)
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
