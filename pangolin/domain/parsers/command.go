package parsers

type command struct {
	language LanguageCommand
	script   ScriptCommand
	head     HeadCommand
	main     MainCommand
	label    LabelCommand
	test     TestCommand
}

func createCommandWithLanguage(
	language LanguageCommand,
) Command {
	return createCommandInternally(language, nil, nil, nil, nil, nil)
}

func createCommandWithScript(
	script ScriptCommand,
) Command {
	return createCommandInternally(nil, script, nil, nil, nil, nil)
}

func createCommandWithHead(
	head HeadCommand,
) Command {
	return createCommandInternally(nil, nil, head, nil, nil, nil)
}

func createCommandWithMain(
	main MainCommand,
) Command {
	return createCommandInternally(nil, nil, nil, main, nil, nil)
}

func createCommandWithLabel(
	label LabelCommand,
) Command {
	return createCommandInternally(nil, nil, nil, nil, label, nil)
}

func createCommandWithTest(
	test TestCommand,
) Command {
	return createCommandInternally(nil, nil, nil, nil, nil, test)
}

func createCommandInternally(
	language LanguageCommand,
	script ScriptCommand,
	head HeadCommand,
	main MainCommand,
	label LabelCommand,
	test TestCommand,
) Command {
	out := command{
		language: language,
		script:   script,
		head:     head,
		main:     main,
		label:    label,
		test:     test,
	}

	return &out
}

// IsLanguage returns true if there is a language command, false otherwise
func (obj *command) IsLanguage() bool {
	return obj.language != nil
}

// Language returns the language command, if any
func (obj *command) Language() LanguageCommand {
	return obj.language
}

// IsScript returns true if there is a script command, false otherwise
func (obj *command) IsScript() bool {
	return obj.script != nil
}

// Script returns the script command, if any
func (obj *command) Script() ScriptCommand {
	return obj.script
}

// IsHead returns true if there is a head, false otherwise
func (obj *command) IsHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *command) Head() HeadCommand {
	return obj.head
}

// IsMain returns true if there is a main, false otherwise
func (obj *command) IsMain() bool {
	return obj.main != nil
}

// Main returns the main, if any
func (obj *command) Main() MainCommand {
	return obj.main
}

// IsLabel returns true if there is a label, false otherwise
func (obj *command) IsLabel() bool {
	return obj.label != nil
}

// Label returns the label, if any
func (obj *command) Label() LabelCommand {
	return obj.label
}

// IsTest returns true if there is a test, false otherwise
func (obj *command) IsTest() bool {
	return obj.test != nil
}

// Test returns the test, if any
func (obj *command) Test() TestCommand {
	return obj.test
}
