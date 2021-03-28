package parsers

type variableName struct {
	local string
}

func createVariableNameWithLocal(local string) VariableName {
	return createVariableNameInternally(local)
}

func createVariableNameInternally(local string) VariableName {
	out := variableName{
		local: local,
	}

	return &out
}

// IsLocal returns true if the variable is local, false otherwise
func (obj *variableName) IsLocal() bool {
	return obj.local != ""
}

// Local returns the local variable, if any
func (obj *variableName) Local() string {
	return obj.local
}

// String returns the stirng
func (obj *variableName) String() string {
	return obj.Local()
}
