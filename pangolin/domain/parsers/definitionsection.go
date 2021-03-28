package parsers

type definitionSection struct {
	variables VariableSection
}

func createDefinitionSectionWithVariables(
	variables VariableSection,
) DefinitionSection {
	return createDefinitionSectionInternally(variables)
}

func createDefinitionSectionInternally(
	variables VariableSection,
) DefinitionSection {
	out := definitionSection{
		variables: variables,
	}

	return &out
}

// HasVariables returns true if threre is variables, false otherwise
func (obj *definitionSection) HasVariables() bool {
	return obj.variables != nil
}

// Variables returns the variables, if any
func (obj *definitionSection) Variables() VariableSection {
	return obj.variables
}
