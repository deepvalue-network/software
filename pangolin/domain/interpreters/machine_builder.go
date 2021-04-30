package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type machineBuilder struct {
	computableValueBuilder computable.Builder
	callLabelFunc          CallLabelFunc
	fetchStackFunc         FetchStackFrameFunc
}

func createMachineBuilder(
	computableValueBuilder computable.Builder,
) MachineBuilder {
	out := machineBuilder{
		computableValueBuilder: computableValueBuilder,
		callLabelFunc:          nil,
		fetchStackFunc:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *machineBuilder) Create() MachineBuilder {
	return createMachineBuilder(
		app.computableValueBuilder,
	)
}

// WithCallLabelFunc adds a callLabelFunc func to the builder
func (app *machineBuilder) WithCallLabelFunc(callLabelFunc CallLabelFunc) MachineBuilder {
	app.callLabelFunc = callLabelFunc
	return app
}

// WithFetchStackFunc adds a fetchStackFunc func to the builder
func (app *machineBuilder) WithFetchStackFunc(fetchStackFunc FetchStackFrameFunc) MachineBuilder {
	app.fetchStackFunc = fetchStackFunc
	return app
}

// Now builds a new Machine instance
func (app *machineBuilder) Now() (Machine, error) {
	if app.callLabelFunc == nil {
		return nil, errors.New("the callLabelFunc is mandatory in order to build a Machine instance")
	}

	if app.fetchStackFunc == nil {
		return nil, errors.New("the fetchStackFunc is mandatory in order to build a Machine instance")
	}

	return createMachine(app.computableValueBuilder, app.callLabelFunc, app.fetchStackFunc), nil
}
