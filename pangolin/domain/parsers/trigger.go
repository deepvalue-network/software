package parsers

type trigger struct {
	variableName VariableName
	event        string
}

func createTrigger(
	variableName VariableName,
	event string,
) Trigger {
	out := trigger{
		variableName: variableName,
		event:        event,
	}

	return &out
}

// Variable returns the variable name
func (obj *trigger) Variable() VariableName {
	return obj.variableName
}

// Event returns the event
func (obj *trigger) Event() string {
	return obj.event
}
