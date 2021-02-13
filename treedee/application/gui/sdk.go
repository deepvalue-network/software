package gui

import (
	"github.com/deepvalue-network/software/treedee/domain/windows"
	"github.com/deepvalue-network/software/treedee/domain/worlds"
)

// Application represents a gui application
type Application interface {
	Execute(win windows.Window, world worlds.World) error
}
