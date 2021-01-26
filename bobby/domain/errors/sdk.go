package errors

import (
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/libs/hash"
)

const (
	// CannotProcessTrx represents the cannot process transaction code
	CannotProcessTrx = iota
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := resources.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents an error builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithCode(code uint) Builder
	WithParent(parent Error) Builder
	Now() (Error, error)
}

// Error represents an error
type Error interface {
	Resource() resources.Immutable
	Message() string
	Code() uint
	HasParent() bool
	Parent() Error
}
