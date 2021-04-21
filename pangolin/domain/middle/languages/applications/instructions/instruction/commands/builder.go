package commands

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/mains"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/tests"
)

type builder struct {
	lang   languages.Language
	script scripts.Script
	head   heads.Head
	main   mains.Main
	test   tests.Test
	label  labels.Label
}

func createBuilder() Builder {
	out := builder{
		lang:   nil,
		script: nil,
		head:   nil,
		main:   nil,
		test:   nil,
		label:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLanguage adds a language to the builder
func (app *builder) WithLanguage(lang languages.Language) Builder {
	app.lang = lang
	return app
}

// WithScript adds a script to the builder
func (app *builder) WithScript(script scripts.Script) Builder {
	app.script = script
	return app
}

// WithHead adds a head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithMain adds a main to the builder
func (app *builder) WithMain(main mains.Main) Builder {
	app.main = main
	return app
}

// WithTest adds a test to the builder
func (app *builder) WithTest(test tests.Test) Builder {
	app.test = test
	return app
}

// WithLabel adds a label to the builder
func (app *builder) WithLabel(label labels.Label) Builder {
	app.label = label
	return app
}

// Now builds a new Command instance
func (app *builder) Now() (Command, error) {
	if app.lang != nil {
		return createCommandWithLanguage(app.lang), nil
	}

	if app.script != nil {
		return createCommandWithScript(app.script), nil
	}

	if app.head != nil {
		return createCommandWithHead(app.head), nil
	}

	if app.main != nil {
		return createCommandWithMain(app.main), nil
	}

	if app.test != nil {
		return createCommandWithTest(app.test), nil
	}

	if app.label != nil {
		return createCommandWithLabel(app.label), nil
	}

	return nil, errors.New("the Command is invalid")
}
