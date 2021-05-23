package composers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	command_heads "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/heads"
	command_labels "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/labels"
	command_languages "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/languages"
	command_mains "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/mains"
	command_scripts "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/scripts"
	command_tests "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/tests"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type composer struct {
	instructionAdapterBuilder  InstructionAdapterBuilder
	stackFrameBuilder          stackframes.Builder
	programBuilder             parsers.ProgramBuilder
	applicationBuilder         parsers.ApplicationBuilder
	labelSectionBuilder        parsers.LabelSectionBuilder
	mainSectionBuilder         parsers.MainSectionBuilder
	testSectionBuilder         parsers.TestSectionBuilder
	languageBuilder            parsers.LanguageBuilder
	languageApplicationBuilder parsers.LanguageApplicationBuilder
	languageDefinitionBuilder  parsers.LanguageDefinitionBuilder
	languageValueBuilder       parsers.LanguageValueBuilder
	patternMatchBuilder        parsers.PatternMatchBuilder
	patternLabelsBuilder       parsers.PatternLabelsBuilder
	scriptBuilder              parsers.ScriptBuilder
	scriptValueBuilder         parsers.ScriptValueBuilder
	headSectionBuilder         parsers.HeadSectionBuilder
	headValueBuilder           parsers.HeadValueBuilder
	testDeclarationBuilder     parsers.TestDeclarationBuilder
	labelDeclarationBuilder    parsers.LabelDeclarationBuilder
	stackFrame                 stackframes.StackFrame
	linker                     linkers.Linker
	localTestStackFrames       map[string]stackframes.StackFrame
	localLabelStackFrames      map[string]stackframes.StackFrame
	languageValues             []parsers.LanguageValue
	scriptValues               []parsers.ScriptValue
	headValues                 []parsers.HeadValue
	mainInstructions           []parsers.Instruction
	testDeclarations           []parsers.TestDeclaration
	labelDeclarations          []parsers.LabelDeclaration
}

func createComposer(
	instructionAdapterBuilder InstructionAdapterBuilder,
	stackFrameBuilder stackframes.Builder,
	programBuilder parsers.ProgramBuilder,
	applicationBuilder parsers.ApplicationBuilder,
	labelSectionBuilder parsers.LabelSectionBuilder,
	mainSectionBuilder parsers.MainSectionBuilder,
	testSectionBuilder parsers.TestSectionBuilder,
	languageBuilder parsers.LanguageBuilder,
	languageApplicationBuilder parsers.LanguageApplicationBuilder,
	languageDefinitionBuilder parsers.LanguageDefinitionBuilder,
	languageValueBuilder parsers.LanguageValueBuilder,
	patternMatchBuilder parsers.PatternMatchBuilder,
	patternLabelsBuilder parsers.PatternLabelsBuilder,
	scriptBuilder parsers.ScriptBuilder,
	scriptValueBuilder parsers.ScriptValueBuilder,
	headSectionBuilder parsers.HeadSectionBuilder,
	headValueBuilder parsers.HeadValueBuilder,
	testDeclarationBuilder parsers.TestDeclarationBuilder,
	labelDeclarationBuilder parsers.LabelDeclarationBuilder,
	stackFrame stackframes.StackFrame,
	linker linkers.Linker,
) Composer {
	out := composer{
		instructionAdapterBuilder:  instructionAdapterBuilder,
		stackFrameBuilder:          stackFrameBuilder,
		programBuilder:             programBuilder,
		applicationBuilder:         applicationBuilder,
		labelSectionBuilder:        labelSectionBuilder,
		mainSectionBuilder:         mainSectionBuilder,
		testSectionBuilder:         testSectionBuilder,
		languageBuilder:            languageBuilder,
		languageApplicationBuilder: languageApplicationBuilder,
		languageDefinitionBuilder:  languageDefinitionBuilder,
		languageValueBuilder:       languageValueBuilder,
		patternMatchBuilder:        patternMatchBuilder,
		patternLabelsBuilder:       patternLabelsBuilder,
		scriptBuilder:              scriptBuilder,
		scriptValueBuilder:         scriptValueBuilder,
		headSectionBuilder:         headSectionBuilder,
		headValueBuilder:           headValueBuilder,
		testDeclarationBuilder:     testDeclarationBuilder,
		labelDeclarationBuilder:    labelDeclarationBuilder,
		stackFrame:                 stackFrame,
		linker:                     linker,
		localTestStackFrames:       map[string]stackframes.StackFrame{},
		localLabelStackFrames:      map[string]stackframes.StackFrame{},
		languageValues:             []parsers.LanguageValue{},
		scriptValues:               []parsers.ScriptValue{},
		headValues:                 []parsers.HeadValue{},
		mainInstructions:           []parsers.Instruction{},
		testDeclarations:           []parsers.TestDeclaration{},
		labelDeclarations:          []parsers.LabelDeclaration{},
	}

	return &out
}

