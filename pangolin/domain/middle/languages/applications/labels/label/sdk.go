package label

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions"

// Label represents a label
type Label interface {
	Name() string
	Instructions() instructions.Instructions
}
