package labels

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label"

// Labels represents labels
type Labels interface {
	All() []label.Label
}