// Receive receives a command
func (app *composer) Receive(command commands.Command) error {
	if command.IsLanguage() {
		lang := command.Language()
		err := app.receiveLanguage(lang)
		if err != nil {
			return err
		}
	}

	if command.IsScript() {
		script := command.Script()
		err := app.receiveScript(script)
		if err != nil {
			return err
		}
	}

	if command.IsHead() {
		head := command.Head()
		err := app.receiveHead(head)
		if err != nil {
			return err
		}
	}

	if command.IsMain() {
		main := command.Main()
		err := app.receiveMain(main)
		if err != nil {
			return err
		}
	}

	if command.IsTest() {
		test := command.Test()
		err := app.receiveTest(test)
		if err != nil {
			return err
		}
	}

	if command.IsLabel() {
		label := command.Label()
		err := app.receiveLabel(label)
		if err != nil {
			return err
		}
	}

	return nil
}

// Now parses the software and returns a middle application
func (app *composer) Now() (linkers.Application, error) {
	builder := app.applicationBuilder.Create()
	headSection, err := app.headSectionBuilder.Create().WithValues(app.headValues).Now()
	if err != nil {
		return nil, err
	}

	mainInstructions, err := app.mainSectionBuilder.Create().WithInstructions(app.mainInstructions).Now()
	if err != nil {
		return nil, err
	}

	builder.WithHead(headSection).WithMain(mainInstructions)
	if len(app.labelDeclarations) > 0 {
		labelSection, err := app.labelSectionBuilder.Create().WithDeclarations(app.labelDeclarations).Now()
		if err != nil {
			return nil, err
		}

		builder.WithLabel(labelSection)
	}

	if len(app.testDeclarations) > 0 {
		testSection, err := app.testSectionBuilder.Create().WithDeclarations(app.testDeclarations).Now()
		if err != nil {
			return nil, err
		}

		builder.WithTest(testSection)
	}

	parsedApp, err := builder.Now()
	if err != nil {
		return nil, err
	}

	parsedProg, err := app.programBuilder.Create().WithApplication(parsedApp).Now()
	if err != nil {
		return nil, err
	}

	linkedProg, err := app.linker.Execute(parsedProg)
	if err != nil {
		return nil, err
	}

	if linkedProg.IsApplication() {
		return nil, errors.New("the composed program was expected to be an application")
	}

	return linkedProg.Application(), nil
}

func (app *composer) receiveLabel(label command_labels.Label) error {
	name := label.Name()
	if _, ok := app.localLabelStackFrames[name]; !ok {
		app.localLabelStackFrames[name] = app.stackFrameBuilder.Create().Now()
	}

	instructionAdapter, err := app.instructionAdapterBuilder.Create().WithLocalStackFrame(app.localLabelStackFrames[name]).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return err
	}

	parsedInstructions := []parsers.LabelInstruction{}
	instructions := label.Instructions()
	for _, oneIns := range instructions {
		parsedInsList, err := instructionAdapter.Label(oneIns)
		if err != nil {
			return err
		}

		parsedInstructions = append(parsedInstructions, parsedInsList...)
	}

	labelDeclaration, err := app.labelDeclarationBuilder.Create().WithName(name).WithInstructions(parsedInstructions).Now()
	if err != nil {
		return err
	}

	app.labelDeclarations = append(app.labelDeclarations, labelDeclaration)
	return nil
}

