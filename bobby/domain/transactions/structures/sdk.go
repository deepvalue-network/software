package structures

import (
	"time"

	"github.com/steve-care-software/products/bobby/domain/structures"
)

// Builder represents a structure builder
type Builder interface {
	Create() Builder
	WithStructure(structure structures.Structure) Builder
	ExecutesOn(executesOn time.Time) Builder
	ExpiresOn(expiresOn time.Time) Builder
	IsDeleted() Builder
	Now() (Structure, error)
}

// Structure represents a structure
type Structure interface {
	Structure() structures.Structure
	IsActive() bool
	IsDeleted() bool
	HasExecutesOn() bool
	ExecutesOn() *time.Time
	HasExpiresOn() bool
	ExpiresOn() *time.Time
}
