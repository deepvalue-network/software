package parsers

type variableName struct {
	global string
	local  string
}

func createVariableNameWithGlobal(global string) VariableName {
	return createVariableNameInternally(global, "")
}

func createVariableNameWithLocal(local string) VariableName {
	return createVariableNameInternally("", local)
}

func createVariableNameInternally(global string, local string) VariableName {
	out := variableName{
		global: global,
		local:  local,
	}

	return &out
}

// IsGlobal returns true if the variable is global, false otherwise
func (obj *variableName) IsGlobal() bool {
	return obj.global != ""
}

// Global returns the global variable, if any
func (obj *variableName) Global() string {
	return obj.global
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
	if obj.IsGlobal() {
		return obj.Global()
	}

	return obj.Local()
}
