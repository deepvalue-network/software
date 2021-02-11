package events

import "github.com/deepvalue-network/software/libs/hash"

type event struct {
	hash       hash.Hash
	identifier int
	onEnter    EventFn
	onExit     EventFn
}

func createEventWithOnEnter(
	hash hash.Hash,
	identifier int,
	onEnter EventFn,
) Event {
	return createEventInternally(hash, identifier, onEnter, nil)
}

func createEventWithOnExit(
	hash hash.Hash,
	identifier int,
	onExit EventFn,
) Event {
	return createEventInternally(hash, identifier, nil, onExit)
}

func createEventWithOnEnterAndOnExit(
	hash hash.Hash,
	identifier int,
	onEnter EventFn,
	onExit EventFn,
) Event {
	return createEventInternally(hash, identifier, onEnter, onExit)
}

func createEventInternally(
	hash hash.Hash,
	identifier int,
	onEnter EventFn,
	onExit EventFn,
) Event {
	out := event{
		hash:       hash,
		identifier: identifier,
		onEnter:    onEnter,
		onExit:     onExit,
	}

	return &out
}

// Hash returns the hash
func (obj *event) Hash() hash.Hash {
	return obj.hash
}

// Identifier returns the identifier
func (obj *event) Identifier() int {
	return obj.identifier
}

// HasOnEnter returns true if there is an onEnter func, false otherwise
func (obj *event) HasOnEnter() bool {
	return obj.onEnter != nil
}

// OnEnter returns the onEnter func, if any
func (obj *event) OnEnter() EventFn {
	return obj.onEnter
}

// HasOnExit returns true if there is an onExit func, false otherwise
func (obj *event) HasOnExit() bool {
	return obj.onExit != nil
}

// OnExit returns the onExit func, if any
func (obj *event) OnExit() EventFn {
	return obj.onExit
}
