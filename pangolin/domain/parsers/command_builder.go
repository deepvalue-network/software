package parsers

import "errors"

type commandBuilder struct {
	language LanguageCommand
	script   ScriptCommand
	head     HeadCommand
	main     MainCommand
	label    LabelCommand
	test     TestCommand
}

func createCommandBuilder() CommandBuilder {
	out := commandBuilder{
		language: nil,
		script:   nil,
		head:     nil,
		main:     nil,
		label:    nil,
		test:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *commandBuilder) Create() CommandBuilder {
	return createCommandBuilder()
}

// WithLanguage adds a language to the builder
func (app *commandBuilder) WithLanguage(language LanguageCommand) CommandBuilder {
	app.language = language
	return app
}

// WithScript adds a script to the builder
func (app *commandBuilder) WithScript(script ScriptCommand) CommandBuilder {
	app.script = script
	return app
}

// WithHead adds an head to the builder
func (app *commandBuilder) WithHead(head HeadCommand) CommandBuilder {
	app.head = head
	return app
}

// WithMain adds a main to the builder
func (app *commandBuilder) WithMain(main MainCommand) CommandBuilder {
	app.main = main
	return app
}

// WithLabel adds a label to the builder
func (app *commandBuilder) WithLabel(label LabelCommand) CommandBuilder {
	app.label = label
	return app
}

// WithTest adds a test to the builder
func (app *commandBuilder) WithTest(test TestCommand) CommandBuilder {
	app.test = test
	return app
}

// Now builds a new Command instance
func (app *commandBuilder) Now() (Command, error) {
	if app.language != nil {
		return createCommandWithLanguage(app.language), nil
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

	if app.label != nil {
		return createCommandWithLabel(app.label), nil
	}

	if app.test != nil {
		return createCommandWithTest(app.test), nil
	}

	return nil, errors.New("the Command is invalid")
}