func (app *composer) receiveTest(test command_tests.Test) error {
	name := test.Name()
	if _, ok := app.localTestStackFrames[name]; !ok {
		app.localTestStackFrames[name] = app.stackFrameBuilder.Create().Now()
	}

	instructionAdapter, err := app.instructionAdapterBuilder.Create().WithLocalStackFrame(app.localTestStackFrames[name]).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return err
	}

	parsedInstructions := []parsers.TestInstruction{}
	instructions := test.Instructions()
	for _, oneIns := range instructions {
		parsedInsList, err := instructionAdapter.Test(oneIns)
		if err != nil {
			return err
		}

		parsedInstructions = append(parsedInstructions, parsedInsList...)
	}

	testDeclaration, err := app.testDeclarationBuilder.Create().WithName(name).WithInstructions(parsedInstructions).Now()
	if err != nil {
		return err
	}

	app.testDeclarations = append(app.testDeclarations, testDeclaration)
	return nil
}

func (app *composer) receiveMain(main command_mains.Main) error {
	localStackFrame := app.stackFrameBuilder.Create().Now()
	instructionAdapter, err := app.instructionAdapterBuilder.Create().WithLocalStackFrame(localStackFrame).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return err
	}

	instructions := main.Instructions()
	for _, oneIns := range instructions {
		parsedInsList, err := instructionAdapter.Application(oneIns)
		if err != nil {
			return err
		}

		app.mainInstructions = append(app.mainInstructions, parsedInsList...)
	}

	return nil
}

func (app *composer) receiveHead(head command_heads.Head) error {
	values := head.Values()
	for _, oneValue := range values {
		err := app.receiveHeadValue(oneValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *composer) receiveHeadValue(value command_heads.Value) error {
	valueBuilder := app.headValueBuilder.Create()
	if value.IsName() {
		name := value.Name()
		valueBuilder.WithName(name)
	}

	if value.IsVersion() {
		version := value.Version()
		valueBuilder.WithVersion(version)
	}

	if value.IsImports() {
		imports := value.Imports()
		valueBuilder.WithImport(imports)
	}

	val, err := valueBuilder.Now()
	if err != nil {
		return err
	}

	app.headValues = append(app.headValues, val)
	return nil
}

func (app *composer) receiveScript(script command_scripts.Script) error {
	values := script.Values()
	for _, oneValue := range values {
		err := app.receiveScriptValue(oneValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *composer) receiveScriptValue(script command_scripts.Value) error {
	valueBuilder := app.scriptValueBuilder.Create()
	if script.IsName() {
		name := script.Name()
		valueBuilder.WithName(name)
	}

	if script.IsVersion() {
		version := script.Version()
		valueBuilder.WithVersion(version)
	}

	if script.IsLanguagePath() {
		langPath := script.LanguagePath()
		valueBuilder.WithLanguagePath(langPath)
	}

	if script.IsScriptPath() {
		scriptPath := script.ScriptPath()
		valueBuilder.WithScriptPath(scriptPath)
	}

	val, err := valueBuilder.Now()
	if err != nil {
		return err
	}

	app.scriptValues = append(app.scriptValues, val)
	return nil
}

func (app *composer) receiveLanguage(language command_languages.Language) error {
	values := language.Values()
	for _, oneValue := range values {
		err := app.receiveLanguageValue(oneValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *composer) receiveLanguageValue(value command_languages.Value) error {
	valueBuilder := app.languageValueBuilder.Create()
	if value.IsRoot() {
		root := value.Root()
		valueBuilder.WithRoot(root)
	}

	if value.IsTokensPath() {
		tokensPath := value.TokensPath()
		valueBuilder.WithTokens(tokensPath)
	}

	if value.IsRulesPath() {
		rulesPath := value.RulesPath()
		valueBuilder.WithRules(rulesPath)
	}

	if value.IsLogicsPath() {
		logicsPath := value.LogicsPath()
		valueBuilder.WithLogic(logicsPath)
	}

	if value.IsPatternMatches() {
		//patternMatches := value.PatternMatches()
		//valueBuilder.WithPatternMatches(patternMatches)
	}

	/*

		patternLabelsBuilder      parsers.PatternLabelsBuilder
		patternMatchBuilder       parsers.PatternMatchBuilder

	*/

	if value.IsInputVariable() {
		inputVariable := value.InputVariable()
		valueBuilder.WithInputVariable(inputVariable)
	}

	if value.IsChannelsPath() {
		channelsPath := value.ChannelsPath()
		valueBuilder.WithLogic(channelsPath)
	}

	if value.IsExtends() {
		extends := value.Extends()
		valueBuilder.WithExtends(extends)
	}

	val, err := valueBuilder.Now()
	if err != nil {
		return err
	}

	app.languageValues = append(app.languageValues, val)
	return nil
}
