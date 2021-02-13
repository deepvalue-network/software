package windows

import (
	"time"

	"github.com/deepvalue-network/software/treedee/domain/windows"
)

// UpdateFn represents the update func
type UpdateFn func(prev time.Time, current time.Time) error

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithWindow(win windows.Window) Builder
	Now() (Application, error)
}

// Application represents a window application
type Application interface {
	Execute(fn UpdateFn) error
}
