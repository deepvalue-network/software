package parsers

import (
	"errors"
)

type definitionSectionBuilder struct {
	variables VariableSection
}

func createDefinitionSectionBuilder() DefinitionSectionBuilder {
	out := definitionSectionBuilder{
		variables: nil,
	}

	return &out
}

// Create initializes the builder
func (obj *definitionSectionBuilder) Create() DefinitionSectionBuilder {
	return createDefinitionSectionBuilder()
}

// WithVariables add variables to the builder
func (obj *definitionSectionBuilder) WithVariables(variables VariableSection) DefinitionSectionBuilder {
	obj.variables = variables
	return obj
}

// Now builds a new DefinitionSection instance
func (obj *definitionSectionBuilder) Now() (DefinitionSection, error) {
	if obj.variables != nil {
		return createDefinitionSectionWithVariables(obj.variables), nil
	}

	return nil, errors.New("the DefinitionSection is invalid")
}
