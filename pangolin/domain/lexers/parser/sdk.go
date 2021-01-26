package parsers

import (
	"github.com/steve-care-software/products/pangolin/domain/lexers"
)

// NewEventsAdapter creates a new eventsAdapter
func NewEventsAdapter() EventsAdapter {
	eventsBuilder := createEventsBuilder()
	eventBuilder := createEventBuilder()
	return createEventsAdapter(eventsBuilder, eventBuilder)
}

// NewBuilder creates a new parser builder
func NewBuilder() Builder {
	eventsBuilder := createEventsBuilder()
	eventBuilder := createEventBuilder()
	eventsAdapter := createEventsAdapter(eventsBuilder, eventBuilder)
	return createBuilder(eventsAdapter)
}

// NewApplication creates a new parser application
func NewApplication() Application {
	return createApplication()
}

// EventFn represents an event func
type EventFn func(tree lexers.NodeTree) (interface{}, error)

// SetFn represents a set func
type SetFn func(code string, ins interface{})

// RetrieveReplacementsFn represents a retrieve replacements func
// first keyname == token, section keyname == code, interface{} == instance to add
type RetrieveReplacementsFn func(tree lexers.NodeTree) (map[string]map[string]interface{}, error)

// ToEventsParams represents the toEventsParams
type ToEventsParams struct {
	Token               string
	Set                 SetFn
	OnEnter             EventFn
	OnExit              EventFn
	RetrieveReplacement RetrieveReplacementsFn
}

// Application represents a parser application
type Application interface {
	Execute(parser Parser) (interface{}, error)
}

// Builder represents a parser builder
type Builder interface {
	Create() Builder
	WithLexer(lexer lexers.Lexer) Builder
	WithEventParams(params []ToEventsParams) Builder
	WithReplacements(replacements map[string]RetrieveReplacementsFn) Builder
	Now() (Parser, error)
}

// Parser represents a parser
type Parser interface {
	Lexer() lexers.Lexer
	Events() Events
}

// EventsAdapter represents an events adapter
type EventsAdapter interface {
	ToEvents(params []ToEventsParams) (Events, error)
}

// EventsBuilder represents an events builder
type EventsBuilder interface {
	Create() EventsBuilder
	WithEvents(evts []Event) EventsBuilder
	Now() (Events, error)
}

// Events represents events
type Events interface {
	Events() []Event
	GetByToken(token string) (Event, error)
}

// EventBuilder represents an event builder
type EventBuilder interface {
	Create() EventBuilder
	WithToken(token string) EventBuilder
	WithSet(set SetFn) EventBuilder
	WithOnEnter(onEnter EventFn) EventBuilder
	WithOnExit(onExit EventFn) EventBuilder
	WithRetrieveReplacement(retrieveReplacement RetrieveReplacementsFn) EventBuilder
	Now() (Event, error)
}

// Event represents a parser event
type Event interface {
	Token() string
	HasSet() bool
	Set() SetFn
	HasOnEnter() bool
	OnEnter() EventFn
	HasOnExit() bool
	OnExit() EventFn
	Fn(isEnter bool) EventFn
	HasRetrieveReplacement() bool
	RetrieveReplacement() RetrieveReplacementsFn
}
