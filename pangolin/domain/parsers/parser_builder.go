package parsers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
)

type parserBuilder struct {
	lexerBuilder                    lexers.Builder
	lexerParserApplication          lexer_parser.Application
	lexerParserBuilder              lexer_parser.Builder
	programBuilder                  ProgramBuilder
	languageBuilder                 LanguageBuilder
	scopesBuilder                   ScopesBuilder
	scopeBuilder                    ScopeBuilder
	commandBuilder                  CommandBuilder
	languageCommandBuilder          LanguageCommandBuilder
	scriptCommandBuilder            ScriptCommandBuilder
	headCommandBuilder              HeadCommandBuilder
	mainCommandBuilder              MainCommandBuilder
	mainCommandInstructionBuilder   MainCommandInstructionBuilder
	testCommandBuilder              TestCommandBuilder
	testCommandInstructionBuilder   TestCommandInstructionBuilder
	labelCommandBuilder             LabelCommandBuilder
	labelCommandInstructionBuilder  LabelCommandInstructionBuilder
	languageApplicationBuilder      LanguageApplicationBuilder
	languageMainSectionBuilder      LanguageMainSectionBuilder
	languageTestSectionBuilder      LanguageTestSectionBuilder
	languageTestDeclarationBuilder  LanguageTestDeclarationBuilder
	languageTestInstructionBuilder  LanguageTestInstructionBuilder
	languageLabelSectionBuilder     LanguageLabelSectionBuilder
	languageLabelDeclarationBuilder LanguageLabelDeclarationBuilder
	languageLabelInstructionBuilder LanguageLabelInstructionBuilder
	languageInstructionBuilder      LanguageInstructionBuilder
	languageDefinitionBuilder       LanguageDefinitionBuilder
	languageValueBuilder            LanguageValueBuilder
	scriptBuilder                   ScriptBuilder
	scriptValueBuilder              ScriptValueBuilder
	patternMatchBuilder             PatternMatchBuilder
	patternLabelsBuilder            PatternLabelsBuilder
	relativePathsBuilder            RelativePathsBuilder
	relativePathBuilder             RelativePathBuilder
	folderSectionBuilder            FolderSectionBuilder
	folderNameBuilder               FolderNameBuilder
	applicationBuilder              ApplicationBuilder
	testSectionBuilder              TestSectionBuilder
	testDeclarationBuilder          TestDeclarationBuilder
	testInstructionBuilder          TestInstructionBuilder
	assertBuilder                   AssertBuilder
	readFileBuilder                 ReadFileBuilder
	headSectionBuilder              HeadSectionBuilder
	headValueBuilder                HeadValueBuilder
	importSingleBuilder             ImportSingleBuilder
	labelSectionBuilder             LabelSectionBuilder
	labelDeclarationBuilder         LabelDeclarationBuilder
	labelInstructionBuilder         LabelInstructionBuilder
	mainSectionBuilder              MainSectionBuilder
	instructionBuilder              InstructionBuilder
	specificTokenCodeBuilder        SpecificTokenCodeBuilder
	tokenSectionBuilder             TokenSectionBuilder
	codeMatchBuilder                CodeMatchBuilder
	tokenBuilder                    TokenBuilder
	variableBuilder                 VariableBuilder
	concatenationBuilder            ConcatenationBuilder
	declarationBuilder              DeclarationBuilder
	assignmentBuilder               AssignmentBuilder
	valueRepresentationBuilder      ValueRepresentationBuilder
	valueBuilder                    ValueBuilder
	numericValueBuilder             NumericValueBuilder
	typeBuilder                     TypeBuilder
	operationBuilder                OperationBuilder
	arythmeticBuilder               ArythmeticBuilder
	relationalBuilder               RelationalBuilder
	logicalBuilder                  LogicalBuilder
	transformOperationBuilder       TransformOperationBuilder
	standardOperationBuilder        StandardOperationBuilder
	remainingOperationBuilder       RemainingOperationBuilder
	printBuilder                    PrintBuilder
	jumpBuilder                     JumpBuilder
	matchBuilder                    MatchBuilder
	exitBuilder                     ExitBuilder
	callBuilder                     CallBuilder
	stackFrameBuilder               StackFrameBuilder
	indexBuilder                    IndexBuilder
	skipBuilder                     SkipBuilder
	intPointerBuilder               IntPointerBuilder
	lexerAdapter                    lexers.Adapter
}

