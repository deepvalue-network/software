package events

import (
	"github.com/deepvalue-network/software/libs/hash"
)

// NewManagerFactory creates a new manager factory
func NewManagerFactory() ManagerFactory {
	return createManagerFactory()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// EventFn represents an event function
type EventFn func(data interface{}, event Event) error

// TriggerFn represents a trigger Fn
type TriggerFn func() error

// ManagerFactory represents a manager factory
type ManagerFactory interface {
	Create() Manager
}

// Manager represents an event manager
type Manager interface {
	Add(evt Event) error
	AddList(evts []Event) error
	Trigger(identifier int, data interface{}, triggerFn TriggerFn) error
}

// Builder represents an event builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier int) Builder
	OnEnter(onEnter EventFn) Builder
	OnExit(onExit EventFn) Builder
	Now() (Event, error)
}

// Event represents an event
type Event interface {
	Hash() hash.Hash
	Identifier() int
	HasOnEnter() bool
	OnEnter() EventFn
	HasOnExit() bool
	OnExit() EventFn
}
