package variables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	variableAdapter variable.Adapter
	builder         Builder
}

func createAdapter(variableAdapterBuilder variable.AdapterBuilder, builder Builder) Adapter {
	out := adapter{
		variableAdapter: variableAdapterBuilder.Create().IsGlobal().Now(),
		builder:         builder,
	}

	return &out
}

// FromConstants converts a constant section to a variables
func (app *adapter) FromConstants(section parsers.ConstantSection) (Variables, error) {
	vrs := []variable.Variable{}
	declarations := section.Declarations()
	for _, oneDeclaration := range declarations {
		vr, err := app.variableAdapter.FromConstant(oneDeclaration)
		if err != nil {
			return nil, err
		}

		vrs = append(vrs, vr)
	}

	return app.builder.Create().WithVariables(vrs).Now()
}

// FromVariables converts a variables section to a variables
func (app *adapter) FromVariables(section parsers.VariableSection) (Variables, error) {
	vrs := []variable.Variable{}
	declarations := section.Declarations()
	for _, oneDeclaration := range declarations {
		vr, err := app.variableAdapter.FromVariable(oneDeclaration)
		if err != nil {
			return nil, err
		}

		vrs = append(vrs, vr)
	}

	return app.builder.Create().WithVariables(vrs).Now()
}
