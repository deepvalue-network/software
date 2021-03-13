package targets

type target struct {
	name   string
	path   string
	events []Event
}

func createTarget(
	name string,
	path string,
	events []Event,
) Target {
	out := target{
		name:   name,
		path:   path,
		events: events,
	}

	return &out
}

// Name returns the name
func (obj *target) Name() string {
	return obj.name
}

// Path returns the path
func (obj *target) Path() string {
	return obj.path
}

// Events returns the events
func (obj *target) Events() []Event {
	return obj.events
}