func createParserBuilder(
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerBuilder lexers.Builder,
	programBuilder ProgramBuilder,
	languageBuilder LanguageBuilder,
	scopesBuilder ScopesBuilder,
	scopeBuilder ScopeBuilder,
	commandBuilder CommandBuilder,
	languageCommandBuilder LanguageCommandBuilder,
	scriptCommandBuilder ScriptCommandBuilder,
	headCommandBuilder HeadCommandBuilder,
	mainCommandBuilder MainCommandBuilder,
	mainCommandInstructionBuilder MainCommandInstructionBuilder,
	testCommandBuilder TestCommandBuilder,
	testCommandInstructionBuilder TestCommandInstructionBuilder,
	labelCommandBuilder LabelCommandBuilder,
	labelCommandInstructionBuilder LabelCommandInstructionBuilder,
	languageApplicationBuilder LanguageApplicationBuilder,
	languageMainSectionBuilder LanguageMainSectionBuilder,
	languageTestSectionBuilder LanguageTestSectionBuilder,
	languageTestDeclarationBuilder LanguageTestDeclarationBuilder,
	languageTestInstructionBuilder LanguageTestInstructionBuilder,
	languageLabelSectionBuilder LanguageLabelSectionBuilder,
	languageLabelDeclarationBuilder LanguageLabelDeclarationBuilder,
	languageLabelInstructionBuilder LanguageLabelInstructionBuilder,
	languageInstructionBuilder LanguageInstructionBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	languageValueBuilder LanguageValueBuilder,
	scriptBuilder ScriptBuilder,
	scriptValueBuilder ScriptValueBuilder,
	patternMatchBuilder PatternMatchBuilder,
	patternLabelsBuilder PatternLabelsBuilder,
	relativePathsBuilder RelativePathsBuilder,
	relativePathBuilder RelativePathBuilder,
	folderSectionBuilder FolderSectionBuilder,
	folderNameBuilder FolderNameBuilder,
	applicationBuilder ApplicationBuilder,
	testSectionBuilder TestSectionBuilder,
	testDeclarationBuilder TestDeclarationBuilder,
	testInstructionBuilder TestInstructionBuilder,
	assertBuilder AssertBuilder,
	readFileBuilder ReadFileBuilder,
	headSectionBuilder HeadSectionBuilder,
	headValueBuilder HeadValueBuilder,
	importSingleBuilder ImportSingleBuilder,
	labelSectionBuilder LabelSectionBuilder,
	labelDeclarationBuilder LabelDeclarationBuilder,
	labelInstructionBuilder LabelInstructionBuilder,
	mainSectionBuilder MainSectionBuilder,
	instructionBuilder InstructionBuilder,
	specificTokenCodeBuilder SpecificTokenCodeBuilder,
	tokenSectionBuilder TokenSectionBuilder,
	codeMatchBuilder CodeMatchBuilder,
	tokenBuilder TokenBuilder,
	variableBuilder VariableBuilder,
	concatenationBuilder ConcatenationBuilder,
	declarationBuilder DeclarationBuilder,
	assignmentBuilder AssignmentBuilder,
	valueRepresentationBuilder ValueRepresentationBuilder,
	valueBuilder ValueBuilder,
	numericValueBuilder NumericValueBuilder,
	typeBuilder TypeBuilder,
	operationBuilder OperationBuilder,
	arythmeticBuilder ArythmeticBuilder,
	relationalBuilder RelationalBuilder,
	logicalBuilder LogicalBuilder,
	transformOperationBuilder TransformOperationBuilder,
	standardOperationBuilder StandardOperationBuilder,
	remainingOperationBuilder RemainingOperationBuilder,
	printBuilder PrintBuilder,
	jumpBuilder JumpBuilder,
	matchBuilder MatchBuilder,
	exitBuilder ExitBuilder,
	callBuilder CallBuilder,
	stackFrameBuilder StackFrameBuilder,
	indexBuilder IndexBuilder,
	skipBuilder SkipBuilder,
	intPointerBuilder IntPointerBuilder,
) ParserBuilder {
	out := parserBuilder{
		lexerParserApplication:          lexerParserApplication,
		lexerParserBuilder:              lexerParserBuilder,
		lexerBuilder:                    lexerBuilder,
		programBuilder:                  programBuilder,
		languageBuilder:                 languageBuilder,
		scopesBuilder:                   scopesBuilder,
		scopeBuilder:                    scopeBuilder,
		commandBuilder:                  commandBuilder,
		languageCommandBuilder:          languageCommandBuilder,
		scriptCommandBuilder:            scriptCommandBuilder,
		headCommandBuilder:              headCommandBuilder,
		mainCommandBuilder:              mainCommandBuilder,
		mainCommandInstructionBuilder:   mainCommandInstructionBuilder,
		testCommandBuilder:              testCommandBuilder,
		testCommandInstructionBuilder:   testCommandInstructionBuilder,
		labelCommandBuilder:             labelCommandBuilder,
		labelCommandInstructionBuilder:  labelCommandInstructionBuilder,
		languageApplicationBuilder:      languageApplicationBuilder,
		languageMainSectionBuilder:      languageMainSectionBuilder,
		languageTestSectionBuilder:      languageTestSectionBuilder,
		languageTestDeclarationBuilder:  languageTestDeclarationBuilder,
		languageTestInstructionBuilder:  languageTestInstructionBuilder,
		languageLabelSectionBuilder:     languageLabelSectionBuilder,
		languageLabelDeclarationBuilder: languageLabelDeclarationBuilder,
		languageLabelInstructionBuilder: languageLabelInstructionBuilder,
		languageInstructionBuilder:      languageInstructionBuilder,
		languageDefinitionBuilder:       languageDefinitionBuilder,
		languageValueBuilder:            languageValueBuilder,
		scriptBuilder:                   scriptBuilder,
		scriptValueBuilder:              scriptValueBuilder,
		patternMatchBuilder:             patternMatchBuilder,
		patternLabelsBuilder:            patternLabelsBuilder,
		relativePathsBuilder:            relativePathsBuilder,
		relativePathBuilder:             relativePathBuilder,
		folderSectionBuilder:            folderSectionBuilder,
		folderNameBuilder:               folderNameBuilder,
		applicationBuilder:              applicationBuilder,
		testSectionBuilder:              testSectionBuilder,
		testDeclarationBuilder:          testDeclarationBuilder,
		testInstructionBuilder:          testInstructionBuilder,
		assertBuilder:                   assertBuilder,
		readFileBuilder:                 readFileBuilder,
		headSectionBuilder:              headSectionBuilder,
		headValueBuilder:                headValueBuilder,
		importSingleBuilder:             importSingleBuilder,
		labelSectionBuilder:             labelSectionBuilder,
		labelDeclarationBuilder:         labelDeclarationBuilder,
		labelInstructionBuilder:         labelInstructionBuilder,
		mainSectionBuilder:              mainSectionBuilder,
		instructionBuilder:              instructionBuilder,
		specificTokenCodeBuilder:        specificTokenCodeBuilder,
		tokenSectionBuilder:             tokenSectionBuilder,
		codeMatchBuilder:                codeMatchBuilder,
		tokenBuilder:                    tokenBuilder,
		variableBuilder:                 variableBuilder,
		concatenationBuilder:            concatenationBuilder,
		declarationBuilder:              declarationBuilder,
		assignmentBuilder:               assignmentBuilder,
		valueRepresentationBuilder:      valueRepresentationBuilder,
		valueBuilder:                    valueBuilder,
		numericValueBuilder:             numericValueBuilder,
		typeBuilder:                     typeBuilder,
		operationBuilder:                operationBuilder,
		arythmeticBuilder:               arythmeticBuilder,
		relationalBuilder:               relationalBuilder,
		logicalBuilder:                  logicalBuilder,
		transformOperationBuilder:       transformOperationBuilder,
		standardOperationBuilder:        standardOperationBuilder,
		remainingOperationBuilder:       remainingOperationBuilder,
		printBuilder:                    printBuilder,
		jumpBuilder:                     jumpBuilder,
		matchBuilder:                    matchBuilder,
		exitBuilder:                     exitBuilder,
		callBuilder:                     callBuilder,
		stackFrameBuilder:               stackFrameBuilder,
		indexBuilder:                    indexBuilder,
		skipBuilder:                     skipBuilder,
		intPointerBuilder:               intPointerBuilder,
		lexerAdapter:                    nil,
	}

	return &out
}

