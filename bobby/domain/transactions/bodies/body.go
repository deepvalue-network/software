package bodies

import (
	"time"

	"github.com/deepvalue-network/software/bobby/domain/resources"
)

type body struct {
	resource   resources.Immutable
	content    Content
	executesOn *time.Time
}

func createBody(
	resource resources.Immutable,
	content Content,
) Body {
	return createBodyInternally(resource, content, nil)
}

func createBodyWithExecutesOn(
	resource resources.Immutable,
	content Content,
	executesOn *time.Time,
) Body {
	return createBodyInternally(resource, content, executesOn)
}

func createBodyInternally(
	resource resources.Immutable,
	content Content,
	executesOn *time.Time,
) Body {
	out := body{
		resource:   resource,
		content:    content,
		executesOn: executesOn,
	}

	return &out
}

// Resource returns the resource
func (obj *body) Resource() resources.Immutable {
	return obj.resource
}

// Content returns the content
func (obj *body) Content() Content {
	return obj.content
}

// HasExecutesOn returns true if there is an execution time, false otherwise
func (obj *body) HasExecutesOn() bool {
	return obj.executesOn != nil
}

// ExecutesOn returns the executes on time, if any
func (obj *body) ExecutesOn() *time.Time {
	return obj.executesOn
}
