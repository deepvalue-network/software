package parsers

type definitionSection struct {
    constants ConstantSection
    variables VariableSection
}

func createDefinitionSectionWithConstants(
    constants ConstantSection,
) DefinitionSection {
    return createDefinitionSectionInternally(constants, nil)
}


func createDefinitionSectionWithVariables(
    variables VariableSection,
) DefinitionSection {
    return createDefinitionSectionInternally(nil, variables)
}

func createDefinitionSectionWithConstantsWithVariables(
    constants ConstantSection,
    variables VariableSection,
) DefinitionSection {
    return createDefinitionSectionInternally(constants, variables)
}

func createDefinitionSectionInternally(
    constants ConstantSection,
    variables VariableSection,
    ) DefinitionSection {
    out := definitionSection{
        constants: constants,
        variables: variables,
    }

    return &out
}

// HasConstants returns true if threre is constants, false otherwise
func (obj *definitionSection) HasConstants() bool {
    return obj.constants != nil
}

// Constants returns the constants, if any
func (obj *definitionSection) Constants() ConstantSection {
    return obj.constants
}

// HasVariables returns true if threre is variables, false otherwise
func (obj *definitionSection) HasVariables() bool {
    return obj.variables != nil
}

// Variables returns the variables, if any
func (obj *definitionSection) Variables() VariableSection {
    return obj.variables
}