// Create initializes the builder
func (app *parserBuilder) Create() ParserBuilder {
	return createParserBuilder(
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerBuilder,
		app.programBuilder,
		app.languageBuilder,
		app.scopesBuilder,
		app.scopeBuilder,
		app.commandBuilder,
		app.languageCommandBuilder,
		app.scriptCommandBuilder,
		app.headCommandBuilder,
		app.mainCommandBuilder,
		app.mainCommandInstructionBuilder,
		app.testCommandBuilder,
		app.testCommandInstructionBuilder,
		app.labelCommandBuilder,
		app.labelCommandInstructionBuilder,
		app.languageApplicationBuilder,
		app.languageMainSectionBuilder,
		app.languageTestSectionBuilder,
		app.languageTestDeclarationBuilder,
		app.languageTestInstructionBuilder,
		app.languageLabelSectionBuilder,
		app.languageLabelDeclarationBuilder,
		app.languageLabelInstructionBuilder,
		app.languageInstructionBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.relativePathsBuilder,
		app.relativePathBuilder,
		app.folderSectionBuilder,
		app.folderNameBuilder,
		app.applicationBuilder,
		app.testSectionBuilder,
		app.testDeclarationBuilder,
		app.testInstructionBuilder,
		app.assertBuilder,
		app.readFileBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.importSingleBuilder,
		app.labelSectionBuilder,
		app.labelDeclarationBuilder,
		app.labelInstructionBuilder,
		app.mainSectionBuilder,
		app.instructionBuilder,
		app.specificTokenCodeBuilder,
		app.tokenSectionBuilder,
		app.codeMatchBuilder,
		app.tokenBuilder,
		app.variableBuilder,
		app.concatenationBuilder,
		app.declarationBuilder,
		app.assignmentBuilder,
		app.valueRepresentationBuilder,
		app.valueBuilder,
		app.numericValueBuilder,
		app.typeBuilder,
		app.operationBuilder,
		app.arythmeticBuilder,
		app.relationalBuilder,
		app.logicalBuilder,
		app.transformOperationBuilder,
		app.standardOperationBuilder,
		app.remainingOperationBuilder,
		app.printBuilder,
		app.jumpBuilder,
		app.matchBuilder,
		app.exitBuilder,
		app.callBuilder,
		app.stackFrameBuilder,
		app.indexBuilder,
		app.skipBuilder,
		app.intPointerBuilder,
	)
}

