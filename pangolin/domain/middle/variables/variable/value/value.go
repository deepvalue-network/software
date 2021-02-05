package value

import "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"

type value struct {
	comp           computable.Value
	globalVariable string
	localVariable  string
}

func createValueWithComputable(comp computable.Value) Value {
	return createValueInternally(comp, "", "")
}

func createValueWithGlobalVariabe(globalVariable string) Value {
	return createValueInternally(nil, globalVariable, "")
}

func createValueWithLocalVariabe(localVariable string) Value {
	return createValueInternally(nil, "", localVariable)
}

func createValueInternally(comp computable.Value, globalVariable string, localVariable string) Value {
	out := value{
		comp:           comp,
		globalVariable: globalVariable,
		localVariable:  localVariable,
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

// IsGlobalVariable returns true if the value is globalVariable, false otherwise
func (obj *value) IsGlobalVariable() bool {
	return obj.globalVariable != ""
}

// GlobalVariable returns the globalVariable, if any
func (obj *value) GlobalVariable() string {
	return obj.globalVariable
}

// IsLocalVariable returns true if the value is localVariable, false otherwise
func (obj *value) IsLocalVariable() bool {
	return obj.localVariable != ""
}

// LocalVariable returns the localVariable, if any
func (obj *value) LocalVariable() string {
	return obj.localVariable
}
