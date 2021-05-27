package commands

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/mains"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/tests"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	headAdapter := heads.NewAdapter()
	labelAdapter := labels.NewAdapter()
	languageAdapter := languages.NewAdapter()
	mainAdapter := mains.NewAdapter()
	scriptAdapter := scripts.NewAdapter()
	testAdapter := tests.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		headAdapter,
		labelAdapter,
		languageAdapter,
		mainAdapter,
		scriptAdapter,
		testAdapter,
		builder,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents command adapter
type Adapter interface {
	ToCommand(parsed parsers.Command) (Command, error)
}

// Builder represents a command builder
type Builder interface {
	Create() Builder
	WithLanguage(lang languages.Language) Builder
	WithScript(script scripts.Script) Builder
	WithHead(head heads.Head) Builder
	WithMain(main mains.Main) Builder
	WithTest(test tests.Test) Builder
	WithLabel(label labels.Label) Builder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	IsLanguage() bool
	Language() languages.Language
	IsScript() bool
	Script() scripts.Script
	IsHead() bool
	Head() heads.Head
	IsMain() bool
	Main() mains.Main
	IsTest() bool
	Test() tests.Test
	IsLabel() bool
	Label() labels.Label
}
