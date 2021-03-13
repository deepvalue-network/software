package trigger

type trigger struct {
	variable string
	event    string
}

func createTrigger(
	variable string,
	event string,
) Trigger {
	out := trigger{
		variable: variable,
		event:    event,
	}

	return &out
}

// Variable returns the variable name
func (obj *trigger) Variable() string {
	return obj.variable
}

// Event returns the event name
func (obj *trigger) Event() string {
	return obj.event
}
