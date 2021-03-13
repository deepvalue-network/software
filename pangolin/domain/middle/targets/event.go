package targets

type event struct {
	name  string
	label string
}

func createEvent(
	name string,
	label string,
) Event {
	out := event{
		name:  name,
		label: label,
	}

	return &out
}

// Name returns the name
func (obj *event) Name() string {
	return obj.name
}

// Label returns the label
func (obj *event) Label() string {
	return obj.label
}
