package tests

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test"

// Tests represents tests
type Tests interface {
	All() []test.Test
}
