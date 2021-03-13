package targets

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTargetBuilder creates a new target builder instance
func NewTargetBuilder() TargetBuilder {
	return createTargetBuilder()
}

// NewEventBuilder creates a new event builder instance
func NewEventBuilder() EventBuilder {
	return createEventBuilder()
}

// Builder represents a targets builder
type Builder interface {
	Create() Builder
	WithTargets(targets []Target) Builder
	Now() (Targets, error)
}

// Targets represents a targets
type Targets interface {
	All() []Target
	Fetch(name string) (Target, error)
}

// TargetBuilder represents a target builder
type TargetBuilder interface {
	Create() TargetBuilder
	WithName(name string) TargetBuilder
	WithPath(path string) TargetBuilder
	WithEvents(events []Event) TargetBuilder
	Now() (Target, error)
}

// Target represents a target
type Target interface {
	Name() string
	Path() string
	Events() []Event
}

// EventBuilder represents an event builder
type EventBuilder interface {
	Create() EventBuilder
	WithName(name string) EventBuilder
	WithLabel(label string) EventBuilder
	Now() (Event, error)
}

// Event represents an event
type Event interface {
	Name() string
	Label() string
}
