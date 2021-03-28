package parsers

type trigger struct {
	variableName string
	event        string
}

func createTrigger(
	variableName string,
	event string,
) Trigger {
	out := trigger{
		variableName: variableName,
		event:        event,
	}

	return &out
}

// Variable returns the variable name
func (obj *trigger) Variable() string {
	return obj.variableName
}

// Event returns the event
func (obj *trigger) Event() string {
	return obj.event
}
