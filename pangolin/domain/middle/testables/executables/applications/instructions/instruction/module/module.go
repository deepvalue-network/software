package module

type module struct {
	stackFrame string
	name       string
	symbol     string
}

func createModule(
	stackFrame string,
	name string,
	symbol string,
) Module {
	out := module{
		stackFrame: stackFrame,
		name:       name,
		symbol:     symbol,
	}

	return &out
}

// StackFrame returns the stackFrame
func (obj *module) StackFrame() string {
	return obj.stackFrame
}

// Name returns the name
func (obj *module) Name() string {
	return obj.name
}

// Symbol returns the symbol
func (obj *module) Symbol() string {
	return obj.symbol
}
