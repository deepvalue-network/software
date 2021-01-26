package parsers

import (
	"errors"
)

type definitionSectionBuilder struct {
	constants ConstantSection
	variables VariableSection
}

func createDefinitionSectionBuilder() DefinitionSectionBuilder {
	out := definitionSectionBuilder{
		constants: nil,
		variables: nil,
	}

	return &out
}

// Create initializes the builder
func (obj *definitionSectionBuilder) Create() DefinitionSectionBuilder {
	return createDefinitionSectionBuilder()
}

// WithConstants add constants to the builder
func (obj *definitionSectionBuilder) WithConstants(constants ConstantSection) DefinitionSectionBuilder {
	obj.constants = constants
	return obj
}

// WithVariables add variables to the builder
func (obj *definitionSectionBuilder) WithVariables(variables VariableSection) DefinitionSectionBuilder {
	obj.variables = variables
	return obj
}

// Now builds a new DefinitionSection instance
func (obj *definitionSectionBuilder) Now() (DefinitionSection, error) {

	if obj.constants != nil && obj.variables != nil {
		return createDefinitionSectionWithConstantsWithVariables(obj.constants, obj.variables), nil
	}

	if obj.constants != nil {
		return createDefinitionSectionWithConstants(obj.constants), nil
	}

	if obj.variables != nil {
		return createDefinitionSectionWithVariables(obj.variables), nil
	}

	return nil, errors.New("the DefinitionSection is invalid")
}
