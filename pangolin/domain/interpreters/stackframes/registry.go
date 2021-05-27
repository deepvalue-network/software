package stackframes

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
)

type registry struct {
	variables map[string]computable.Value
}

func createRegistry() Registry {
	out := registry{
		variables: map[string]computable.Value{},
	}

	return &out
}

// All returns all variables
func (app *registry) All() map[string]computable.Value {
	return app.variables
}

// Fetch fetches a value by name
func (app *registry) Fetch(name string) (computable.Value, error) {
	if val, ok := app.variables[name]; ok {
		return val, nil
	}

	str := fmt.Sprintf("the variable (name: %s) is not declared in the registry", name)
	return nil, errors.New(str)
}

// Insert inserts a value by name
func (app *registry) Insert(name string, val computable.Value) error {
	if _, ok := app.variables[name]; ok {
		str := fmt.Sprintf("the variable (name: %s) is already declared in the registry", name)
		return errors.New(str)
	}

	app.variables[name] = val
	return nil
}

// Delete deletes a value by name
func (app *registry) Delete(name string) error {
	if _, ok := app.variables[name]; ok {
		delete(app.variables, name)
		return nil
	}

	str := fmt.Sprintf("the variable (name: %s) is not declared in the registry", name)
	return errors.New(str)
}
