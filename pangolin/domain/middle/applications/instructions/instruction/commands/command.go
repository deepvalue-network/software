package commands

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/mains"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/tests"
)

type command struct {
	lang   languages.Language
	script scripts.Script
	head   heads.Head
	main   mains.Main
	test   tests.Test
	label  labels.Label
}

func createCommandWithLanguage(
	lang languages.Language,
) Command {
	return createCommandInternally(lang, nil, nil, nil, nil, nil)
}

func createCommandWithScript(
	script scripts.Script,
) Command {
	return createCommandInternally(nil, script, nil, nil, nil, nil)
}

func createCommandWithHead(
	head heads.Head,
) Command {
	return createCommandInternally(nil, nil, head, nil, nil, nil)
}

func createCommandWithMain(
	main mains.Main,
) Command {
	return createCommandInternally(nil, nil, nil, main, nil, nil)
}

func createCommandWithTest(
	test tests.Test,
) Command {
	return createCommandInternally(nil, nil, nil, nil, test, nil)
}

func createCommandWithLabel(
	label labels.Label,
) Command {
	return createCommandInternally(nil, nil, nil, nil, nil, label)
}

func createCommandInternally(
	lang languages.Language,
	script scripts.Script,
	head heads.Head,
	main mains.Main,
	test tests.Test,
	label labels.Label,
) Command {
	out := command{
		lang:   lang,
		script: script,
		head:   head,
		main:   main,
		test:   test,
		label:  label,
	}

	return &out
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *command) IsLanguage() bool {
	return obj.lang != nil
}

// Language returns the language, if any
func (obj *command) Language() languages.Language {
	return obj.lang
}

// IsScript returns true if there is a script, false otherwise
func (obj *command) IsScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *command) Script() scripts.Script {
	return obj.script
}

// IsHead returns true if there is a head, false otherwise
func (obj *command) IsHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *command) Head() heads.Head {
	return obj.head
}

// IsMain returns true if there is a main, false otherwise
func (obj *command) IsMain() bool {
	return obj.main != nil
}

// Main returns the main, if any
func (obj *command) Main() mains.Main {
	return obj.main
}

// Test returns true if there is a test, false otherwise
func (obj *command) IsTest() bool {
	return obj.test != nil
}

// Test returns the test, if any
func (obj *command) Test() tests.Test {
	return obj.test
}

// IsLabel returns true if there is a label, false otherwise
func (obj *command) IsLabel() bool {
	return obj.label != nil
}

// Label returns the label, if any
func (obj *command) Label() labels.Label {
	return obj.label
}