// WithLexerAdapter adds a lexerAdapter to the builder
func (app *parserBuilder) WithLexerAdapter(lexerAdapter lexers.Adapter) ParserBuilder {
	app.lexerAdapter = lexerAdapter
	return app
}

// Now builds a new Parser instance
func (app *parserBuilder) Now() (Parser, error) {
	if app.lexerAdapter == nil {
		return nil, errors.New("the lexerAdapter is mandatory in order to build a Parser instance")
	}

	return createParser(
		app.lexerAdapter,
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerBuilder,
		app.programBuilder,
		app.languageBuilder,
		app.scopesBuilder,
		app.scopeBuilder,
		app.commandBuilder,
		app.languageCommandBuilder,
		app.scriptCommandBuilder,
		app.headCommandBuilder,
		app.mainCommandBuilder,
		app.mainCommandInstructionBuilder,
		app.testCommandBuilder,
		app.testCommandInstructionBuilder,
		app.labelCommandBuilder,
		app.labelCommandInstructionBuilder,
		app.languageApplicationBuilder,
		app.languageMainSectionBuilder,
		app.languageTestSectionBuilder,
		app.languageTestDeclarationBuilder,
		app.languageTestInstructionBuilder,
		app.languageLabelSectionBuilder,
		app.languageLabelDeclarationBuilder,
		app.languageLabelInstructionBuilder,
		app.languageInstructionBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.relativePathsBuilder,
		app.relativePathBuilder,
		app.folderSectionBuilder,
		app.folderNameBuilder,
		app.applicationBuilder,
		app.testSectionBuilder,
		app.testDeclarationBuilder,
		app.testInstructionBuilder,
		app.assertBuilder,
		app.readFileBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.importSingleBuilder,
		app.labelSectionBuilder,
		app.labelDeclarationBuilder,
		app.labelInstructionBuilder,
		app.mainSectionBuilder,
		app.instructionBuilder,
		app.specificTokenCodeBuilder,
		app.tokenSectionBuilder,
		app.codeMatchBuilder,
		app.tokenBuilder,
		app.variableBuilder,
		app.concatenationBuilder,
		app.declarationBuilder,
		app.assignmentBuilder,
		app.valueRepresentationBuilder,
		app.valueBuilder,
		app.numericValueBuilder,
		app.typeBuilder,
		app.operationBuilder,
		app.arythmeticBuilder,
		app.relationalBuilder,
		app.logicalBuilder,
		app.transformOperationBuilder,
		app.standardOperationBuilder,
		app.remainingOperationBuilder,
		app.printBuilder,
		app.jumpBuilder,
		app.matchBuilder,
		app.exitBuilder,
		app.callBuilder,
		app.stackFrameBuilder,
		app.indexBuilder,
		app.skipBuilder,
		app.intPointerBuilder,
	)
}
