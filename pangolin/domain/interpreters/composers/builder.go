package composers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	instructionAdapterBuilder  InstructionAdapterBuilder
	stackFrameBuilder          stackframes.Builder
	programBuilder             parsers.ProgramBuilder
	testableBuilder            parsers.TestableBuilder
	executableBuilder          parsers.ExecutableBuilder
	applicationBuilder         parsers.ApplicationBuilder
	labelSectionBuilder        parsers.LabelSectionBuilder
	mainSectionBuilder         parsers.MainSectionBuilder
	testSectionBuilder         parsers.TestSectionBuilder
	languageApplicationBuilder parsers.LanguageApplicationBuilder
	languageDefinitionBuilder  parsers.LanguageDefinitionBuilder
	languageValueBuilder       parsers.LanguageValueBuilder
	patternMatchBuilder        parsers.PatternMatchBuilder
	patternLabelsBuilder       parsers.PatternLabelsBuilder
	scriptBuilder              parsers.ScriptBuilder
	scriptValueBuilder         parsers.ScriptValueBuilder
	headSectionBuilder         parsers.HeadSectionBuilder
	headValueBuilder           parsers.HeadValueBuilder
	loadSingleBuilder          parsers.LoadSingleBuilder
	testDeclarationBuilder     parsers.TestDeclarationBuilder
	labelDeclarationBuilder    parsers.LabelDeclarationBuilder
	stackFrame                 stackframes.StackFrame
	linker                     linkers.Linker
}

func createBuilder(
	instructionAdapterBuilder InstructionAdapterBuilder,
	stackFrameBuilder stackframes.Builder,
	programBuilder parsers.ProgramBuilder,
	testableBuilder parsers.TestableBuilder,
	executableBuilder parsers.ExecutableBuilder,
	applicationBuilder parsers.ApplicationBuilder,
	labelSectionBuilder parsers.LabelSectionBuilder,
	mainSectionBuilder parsers.MainSectionBuilder,
	testSectionBuilder parsers.TestSectionBuilder,
	languageApplicationBuilder parsers.LanguageApplicationBuilder,
	languageDefinitionBuilder parsers.LanguageDefinitionBuilder,
	languageValueBuilder parsers.LanguageValueBuilder,
	patternMatchBuilder parsers.PatternMatchBuilder,
	patternLabelsBuilder parsers.PatternLabelsBuilder,
	scriptBuilder parsers.ScriptBuilder,
	scriptValueBuilder parsers.ScriptValueBuilder,
	headSectionBuilder parsers.HeadSectionBuilder,
	headValueBuilder parsers.HeadValueBuilder,
	loadSingleBuilder parsers.LoadSingleBuilder,
	testDeclarationBuilder parsers.TestDeclarationBuilder,
	labelDeclarationBuilder parsers.LabelDeclarationBuilder,
) Builder {
	out := builder{
		instructionAdapterBuilder:  instructionAdapterBuilder,
		stackFrameBuilder:          stackFrameBuilder,
		programBuilder:             programBuilder,
		testableBuilder:            testableBuilder,
		executableBuilder:          executableBuilder,
		applicationBuilder:         applicationBuilder,
		labelSectionBuilder:        labelSectionBuilder,
		mainSectionBuilder:         mainSectionBuilder,
		testSectionBuilder:         testSectionBuilder,
		languageApplicationBuilder: languageApplicationBuilder,
		languageDefinitionBuilder:  languageDefinitionBuilder,
		languageValueBuilder:       languageValueBuilder,
		patternMatchBuilder:        patternMatchBuilder,
		patternLabelsBuilder:       patternLabelsBuilder,
		scriptBuilder:              scriptBuilder,
		scriptValueBuilder:         scriptValueBuilder,
		headSectionBuilder:         headSectionBuilder,
		headValueBuilder:           headValueBuilder,
		loadSingleBuilder:          loadSingleBuilder,
		testDeclarationBuilder:     testDeclarationBuilder,
		labelDeclarationBuilder:    labelDeclarationBuilder,
		stackFrame:                 nil,
		linker:                     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.instructionAdapterBuilder,
		app.stackFrameBuilder,
		app.programBuilder,
		app.testableBuilder,
		app.executableBuilder,
		app.applicationBuilder,
		app.labelSectionBuilder,
		app.mainSectionBuilder,
		app.testSectionBuilder,
		app.languageApplicationBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.loadSingleBuilder,
		app.testDeclarationBuilder,
		app.labelDeclarationBuilder,
	)
}

// WithLinker adds a linker to the builder
func (app *builder) WithLinker(linker linkers.Linker) Builder {
	app.linker = linker
	return app
}

// WithStackFrame adds a stackFrameto the builder
func (app *builder) WithStackFrame(stackFrame stackframes.StackFrame) Builder {
	app.stackFrame = stackFrame
	return app
}

// Now builds a new Composer instance
func (app *builder) Now() (Composer, error) {
	if app.linker == nil {
		return nil, errors.New("the linker is mandatory in order to build a Composer instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a Composer instance")
	}

	return createComposer(
		app.instructionAdapterBuilder,
		app.stackFrameBuilder,
		app.programBuilder,
		app.testableBuilder,
		app.executableBuilder,
		app.applicationBuilder,
		app.labelSectionBuilder,
		app.mainSectionBuilder,
		app.testSectionBuilder,
		app.languageApplicationBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.loadSingleBuilder,
		app.testDeclarationBuilder,
		app.labelDeclarationBuilder,
		app.stackFrame,
		app.linker,
	), nil
}
