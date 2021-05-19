package composers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type composerBuilder struct {
	linker                     linkers.Linker
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
}

func createComposerBuilder(
	linker linkers.Linker,
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
) ComposerBuilder {
	out := composerBuilder{
		linker:                     linker,
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
		stackFrame:                 nil,
	}

	return &out
}

// Create initializes the builder
func (app *composerBuilder) Create() ComposerBuilder {
	return createComposerBuilder(
		app.linker,
		app.instructionAdapterBuilder,
		app.stackFrameBuilder,
		app.programBuilder,
		app.applicationBuilder,
		app.labelSectionBuilder,
		app.mainSectionBuilder,
		app.testSectionBuilder,
		app.languageBuilder,
		app.languageApplicationBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.testDeclarationBuilder,
		app.labelDeclarationBuilder,
	)
}

// WithStackFrame adds a stackFrameto the builder
func (app *composerBuilder) WithStackFrame(stackFrame stackframes.StackFrame) ComposerBuilder {
	app.stackFrame = stackFrame
	return app
}

// Now builds a new Composer instance
func (app *composerBuilder) Now() (Composer, error) {
	if app.stackFrame != nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a Composer instance")
	}

	return createComposer(
		app.linker,
		app.instructionAdapterBuilder,
		app.stackFrameBuilder,
		app.programBuilder,
		app.applicationBuilder,
		app.labelSectionBuilder,
		app.mainSectionBuilder,
		app.testSectionBuilder,
		app.languageBuilder,
		app.languageApplicationBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.testDeclarationBuilder,
		app.labelDeclarationBuilder,
		app.stackFrame,
	), nil
}
