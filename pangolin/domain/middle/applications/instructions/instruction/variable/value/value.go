package value

import "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"

type value struct {
	comp     computable.Value
	variable string
}

func createValueWithComputable(comp computable.Value) Value {
	return createValueInternally(comp, "")
}

func createValueWithVariabe(variable string) Value {
	return createValueInternally(nil, variable)
}

func createValueInternally(comp computable.Value, variable string) Value {
	out := value{
		comp:     comp,
		variable: variable,
	}

	return &out
}

// IsComputable returns true if the value is computable, false otherwise
func (obj *value) IsComputable() bool {
	return obj.comp != nil
}

// Computable returns the computable value, if any
func (obj *value) Computable() computable.Value {
	return obj.comp
}

// IsVariable returns true if the value is variable, false otherwise
func (obj *value) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *value) Variable() string {
	return obj.variable
}
