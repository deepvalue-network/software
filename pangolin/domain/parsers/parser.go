package parsers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	lparser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
)

type parser struct {
	lexerAdapter                     lexers.Adapter
	lexerApplication                 lparser.Application
	lexerBuilder                     lexers.Builder
	parserBuilder                    lparser.Builder
	programBuilder                   ProgramBuilder
	testableBuilder                  TestableBuilder
	executableBuilder                ExecutableBuilder
	scopesBuilder                    ScopesBuilder
	scopeBuilder                     ScopeBuilder
	commandBuilder                   CommandBuilder
	languageCommandBuilder           LanguageCommandBuilder
	scriptCommandBuilder             ScriptCommandBuilder
	headCommandBuilder               HeadCommandBuilder
	mainCommandBuilder               MainCommandBuilder
	mainCommandInstructionBuilder    MainCommandInstructionBuilder
	testCommandBuilder               TestCommandBuilder
	testCommandInstructionBuilder    TestCommandInstructionBuilder
	labelCommandBuilder              LabelCommandBuilder
	labelCommandInstructionBuilder   LabelCommandInstructionBuilder
	languageApplicationBuilder       LanguageApplicationBuilder
	languageMainSectionBuilder       LanguageMainSectionBuilder
	languageTestSectionBuilder       LanguageTestSectionBuilder
	languageTestDeclarationBuilder   LanguageTestDeclarationBuilder
	languageTestInstructionBuilder   LanguageTestInstructionBuilder
	languageLabelSectionBuilder      LanguageLabelSectionBuilder
	languageLabelDeclarationBuilder  LanguageLabelDeclarationBuilder
	languageLabelInstructionBuilder  LanguageLabelInstructionBuilder
	languageInstructionBuilder       LanguageInstructionBuilder
	languageInstructionCommonBuilder LanguageInstructionCommonBuilder
	languageDefinitionBuilder        LanguageDefinitionBuilder
	languageValueBuilder             LanguageValueBuilder
	scriptBuilder                    ScriptBuilder
	scriptValueBuilder               ScriptValueBuilder
	scriptTestsBuilder               ScriptTestsBuilder
	scriptTestBuilder                ScriptTestBuilder
	patternMatchBuilder              PatternMatchBuilder
	patternLabelsBuilder             PatternLabelsBuilder
	relativePathsBuilder             RelativePathsBuilder
	relativePathBuilder              RelativePathBuilder
	folderSectionBuilder             FolderSectionBuilder
	folderNameBuilder                FolderNameBuilder
	applicationBuilder               ApplicationBuilder
	testSectionBuilder               TestSectionBuilder
	testDeclarationBuilder           TestDeclarationBuilder
	testInstructionBuilder           TestInstructionBuilder
	assertBuilder                    AssertBuilder
	readFileBuilder                  ReadFileBuilder
	headSectionBuilder               HeadSectionBuilder
	headValueBuilder                 HeadValueBuilder
	loadSingleBuilder                LoadSingleBuilder
	importSingleBuilder              ImportSingleBuilder
	labelSectionBuilder              LabelSectionBuilder
	labelDeclarationBuilder          LabelDeclarationBuilder
	labelInstructionBuilder          LabelInstructionBuilder
	mainSectionBuilder               MainSectionBuilder
	instructionBuilder               InstructionBuilder
	registryBuilder                  RegistryBuilder
	fetchRegistryBuilder             FetchRegistryBuilder
	unregisterBuilder                UnregisterBuilder
	registerBuilder                  RegisterBuilder
	specificTokenCodeBuilder         SpecificTokenCodeBuilder
	tokenSectionBuilder              TokenSectionBuilder
	codeMatchBuilder                 CodeMatchBuilder
	tokenBuilder                     TokenBuilder
	variableBuilder                  VariableBuilder
	concatenationBuilder             ConcatenationBuilder
	declarationBuilder               DeclarationBuilder
	assignmentBuilder                AssignmentBuilder
	valueRepresentationBuilder       ValueRepresentationBuilder
	valueBuilder                     ValueBuilder
	numericValueBuilder              NumericValueBuilder
	typeBuilder                      TypeBuilder
	operationBuilder                 OperationBuilder
	arythmeticBuilder                ArythmeticBuilder
	relationalBuilder                RelationalBuilder
	logicalBuilder                   LogicalBuilder
	standardOperationBuilder         StandardOperationBuilder
	remainingOperationBuilder        RemainingOperationBuilder
	printBuilder                     PrintBuilder
	jumpBuilder                      JumpBuilder
	matchBuilder                     MatchBuilder
	exitBuilder                      ExitBuilder
	callBuilder                      CallBuilder
	moduleBuilder                    ModuleBuilder
	switchBuilder                    SwitchBuilder
	saveBuilder                      SaveBuilder
	stackFrameBuilder                StackFrameBuilder
	indexBuilder                     IndexBuilder
	skipBuilder                      SkipBuilder
	intPointerBuilder                IntPointerBuilder
	program                          map[string]Program
	testable                         map[string]Testable
	executable                       map[string]Executable
	scopes                           map[string]Scopes
	scope                            map[string]Scope
	command                          map[string]Command
	languageCommand                  map[string]LanguageCommand
	scriptCommand                    map[string]ScriptCommand
	headCommand                      map[string]HeadCommand
	mainCommand                      map[string]MainCommand
	mainCommandInstruction           map[string]MainCommandInstruction
	testCommand                      map[string]TestCommand
	testCommandInstruction           map[string]TestCommandInstruction
	labelCommand                     map[string]LabelCommand
	labelCommandInstruction          map[string]LabelCommandInstruction
	languageApplication              map[string]LanguageApplication
	languageMainSection              map[string]LanguageMainSection
	languageTestSection              map[string]LanguageTestSection
	languageTestDeclaration          map[string]LanguageTestDeclaration
	languageTestInstruction          map[string]LanguageTestInstruction
	languageLabelSection             map[string]LanguageLabelSection
	languageLabelDeclaration         map[string]LanguageLabelDeclaration
	languageLabelInstruction         map[string]LanguageLabelInstruction
	languageInstruction              map[string]LanguageInstruction
	languageInstructionCommon        map[string]LanguageInstructionCommon
	languageDefinition               map[string]LanguageDefinition
	languageValue                    map[string]LanguageValue
	targetPath                       map[string]RelativePath
	script                           map[string]Script
	scriptValue                      map[string]ScriptValue
	scriptTests                      map[string]ScriptTests
	scriptTest                       map[string]ScriptTest
	patternMatch                     map[string]PatternMatch
	patternLabels                    map[string]PatternLabels
	patternLabelEnter                map[string]string
	patternLabelExit                 map[string]string
	relativePaths                    map[string]RelativePaths
	relativePath                     map[string]RelativePath
	folderSection                    map[string]FolderSection
	folderName                       map[string]FolderName
	application                      map[string]Application
	testSection                      map[string]TestSection
	testDeclaration                  map[string]TestDeclaration
	testInstruction                  map[string]TestInstruction
	assert                           map[string]Assert
	readFile                         map[string]ReadFile
	headSection                      map[string]HeadSection
	headValue                        map[string]HeadValue
	loadSingle                       map[string]LoadSingle
	importSingle                     map[string]ImportSingle
	labelSection                     map[string]LabelSection
	labelDeclaration                 map[string]LabelDeclaration
	labelInstruction                 map[string]LabelInstruction
	mainSection                      map[string]MainSection
	instruction                      map[string]Instruction
	registry                         map[string]Registry
	fetchRegistry                    map[string]FetchRegistry
	unregister                       map[string]Unregister
	register                         map[string]Register
	specificTokenCode                map[string]SpecificTokenCode
	tokenSection                     map[string]TokenSection
	codeMatch                        map[string]CodeMatch
	token                            map[string]Token
	callPattern                      map[string]string
	callPatterns                     map[string][]string
	patternOrRule                    map[string]string
	importNames                      map[string]string
	extendNames                      map[string]string
	variable                         map[string]Variable
	concatenation                    map[string]Concatenation
	declaration                      map[string]Declaration
	assignment                       map[string]Assignment
	valueRepresentation              map[string]ValueRepresentation
	value                            map[string]Value
	numericValue                     map[string]NumericValue
	boolValue                        map[string]bool
	floatValue                       map[string]float64
	stringValue                      map[string]string
	typ                              map[string]Type
	operation                        map[string]Operation
	arythmetic                       map[string]Arythmetic
	relational                       map[string]Relational
	logical                          map[string]Logical
	standardOperation                map[string]StandardOperation
	remainingOperation               map[string]RemainingOperation
	print                            map[string]Print
	jump                             map[string]Jump
	match                            map[string]Match
	matchPattern                     map[string]string
	exit                             map[string]Exit
	call                             map[string]Call
	module                           map[string]Module
	swtch                            map[string]Switch
	save                             map[string]Save
	stackFrame                       map[string]StackFrame
	index                            map[string]Index
	skip                             map[string]Skip
	intPointer                       map[string]IntPointer
}

func createParser(
	lexerAdapter lexers.Adapter,
	lexerApplication lparser.Application,
	parserBuilder lparser.Builder,
	lexerBuilder lexers.Builder,
	programBuilder ProgramBuilder,
	testableBuilder TestableBuilder,
	executableBuilder ExecutableBuilder,
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
	languageInstructionCommonBuilder LanguageInstructionCommonBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	languageValueBuilder LanguageValueBuilder,
	scriptBuilder ScriptBuilder,
	scriptValueBuilder ScriptValueBuilder,
	scriptTestsBuilder ScriptTestsBuilder,
	scriptTestBuilder ScriptTestBuilder,
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
	loadSingleBuilder LoadSingleBuilder,
	importSingleBuilder ImportSingleBuilder,
	labelSectionBuilder LabelSectionBuilder,
	labelDeclarationBuilder LabelDeclarationBuilder,
	labelInstructionBuilder LabelInstructionBuilder,
	mainSectionBuilder MainSectionBuilder,
	instructionBuilder InstructionBuilder,
	registryBuilder RegistryBuilder,
	fetchRegistryBuilder FetchRegistryBuilder,
	unregisterBuilder UnregisterBuilder,
	registerBuilder RegisterBuilder,
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
	standardOperationBuilder StandardOperationBuilder,
	remainingOperationBuilder RemainingOperationBuilder,
	printBuilder PrintBuilder,
	jumpBuilder JumpBuilder,
	matchBuilder MatchBuilder,
	exitBuilder ExitBuilder,
	callBuilder CallBuilder,
	moduleBuilder ModuleBuilder,
	switchBuilder SwitchBuilder,
	saveBuilder SaveBuilder,
	stackFrameBuilder StackFrameBuilder,
	indexBuilder IndexBuilder,
	skipBuilder SkipBuilder,
	intPointerBuilder IntPointerBuilder,
) (*parser, error) {
	out := &parser{
		lexerApplication:                 lexerApplication,
		parserBuilder:                    parserBuilder,
		lexerBuilder:                     lexerBuilder,
		lexerAdapter:                     lexerAdapter,
		programBuilder:                   programBuilder,
		testableBuilder:                  testableBuilder,
		executableBuilder:                executableBuilder,
		scopesBuilder:                    scopesBuilder,
		scopeBuilder:                     scopeBuilder,
		commandBuilder:                   commandBuilder,
		languageCommandBuilder:           languageCommandBuilder,
		scriptCommandBuilder:             scriptCommandBuilder,
		headCommandBuilder:               headCommandBuilder,
		mainCommandBuilder:               mainCommandBuilder,
		mainCommandInstructionBuilder:    mainCommandInstructionBuilder,
		testCommandBuilder:               testCommandBuilder,
		testCommandInstructionBuilder:    testCommandInstructionBuilder,
		labelCommandBuilder:              labelCommandBuilder,
		labelCommandInstructionBuilder:   labelCommandInstructionBuilder,
		languageApplicationBuilder:       languageApplicationBuilder,
		languageMainSectionBuilder:       languageMainSectionBuilder,
		languageTestSectionBuilder:       languageTestSectionBuilder,
		languageTestDeclarationBuilder:   languageTestDeclarationBuilder,
		languageTestInstructionBuilder:   languageTestInstructionBuilder,
		languageLabelSectionBuilder:      languageLabelSectionBuilder,
		languageLabelDeclarationBuilder:  languageLabelDeclarationBuilder,
		languageLabelInstructionBuilder:  languageLabelInstructionBuilder,
		languageInstructionBuilder:       languageInstructionBuilder,
		languageInstructionCommonBuilder: languageInstructionCommonBuilder,
		languageDefinitionBuilder:        languageDefinitionBuilder,
		languageValueBuilder:             languageValueBuilder,
		scriptBuilder:                    scriptBuilder,
		scriptValueBuilder:               scriptValueBuilder,
		scriptTestsBuilder:               scriptTestsBuilder,
		scriptTestBuilder:                scriptTestBuilder,
		patternMatchBuilder:              patternMatchBuilder,
		patternLabelsBuilder:             patternLabelsBuilder,
		relativePathsBuilder:             relativePathsBuilder,
		relativePathBuilder:              relativePathBuilder,
		folderSectionBuilder:             folderSectionBuilder,
		folderNameBuilder:                folderNameBuilder,
		applicationBuilder:               applicationBuilder,
		testSectionBuilder:               testSectionBuilder,
		testDeclarationBuilder:           testDeclarationBuilder,
		testInstructionBuilder:           testInstructionBuilder,
		assertBuilder:                    assertBuilder,
		readFileBuilder:                  readFileBuilder,
		headSectionBuilder:               headSectionBuilder,
		headValueBuilder:                 headValueBuilder,
		loadSingleBuilder:                loadSingleBuilder,
		importSingleBuilder:              importSingleBuilder,
		labelSectionBuilder:              labelSectionBuilder,
		labelDeclarationBuilder:          labelDeclarationBuilder,
		labelInstructionBuilder:          labelInstructionBuilder,
		mainSectionBuilder:               mainSectionBuilder,
		instructionBuilder:               instructionBuilder,
		registryBuilder:                  registryBuilder,
		fetchRegistryBuilder:             fetchRegistryBuilder,
		unregisterBuilder:                unregisterBuilder,
		registerBuilder:                  registerBuilder,
		specificTokenCodeBuilder:         specificTokenCodeBuilder,
		tokenSectionBuilder:              tokenSectionBuilder,
		codeMatchBuilder:                 codeMatchBuilder,
		tokenBuilder:                     tokenBuilder,
		variableBuilder:                  variableBuilder,
		concatenationBuilder:             concatenationBuilder,
		declarationBuilder:               declarationBuilder,
		assignmentBuilder:                assignmentBuilder,
		valueRepresentationBuilder:       valueRepresentationBuilder,
		valueBuilder:                     valueBuilder,
		numericValueBuilder:              numericValueBuilder,
		typeBuilder:                      typeBuilder,
		operationBuilder:                 operationBuilder,
		arythmeticBuilder:                arythmeticBuilder,
		relationalBuilder:                relationalBuilder,
		logicalBuilder:                   logicalBuilder,
		standardOperationBuilder:         standardOperationBuilder,
		remainingOperationBuilder:        remainingOperationBuilder,
		printBuilder:                     printBuilder,
		jumpBuilder:                      jumpBuilder,
		matchBuilder:                     matchBuilder,
		exitBuilder:                      exitBuilder,
		callBuilder:                      callBuilder,
		moduleBuilder:                    moduleBuilder,
		switchBuilder:                    switchBuilder,
		saveBuilder:                      saveBuilder,
		stackFrameBuilder:                stackFrameBuilder,
		indexBuilder:                     indexBuilder,
		skipBuilder:                      skipBuilder,
		intPointerBuilder:                intPointerBuilder,
	}

	out.init()
	return out, nil
}

// Execute executes the parser from a lexer
func (app *parser) Execute(lexer lexers.Lexer) (interface{}, error) {
	params := []lparser.ToEventsParams{
		lparser.ToEventsParams{
			Token:  "program",
			OnExit: app.exitProgram,
		},
		lparser.ToEventsParams{
			Token:  "testable",
			OnExit: app.exitTestable,
		},
		lparser.ToEventsParams{
			Token:  "executable",
			OnExit: app.exitExecutable,
		},
		lparser.ToEventsParams{
			Token:  "scopesWithArrow",
			OnExit: app.exitScopesWithArrow,
		},
		lparser.ToEventsParams{
			Token:  "scopes",
			OnExit: app.exitScopes,
		},
		lparser.ToEventsParams{
			Token:  "scope",
			OnExit: app.exitScope,
		},
		lparser.ToEventsParams{
			Token:  "command",
			OnExit: app.exitCommand,
		},
		lparser.ToEventsParams{
			Token:  "languageCommand",
			OnExit: app.exitLanguageCommand,
		},
		lparser.ToEventsParams{
			Token:  "scriptCommand",
			OnExit: app.exitScriptCommand,
		},
		lparser.ToEventsParams{
			Token:  "headCommand",
			OnExit: app.exitHeadCommand,
		},
		lparser.ToEventsParams{
			Token:  "mainCommand",
			OnExit: app.exitMainCommand,
		},
		lparser.ToEventsParams{
			Token:  "mainCommandInstruction",
			OnExit: app.exitMainCommandInstruction,
		},
		lparser.ToEventsParams{
			Token:  "testCommand",
			OnExit: app.exitTestCommand,
		},
		lparser.ToEventsParams{
			Token:  "testCommandInstruction",
			OnExit: app.exitTestCommandInstruction,
		},
		lparser.ToEventsParams{
			Token:  "labelCommand",
			OnExit: app.exitLabelCommand,
		},
		lparser.ToEventsParams{
			Token:  "labelCommandInstruction",
			OnExit: app.exitLabelCommandInstruction,
		},
		lparser.ToEventsParams{
			Token:  "languageApplication",
			OnExit: app.exitLanguageApplication,
		},
		lparser.ToEventsParams{
			Token:  "languageMainSection",
			OnExit: app.exitLanguageMainSection,
		},
		lparser.ToEventsParams{
			Token:  "languageTestSection",
			OnExit: app.exitLanguageTestSection,
		},
		lparser.ToEventsParams{
			Token:  "languageTestDeclaration",
			OnExit: app.exitLanguageTestDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "languageTestInstruction",
			OnExit: app.exitLanguageTestInstruction,
		},
		lparser.ToEventsParams{
			Token:  "languageLabelSection",
			OnExit: app.exitLanguageLabelSection,
		},
		lparser.ToEventsParams{
			Token:  "languageLabelDeclaration",
			OnExit: app.exitLanguageLabelDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "languageLabelInstruction",
			OnExit: app.exitLanguageLabelInstruction,
		},
		lparser.ToEventsParams{
			Token:  "languageInstruction",
			OnExit: app.exitLanguageInstruction,
		},
		lparser.ToEventsParams{
			Token:  "languageInstructionCommon",
			OnExit: app.exitLanguageInstructionCommon,
		},
		lparser.ToEventsParams{
			Token:  "languageDefinition",
			OnExit: app.exitLanguageDefinition,
		},
		lparser.ToEventsParams{
			Token:  "languageValue",
			OnExit: app.exitLanguageValue,
		},
		lparser.ToEventsParams{
			Token:  "script",
			OnExit: app.exitScript,
		},
		lparser.ToEventsParams{
			Token:  "scriptValue",
			OnExit: app.exitScriptValue,
		},
		lparser.ToEventsParams{
			Token:  "scriptTests",
			OnExit: app.exitScriptTests,
		},
		lparser.ToEventsParams{
			Token:  "scriptTestWithComma",
			OnExit: app.exitScriptTestWithComma,
		},
		lparser.ToEventsParams{
			Token:  "scriptTest",
			OnExit: app.exitScriptTest,
		},
		lparser.ToEventsParams{
			Token:  "patternMatch",
			OnExit: app.exitPatternMatch,
		},
		lparser.ToEventsParams{
			Token:  "patternLabels",
			OnExit: app.exitPatternLabels,
		},
		lparser.ToEventsParams{
			Token:  "patternLabelEnter",
			OnExit: app.exitPatternLabelEnter,
		},
		lparser.ToEventsParams{
			Token:  "patternLabelExit",
			OnExit: app.exitPatternLabelExit,
		},
		lparser.ToEventsParams{
			Token:  "singleExtend",
			OnExit: app.exitSingleExtend,
		},
		lparser.ToEventsParams{
			Token:  "relativePaths",
			OnExit: app.exitRelativePaths,
		},
		lparser.ToEventsParams{
			Token:  "relativePathWithComma",
			OnExit: app.exitRelativePathWithComma,
		},
		lparser.ToEventsParams{
			Token:  "relativePath",
			OnExit: app.exitRelativePath,
		},
		lparser.ToEventsParams{
			Token:  "folderSection",
			OnExit: app.exitFolderSection,
		},
		lparser.ToEventsParams{
			Token:  "folderName",
			OnExit: app.exitFolderName,
		},
		lparser.ToEventsParams{
			Token:  "application",
			OnExit: app.exitApplication,
		},
		lparser.ToEventsParams{
			Token:  "testSection",
			OnExit: app.exitTestSection,
		},
		lparser.ToEventsParams{
			Token:  "testDeclaration",
			OnExit: app.exitTestDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "testInstruction",
			OnExit: app.exitTestInstruction,
		},
		lparser.ToEventsParams{
			Token:  "assert",
			OnExit: app.exitAssert,
		},
		lparser.ToEventsParams{
			Token:  "readFile",
			OnExit: app.exitReadFile,
		},
		lparser.ToEventsParams{
			Token:  "headSection",
			OnExit: app.exitHeadSection,
		},
		lparser.ToEventsParams{
			Token:  "headValue",
			OnExit: app.exitHeadValue,
		},
		lparser.ToEventsParams{
			Token:  "loadSingle",
			OnExit: app.exitLoadSingle,
		},
		lparser.ToEventsParams{
			Token:  "importSingle",
			OnExit: app.exitImportSingle,
		},
		lparser.ToEventsParams{
			Token:  "labelSection",
			OnExit: app.exitLabelSection,
		},
		lparser.ToEventsParams{
			Token:  "labelDeclaration",
			OnExit: app.exitLabelDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "labelInstruction",
			OnExit: app.exitLabelInstruction,
		},
		lparser.ToEventsParams{
			Token:  "mainSection",
			OnExit: app.exitMainSection,
		},
		lparser.ToEventsParams{
			Token:  "instruction",
			OnExit: app.exitInstruction,
		},
		lparser.ToEventsParams{
			Token:  "registry",
			OnExit: app.exitRegistry,
		},
		lparser.ToEventsParams{
			Token:  "fetchRegistry",
			OnExit: app.exitFetchRegistry,
		},
		lparser.ToEventsParams{
			Token:  "unregister",
			OnExit: app.exitUnregister,
		},
		lparser.ToEventsParams{
			Token:  "register",
			OnExit: app.exitRegister,
		},
		lparser.ToEventsParams{
			Token:  "callPattern",
			OnExit: app.exitCallPattern,
		},
		lparser.ToEventsParams{
			Token:  "pipeCallPattern",
			OnExit: app.exitPipeCallPattern,
		},
		lparser.ToEventsParams{
			Token:  "callPatterns",
			OnExit: app.exitCallPatterns,
		},
		lparser.ToEventsParams{
			Token:  "patternOrRule",
			OnExit: app.exitPatternOrRule,
		},
		lparser.ToEventsParams{
			Token:  "specificTokenCode",
			OnExit: app.exitSpecificTokenCode,
		},
		lparser.ToEventsParams{
			Token:  "specificTokenCodeWithAmount",
			OnExit: app.exitSpecificTokenCodeWithAmount,
		},
		lparser.ToEventsParams{
			Token:  "tokenSection",
			OnExit: app.exitTokenSection,
		},
		lparser.ToEventsParams{
			Token:  "codeMatch",
			OnExit: app.exitCodeMatch,
		},
		lparser.ToEventsParams{
			Token:  "token",
			OnExit: app.exitToken,
		},
		lparser.ToEventsParams{
			Token:  "variable",
			OnExit: app.exitVariable,
		},
		lparser.ToEventsParams{
			Token:  "concatenation",
			OnExit: app.exitConcatenation,
		},
		lparser.ToEventsParams{
			Token:  "declaration",
			OnExit: app.exitDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "assignment",
			OnExit: app.exitAssignment,
		},
		lparser.ToEventsParams{
			Token:  "value",
			OnExit: app.exitValue,
		},
		lparser.ToEventsParams{
			Token:  "valueRepresentation",
			OnExit: app.exitValueRepresentation,
		},
		lparser.ToEventsParams{
			Token:  "numericValue",
			OnExit: app.exitNumericValue,
		},
		lparser.ToEventsParams{
			Token:  "boolValue",
			OnExit: app.exitBool,
		},
		lparser.ToEventsParams{
			Token:  "floatValue",
			OnExit: app.exitFloatValue,
		},
		lparser.ToEventsParams{
			Token:  "stringValue",
			OnExit: app.exitStringValue,
		},
		lparser.ToEventsParams{
			Token:  "type",
			OnExit: app.exitType,
		},
		lparser.ToEventsParams{
			Token:  "operation",
			OnExit: app.exitOperation,
		},
		lparser.ToEventsParams{
			Token:  "arythmetic",
			OnExit: app.exitArythmetic,
		},
		lparser.ToEventsParams{
			Token:  "relational",
			OnExit: app.exitRelational,
		},
		lparser.ToEventsParams{
			Token:  "logical",
			OnExit: app.exitLogical,
		},
		lparser.ToEventsParams{
			Token:  "standardOperation",
			OnExit: app.exitStandardOperation,
		},
		lparser.ToEventsParams{
			Token:  "remainingOperation",
			OnExit: app.exitRemainingOperation,
		},
		lparser.ToEventsParams{
			Token:  "print",
			OnExit: app.exitPrint,
		},
		lparser.ToEventsParams{
			Token:  "jump",
			OnExit: app.exitJump,
		},
		lparser.ToEventsParams{
			Token:  "exit",
			OnExit: app.exitExit,
		},
		lparser.ToEventsParams{
			Token:  "call",
			OnExit: app.exitCall,
		},
		lparser.ToEventsParams{
			Token:  "module",
			OnExit: app.exitModule,
		},
		lparser.ToEventsParams{
			Token:  "switch",
			OnExit: app.exitSwitch,
		},
		lparser.ToEventsParams{
			Token:  "save",
			OnExit: app.exitSave,
		},
		lparser.ToEventsParams{
			Token:  "match",
			OnExit: app.exitMatch,
		},
		lparser.ToEventsParams{
			Token:  "matchPattern",
			OnExit: app.exitMatchPattern,
		},
		lparser.ToEventsParams{
			Token:  "stackFrame",
			OnExit: app.exitStackFrame,
		},
		lparser.ToEventsParams{
			Token:  "index",
			OnExit: app.exitIndex,
		},
		lparser.ToEventsParams{
			Token:  "skip",
			OnExit: app.exitSkip,
		},
		lparser.ToEventsParams{
			Token:  "intPointer",
			OnExit: app.exitIntPointer,
		},
	}

	ins, err := app.parserBuilder.Create().WithEventParams(params).WithLexer(lexer).Now()
	if err != nil {
		return nil, err
	}

	return app.lexerApplication.Execute(ins)
}

// ExecuteFile executes the parser on a script written on file
func (app *parser) ExecuteFile(filePath string) (interface{}, error) {
	if app.lexerAdapter == nil {
		return nil, errors.New("the Lexer lexerAdapter must be set in order to use the ExecuteFile method")
	}

	script, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return app.ExecuteScript(string(script))
}

// ExecuteScript executes the parser on a script
func (app *parser) ExecuteScript(script string) (interface{}, error) {
	if app.lexerAdapter == nil {
		return nil, errors.New("the Lexer lexerAdapter must be set in order to use the ExecuteScript method")
	}

	lexer, err := app.lexerAdapter.ToLexer(string(script))
	if err != nil {
		return nil, err
	}

	return app.Execute(lexer)
}

func (app *parser) init() {
	app.program = map[string]Program{}
	app.testable = map[string]Testable{}
	app.executable = map[string]Executable{}
	app.scopes = map[string]Scopes{}
	app.scope = map[string]Scope{}
	app.command = map[string]Command{}
	app.languageCommand = map[string]LanguageCommand{}
	app.scriptCommand = map[string]ScriptCommand{}
	app.headCommand = map[string]HeadCommand{}
	app.mainCommand = map[string]MainCommand{}
	app.mainCommandInstruction = map[string]MainCommandInstruction{}
	app.testCommand = map[string]TestCommand{}
	app.testCommandInstruction = map[string]TestCommandInstruction{}
	app.labelCommand = map[string]LabelCommand{}
	app.labelCommandInstruction = map[string]LabelCommandInstruction{}
	app.languageApplication = map[string]LanguageApplication{}
	app.languageMainSection = map[string]LanguageMainSection{}
	app.languageTestSection = map[string]LanguageTestSection{}
	app.languageTestDeclaration = map[string]LanguageTestDeclaration{}
	app.languageTestInstruction = map[string]LanguageTestInstruction{}
	app.languageLabelSection = map[string]LanguageLabelSection{}
	app.languageLabelDeclaration = map[string]LanguageLabelDeclaration{}
	app.languageLabelInstruction = map[string]LanguageLabelInstruction{}
	app.languageInstruction = map[string]LanguageInstruction{}
	app.languageInstructionCommon = map[string]LanguageInstructionCommon{}
	app.languageDefinition = map[string]LanguageDefinition{}
	app.languageValue = map[string]LanguageValue{}
	app.targetPath = map[string]RelativePath{}
	app.script = map[string]Script{}
	app.scriptValue = map[string]ScriptValue{}
	app.scriptTest = map[string]ScriptTest{}
	app.scriptTests = map[string]ScriptTests{}
	app.patternMatch = map[string]PatternMatch{}
	app.patternLabels = map[string]PatternLabels{}
	app.patternLabelEnter = map[string]string{}
	app.patternLabelExit = map[string]string{}
	app.relativePaths = map[string]RelativePaths{}
	app.relativePath = map[string]RelativePath{}
	app.folderSection = map[string]FolderSection{}
	app.folderName = map[string]FolderName{}
	app.application = map[string]Application{}
	app.testSection = map[string]TestSection{}
	app.testDeclaration = map[string]TestDeclaration{}
	app.testInstruction = map[string]TestInstruction{}
	app.assert = map[string]Assert{}
	app.readFile = map[string]ReadFile{}
	app.headSection = map[string]HeadSection{}
	app.headValue = map[string]HeadValue{}
	app.loadSingle = map[string]LoadSingle{}
	app.importSingle = map[string]ImportSingle{}
	app.labelSection = map[string]LabelSection{}
	app.labelDeclaration = map[string]LabelDeclaration{}
	app.labelInstruction = map[string]LabelInstruction{}
	app.mainSection = map[string]MainSection{}
	app.instruction = map[string]Instruction{}
	app.registry = map[string]Registry{}
	app.fetchRegistry = map[string]FetchRegistry{}
	app.unregister = map[string]Unregister{}
	app.register = map[string]Register{}
	app.specificTokenCode = map[string]SpecificTokenCode{}
	app.tokenSection = map[string]TokenSection{}
	app.codeMatch = map[string]CodeMatch{}
	app.token = map[string]Token{}
	app.callPattern = map[string]string{}
	app.callPatterns = map[string][]string{}
	app.patternOrRule = map[string]string{}
	app.importNames = map[string]string{}
	app.extendNames = map[string]string{}
	app.variable = map[string]Variable{}
	app.concatenation = map[string]Concatenation{}
	app.declaration = map[string]Declaration{}
	app.assignment = map[string]Assignment{}
	app.valueRepresentation = map[string]ValueRepresentation{}
	app.value = map[string]Value{}
	app.numericValue = map[string]NumericValue{}
	app.boolValue = map[string]bool{}
	app.floatValue = map[string]float64{}
	app.stringValue = map[string]string{}
	app.typ = map[string]Type{}
	app.operation = map[string]Operation{}
	app.arythmetic = map[string]Arythmetic{}
	app.relational = map[string]Relational{}
	app.logical = map[string]Logical{}
	app.standardOperation = map[string]StandardOperation{}
	app.remainingOperation = map[string]RemainingOperation{}
	app.print = map[string]Print{}
	app.jump = map[string]Jump{}
	app.match = map[string]Match{}
	app.matchPattern = map[string]string{}
	app.exit = map[string]Exit{}
	app.call = map[string]Call{}
	app.module = map[string]Module{}
	app.swtch = map[string]Switch{}
	app.save = map[string]Save{}
	app.stackFrame = map[string]StackFrame{}
	app.index = map[string]Index{}
	app.skip = map[string]Skip{}
	app.intPointer = map[string]IntPointer{}
}

func (app *parser) exitProgram(tree lexers.NodeTree) (interface{}, error) {
	builder := app.programBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"testable",
		"languageApplication",
	})

	switch section {
	case "testable":
		if testable, ok := app.testable[code]; ok {
			builder.WithTestable(testable)
		}
		break
	case "languageApplication":
		if lang, ok := app.languageApplication[code]; ok {
			builder.WithLanguage(lang)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.program[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTestable(tree lexers.NodeTree) (interface{}, error) {
	builder := app.testableBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"executable",
		"languageDefinition",
	})

	switch section {
	case "executable":
		if app, ok := app.executable[code]; ok {
			builder.WithExecutable(app)
		}
		break
	case "languageDefinition":
		if langDef, ok := app.languageDefinition[code]; ok {
			builder.WithLanguage(langDef)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.testable[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitExecutable(tree lexers.NodeTree) (interface{}, error) {
	builder := app.executableBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"application",
		"script",
	})

	switch section {
	case "application":
		if app, ok := app.application[code]; ok {
			builder.WithApplication(app)
		}
		break
	case "script":
		if script, ok := app.script[code]; ok {
			builder.WithScript(script)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.executable[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScopesWithArrow(tree lexers.NodeTree) (interface{}, error) {
	scopesCode := tree.CodeFromName("scopes")
	if scopesCode != "" {
		if scopes, ok := app.scopes[scopesCode]; ok {
			app.scopes[tree.Code()] = scopes
			return scopes, nil
		}
	}

	return nil, errors.New("the scopes was expected in the scopesWithArrow token")
}

func (app *parser) exitScopes(tree lexers.NodeTree) (interface{}, error) {
	scopes := []Scope{}
	codes := tree.CodesFromName("scope")
	for _, oneCode := range codes {
		if val, ok := app.scope[oneCode]; ok {
			scopes = append(scopes, val)
		}
	}

	ins, err := app.scopesBuilder.Create().WithScopes(scopes).Now()
	if err != nil {
		return nil, err
	}

	app.scopes[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScope(tree lexers.NodeTree) (interface{}, error) {
	builder := app.scopeBuilder.Create()
	section, _ := tree.BestMatchFromNames([]string{
		"STAR",
		"PLUS",
	})

	switch section {
	case "STAR":
		builder.IsExternal()
	case "PLUS":
		builder.IsInternal()
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.scope[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.commandBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"languageCommand",
		"scriptCommand",
		"headCommand",
		"mainCommand",
		"testCommand",
		"labelCommand",
	})

	switch section {
	case "languageCommand":
		if ins, ok := app.languageCommand[code]; ok {
			builder.WithLanguage(ins)
		}

		break
	case "scriptCommand":
		if ins, ok := app.scriptCommand[code]; ok {
			builder.WithScript(ins)
		}

		break
	case "headCommand":
		if ins, ok := app.headCommand[code]; ok {
			builder.WithHead(ins)
		}

		break
	case "mainCommand":
		if ins, ok := app.mainCommand[code]; ok {
			builder.WithMain(ins)
		}

		break
	case "testCommand":
		if ins, ok := app.testCommand[code]; ok {
			builder.WithTest(ins)
		}

		break
	case "labelCommand":
		if ins, ok := app.labelCommand[code]; ok {
			builder.WithLabel(ins)
		}

		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.command[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageCommandBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	values := []LanguageValue{}
	langValueCodes := tree.CodesFromName("languageValue")
	for _, oneValCode := range langValueCodes {
		if val, ok := app.languageValue[oneValCode]; ok {
			values = append(values, val)
		}
	}

	if len(values) > 0 {
		builder.WithValues(values)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageCommand[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScriptCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.scriptCommandBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	values := []ScriptValue{}
	scriptValueCodes := tree.CodesFromName("scriptValue")
	for _, oneValCode := range scriptValueCodes {
		if val, ok := app.scriptValue[oneValCode]; ok {
			values = append(values, val)
		}
	}

	if len(values) > 0 {
		builder.WithValues(values)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.scriptCommand[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitHeadCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.headCommandBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	values := []HeadValue{}
	headValueCodes := tree.CodesFromName("headValue")
	for _, oneValCode := range headValueCodes {
		if val, ok := app.headValue[oneValCode]; ok {
			values = append(values, val)
		}
	}

	if len(values) > 0 {
		builder.WithValues(values)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.headCommand[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitMainCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.mainCommandBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	instructions := []MainCommandInstruction{}
	instructionCodes := tree.CodesFromName("mainCommandInstruction")
	for _, oneInsCode := range instructionCodes {
		if ins, ok := app.mainCommandInstruction[oneInsCode]; ok {
			instructions = append(instructions, ins)
		}
	}

	if len(instructions) > 0 {
		builder.WithInstructions(instructions)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.mainCommand[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitMainCommandInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.mainCommandInstructionBuilder.Create()
	scopesCode := tree.CodeFromName("scopesWithArrow")
	if scopesCode != "" {
		if ins, ok := app.scopes[scopesCode]; ok {
			builder.WithScopes(ins)
		}
	}

	instructionCode := tree.CodeFromName("instruction")
	if instructionCode != "" {
		if ins, ok := app.instruction[instructionCode]; ok {
			builder.WithInstruction(ins)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.mainCommandInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTestCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.testCommandBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	patternName := tree.CodeFromName("TEST_NAME_PATTERN")
	if patternName != "" {
		builder.WithName(patternName)
	}

	instructions := []TestCommandInstruction{}
	instructionCodes := tree.CodesFromName("testCommandInstruction")
	for _, oneInsCode := range instructionCodes {
		if ins, ok := app.testCommandInstruction[oneInsCode]; ok {
			instructions = append(instructions, ins)
		}
	}

	if len(instructions) > 0 {
		builder.WithInstructions(instructions)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.testCommand[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTestCommandInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.testCommandInstructionBuilder.Create()
	scopesCode := tree.CodeFromName("scopesWithArrow")
	if scopesCode != "" {
		if ins, ok := app.scopes[scopesCode]; ok {
			builder.WithScopes(ins)
		}
	}

	instructionCode := tree.CodeFromName("testInstruction")
	if instructionCode != "" {
		if ins, ok := app.testInstruction[instructionCode]; ok {
			builder.WithInstruction(ins)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.testCommandInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLabelCommand(tree lexers.NodeTree) (interface{}, error) {
	builder := app.labelCommandBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	labelName := tree.CodeFromName("LABEL_PATTERN")
	if labelName != "" {
		builder.WithName(labelName)
	}

	instructions := []LabelCommandInstruction{}
	instructionCodes := tree.CodesFromName("labelCommandInstruction")
	for _, oneInsCode := range instructionCodes {
		if ins, ok := app.labelCommandInstruction[oneInsCode]; ok {
			instructions = append(instructions, ins)
		}
	}

	if len(instructions) > 0 {
		builder.WithInstructions(instructions)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.labelCommand[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLabelCommandInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.labelCommandInstructionBuilder.Create()
	scopesCode := tree.CodeFromName("scopesWithArrow")
	if scopesCode != "" {
		if ins, ok := app.scopes[scopesCode]; ok {
			builder.WithScopes(ins)
		}
	}

	instructionCode := tree.CodeFromName("labelInstruction")
	if instructionCode != "" {
		if ins, ok := app.labelInstruction[instructionCode]; ok {
			builder.WithInstruction(ins)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.labelCommandInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageApplication(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageApplicationBuilder.Create()
	headSectionCode := tree.CodeFromName("headSection")
	if headSectionCode != "" {
		if ins, ok := app.headSection[headSectionCode]; ok {
			builder.WithHead(ins)
		}
	}

	labelSectionCode := tree.CodeFromName("languageLabelSection")
	if labelSectionCode != "" {
		if ins, ok := app.languageLabelSection[labelSectionCode]; ok {
			builder.WithLabels(ins)
		}
	}

	mainSectionCode := tree.CodeFromName("languageMainSection")
	if mainSectionCode != "" {
		if ins, ok := app.languageMainSection[mainSectionCode]; ok {
			builder.WithMain(ins)
		}
	}

	testSectionCode := tree.CodeFromName("languageTestSection")
	if testSectionCode != "" {
		if ins, ok := app.languageTestSection[testSectionCode]; ok {
			builder.WithTests(ins)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageApplication[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageMainSection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageMainSectionBuilder.Create()
	list := []LanguageInstruction{}
	codes := tree.CodesFromName("languageInstruction")
	for _, oneCode := range codes {
		if ins, ok := app.languageInstruction[oneCode]; ok {
			list = append(list, ins)
		}
	}

	if len(list) > 0 {
		builder.WithInstructions(list)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageMainSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageTestSection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageTestSectionBuilder.Create()
	list := []LanguageTestDeclaration{}
	codes := tree.CodesFromName("languageTestDeclaration")
	for _, oneCode := range codes {
		if ins, ok := app.languageTestDeclaration[oneCode]; ok {
			list = append(list, ins)
		}
	}

	if len(list) > 0 {
		builder.WithDeclarations(list)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageTestSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageTestDeclaration(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageTestDeclarationBuilder.Create()
	name := tree.CodeFromName("TEST_NAME_PATTERN")
	if name != "" {
		builder.WithName(name)
	}

	list := []LanguageTestInstruction{}
	codes := tree.CodesFromName("languageTestInstruction")
	for _, oneCode := range codes {
		if ins, ok := app.languageTestInstruction[oneCode]; ok {
			list = append(list, ins)
		}
	}

	if len(list) > 0 {
		builder.WithInstructions(list)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageTestDeclaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageTestInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageTestInstructionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"languageInstructionCommon",
		"testInstruction",
		"INTERPRET",
	})

	switch section {
	case "languageInstructionCommon":
		if ins, ok := app.languageInstructionCommon[code]; ok {
			builder.WithLanguageInstruction(ins)
		}

		break
	case "testInstruction":
		if ins, ok := app.testInstruction[code]; ok {
			builder.WithTestInstruction(ins)
		}

		break
	case "INTERPRET":
		builder.IsInterpret()
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageTestInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageLabelSection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageLabelSectionBuilder.Create()
	list := []LanguageLabelDeclaration{}
	codes := tree.CodesFromName("languageLabelDeclaration")
	for _, oneCode := range codes {
		if ins, ok := app.languageLabelDeclaration[oneCode]; ok {
			list = append(list, ins)
		}
	}

	if len(list) > 0 {
		builder.WithDeclarations(list)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageLabelSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageLabelDeclaration(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageLabelDeclarationBuilder.Create()
	labelName := tree.CodeFromName("LABEL_PATTERN")
	if labelName != "" {
		builder.WithName(labelName)
	}

	list := []LanguageLabelInstruction{}
	codes := tree.CodesFromName("languageLabelInstruction")
	for _, oneCode := range codes {
		if ins, ok := app.languageLabelInstruction[oneCode]; ok {
			list = append(list, ins)
		}
	}

	if len(list) > 0 {
		builder.WithInstructions(list)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageLabelDeclaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageLabelInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageLabelInstructionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"labelInstruction",
		"languageInstruction",
		"token",
	})

	switch section {
	case "labelInstruction":
		if ins, ok := app.labelInstruction[code]; ok {
			builder.WithLabelInstruction(ins)
		}

		break
	case "languageInstruction":
		if ins, ok := app.languageInstruction[code]; ok {
			builder.WithLanguageInstruction(ins)
		}

		break
	case "token":
		if ins, ok := app.token[code]; ok {
			builder.WithToken(ins)
		}

		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageLabelInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageInstructionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"languageInstructionCommon",
		"command",
	})

	switch section {
	case "languageInstructionCommon":
		if ins, ok := app.languageInstructionCommon[code]; ok {
			builder.WithInstruction(ins)
		}

		break
	case "command":
		if ins, ok := app.command[code]; ok {
			builder.WithCommand(ins)
		}

		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageInstructionCommon(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageInstructionCommonBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"instruction",
		"match",
	})

	switch section {
	case "instruction":
		if ins, ok := app.instruction[code]; ok {
			builder.WithInstruction(ins)
		}

		break
	case "match":
		if ins, ok := app.match[code]; ok {
			builder.WithMatch(ins)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageInstructionCommon[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageDefinition(tree lexers.NodeTree) (interface{}, error) {
	values := []LanguageValue{}
	codes := tree.CodesFromName("languageValue")
	for _, oneCode := range codes {
		if val, ok := app.languageValue[oneCode]; ok {
			values = append(values, val)
		}
	}

	ins, err := app.languageDefinitionBuilder.Create().WithValues(values).Now()
	if err != nil {
		return nil, err
	}

	app.languageDefinition[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageValueBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"LANG_ROOT",
		"singleExtend",
		"relativePath",
		"VARIABLE_PATTERN",
		"LANG_PATTERN_MATCHES",
	})

	switch section {
	case "LANG_ROOT":
		pattern := tree.CodeFromName("PATTERN_PATTERN")
		if pattern != "" {
			builder.WithRoot(pattern)
			break
		}
	case "singleExtend":
		if tree.CodeFromName("EXTENDS") != "" {
			extends := []RelativePath{}
			extendCodes := tree.CodesFromName("singleExtend")
			for _, oneExtendCode := range extendCodes {
				if single, ok := app.relativePath[oneExtendCode]; ok {
					extends = append(extends, single)
				}
			}

			builder.WithExtends(extends)
		}
		break
	case "relativePath":
		if filePath, ok := app.relativePath[code]; ok {
			if tree.CodeFromName("LANG_TOKENS") != "" {
				builder.WithTokens(filePath)
				break
			}

			if tree.CodeFromName("LANG_CHANNELS") != "" {
				builder.WithChannels(filePath)
				break
			}

			if tree.CodeFromName("LANG_RULES") != "" {
				builder.WithRules(filePath)
				break
			}

			if tree.CodeFromName("LANG_LOGIC") != "" {
				builder.WithLogic(filePath)
				break
			}
		}

	case "VARIABLE_PATTERN":
		if tree.CodeFromName("IN") != "" {
			builder.WithInputVariable(code)
			break
		}
		break
	case "LANG_PATTERN_MATCHES":
		patternMatches := []PatternMatch{}
		patternMatchCodes := tree.CodesFromName("patternMatch")
		for _, onePatternMatch := range patternMatchCodes {
			if patternMatch, ok := app.patternMatch[onePatternMatch]; ok {
				patternMatches = append(patternMatches, patternMatch)
			}
		}

		builder.WithPatternMatches(patternMatches)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageValue[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScript(tree lexers.NodeTree) (interface{}, error) {
	values := []ScriptValue{}
	codes := tree.CodesFromName("scriptValue")
	for _, oneCode := range codes {
		if val, ok := app.scriptValue[oneCode]; ok {
			values = append(values, val)
		}
	}

	ins, err := app.scriptBuilder.Create().WithValues(values).Now()
	if err != nil {
		return nil, err
	}

	app.script[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScriptValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.scriptValueBuilder.Create()
	section, _ := tree.BestMatchFromNames([]string{
		"SCRIPT_NAME",
		"SCRIPT_VERSION",
		"SCRIPT_SCRIPT",
		"SCRIPT_LANGUAGE",
		"SCRIPT_OUTPUT",
		"SCRIPT_TESTS",
	})

	switch section {
	case "SCRIPT_NAME":
		name := tree.CodeFromName("NAME_PATTERN")
		if name != "" {
			builder.WithName(name)
			break
		}
	case "SCRIPT_VERSION":
		version := tree.CodeFromName("VERSION_PATTERN")
		if version != "" {
			builder.WithVersion(version)
			break
		}
	case "SCRIPT_SCRIPT":
		pathCode := tree.CodeFromName("relativePath")
		if filePath, ok := app.relativePath[pathCode]; ok {
			builder.WithScriptPath(filePath)
			break
		}
	case "SCRIPT_LANGUAGE":
		pathCode := tree.CodeFromName("relativePath")
		if filePath, ok := app.relativePath[pathCode]; ok {
			builder.WithLanguagePath(filePath)
			break
		}
	case "SCRIPT_OUTPUT":
		variableName := tree.CodeFromName("VARIABLE_PATTERN")
		builder.WithOutput(variableName)
		break
	case "SCRIPT_TESTS":
		scriptTestsCode := tree.CodeFromName("scriptTests")
		if scriptTest, ok := app.scriptTests[scriptTestsCode]; ok {
			builder.WithScriptTests(scriptTest)
			break
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.scriptValue[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScriptTests(tree lexers.NodeTree) (interface{}, error) {
	list := []ScriptTest{}
	code := tree.CodeFromName("scriptTest")
	if script, ok := app.scriptTest[code]; ok {
		list = append(list, script)
	}

	codes := tree.CodesFromName("scriptTestWithComma")
	for _, oneCode := range codes {
		if script, ok := app.scriptTest[oneCode]; ok {
			list = append(list, script)
		}
	}

	builder := app.scriptTestsBuilder.Create()
	if len(list) > 0 {
		builder.WithTests(list)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.scriptTests[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitScriptTestWithComma(tree lexers.NodeTree) (interface{}, error) {
	code := tree.CodeFromName("scriptTest")
	if code != "" {
		if test, ok := app.scriptTest[code]; ok {
			app.scriptTest[tree.Code()] = test
			return test, nil
		}
	}

	return nil, errors.New("the scriptTest cannot be found")
}

func (app *parser) exitScriptTest(tree lexers.NodeTree) (interface{}, error) {
	builder := app.scriptTestBuilder.Create()
	name := tree.CodeFromName("TEST_NAME_PATTERN")
	if name != "" {
		builder.WithName(name)
	}

	pathCode := tree.CodeFromName("relativePath")
	if path, ok := app.relativePath[pathCode]; ok {
		builder.WithPath(path)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.scriptTest[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitPatternMatch(tree lexers.NodeTree) (interface{}, error) {
	builder := app.patternMatchBuilder.Create()
	pattern := tree.CodeFromName("PATTERN_PATTERN")
	if pattern != "" {
		builder.WithPattern(pattern)
	}

	patternLabelsCode := tree.CodeFromName("patternLabels")
	if patternLabelsCode != "" {
		if patternLabel, ok := app.patternLabels[patternLabelsCode]; ok {
			builder.WithLabels(patternLabel)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.patternMatch[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitPatternLabels(tree lexers.NodeTree) (interface{}, error) {
	builder := app.patternLabelsBuilder.Create()
	enterCode := tree.CodeFromName("patternLabelEnter")
	if code, ok := app.patternLabelEnter[enterCode]; ok {
		builder.WithEnterLabel(code)
	}

	exitCode := tree.CodeFromName("patternLabelExit")
	if code, ok := app.patternLabelExit[exitCode]; ok {
		builder.WithExitLabel(code)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.patternLabels[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitPatternLabelEnter(tree lexers.NodeTree) (interface{}, error) {
	label := tree.CodeFromName("LABEL_PATTERN")
	if label != "" {
		app.patternLabelEnter[tree.Code()] = label
		return label, nil
	}

	return nil, errors.New("the label is empty")
}

func (app *parser) exitPatternLabelExit(tree lexers.NodeTree) (interface{}, error) {
	label := tree.CodeFromName("LABEL_PATTERN")
	if label != "" {
		app.patternLabelExit[tree.Code()] = label
		return label, nil
	}

	return nil, errors.New("the label is empty")
}

func (app *parser) exitSingleExtend(tree lexers.NodeTree) (interface{}, error) {
	relPathCode := tree.CodeFromName("relativePath")
	if relPathCode != "" {
		if relPath, ok := app.relativePath[relPathCode]; ok {
			app.relativePath[tree.Code()] = relPath
			return relPath, nil
		}
	}

	return nil, errors.New("the singleExtend is invalid")
}

func (app *parser) exitRelativePaths(tree lexers.NodeTree) (interface{}, error) {
	relativePaths := []RelativePath{}
	code := tree.CodeFromName("relativePath")
	if code != "" {
		if rel, ok := app.relativePath[code]; ok {
			relativePaths = append(relativePaths, rel)
		}
	}

	codes := tree.CodesFromName("relativePathWithComma")
	for _, oneCode := range codes {
		if rel, ok := app.relativePath[oneCode]; ok {
			relativePaths = append(relativePaths, rel)
		}
	}

	ins, err := app.relativePathsBuilder.Create().WithRelativePaths(relativePaths).Now()
	if err != nil {
		return nil, err
	}

	app.relativePaths[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitRelativePathWithComma(tree lexers.NodeTree) (interface{}, error) {
	code := tree.CodeFromName("relativePath")
	if code != "" {
		if rel, ok := app.relativePath[code]; ok {
			app.relativePath[tree.Code()] = rel
			return rel, nil
		}
	}

	return nil, errors.New("the relativePath could not be found")
}

func (app *parser) exitRelativePath(tree lexers.NodeTree) (interface{}, error) {
	sections := []FolderSection{}
	codes := tree.CodesFromName("folderSection")
	for _, oneCode := range codes {
		if section, ok := app.folderSection[oneCode]; ok {
			sections = append(sections, section)
		}
	}

	ins, err := app.relativePathBuilder.Create().WithSections(sections).Now()
	if err != nil {
		return nil, err
	}

	app.relativePath[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitFolderSection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.folderSectionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"folderName",
		"FILE_PATTERN",
	})

	switch section {
	case "folderName":
		if folderName, ok := app.folderName[code]; ok {
			builder.WithName(folderName)
		}
		break
	case "FILE_PATTERN":
		folderName, err := app.folderNameBuilder.Create().WithName(code).Now()
		if err != nil {
			return nil, err
		}

		builder.IsTail().WithName(folderName)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.folderSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitFolderName(tree lexers.NodeTree) (interface{}, error) {
	builder := app.folderNameBuilder.Create()
	dots := tree.CodesFromName("DOT")
	amount := len(dots)
	if amount <= 0 {
		name := tree.CodeFromName("FILE_PATTERN")
		builder.WithName(name)
	}

	if amount == 1 {
		builder.IsCurrent()
	}

	if amount == 2 {
		builder.IsPrevious()
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.folderName[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitApplication(tree lexers.NodeTree) (interface{}, error) {
	builder := app.applicationBuilder.Create()
	headSectionCode := tree.CodeFromName("headSection")
	if headSectionCode != "" {
		if headSection, ok := app.headSection[headSectionCode]; ok {
			builder.WithHead(headSection)
		}
	}

	labelSectionCode := tree.CodeFromName("labelSection")
	if labelSectionCode != "" {
		if labelSection, ok := app.labelSection[labelSectionCode]; ok {
			builder.WithLabel(labelSection)
		}
	}

	mainSectionCode := tree.CodeFromName("mainSection")
	if mainSectionCode != "" {
		if mainSection, ok := app.mainSection[mainSectionCode]; ok {
			builder.WithMain(mainSection)
		}
	}

	testSectionCode := tree.CodeFromName("testSection")
	if testSectionCode != "" {
		if testSection, ok := app.testSection[testSectionCode]; ok {
			builder.WithTest(testSection)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.application[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTestSection(tree lexers.NodeTree) (interface{}, error) {
	lst := []TestDeclaration{}
	testDeclarationCodes := tree.CodesFromName("testDeclaration")
	builder := app.testSectionBuilder.Create()
	for _, oneTestDecl := range testDeclarationCodes {
		if testDecl, ok := app.testDeclaration[oneTestDecl]; ok {
			lst = append(lst, testDecl)
		}
	}

	if len(lst) > 0 {
		builder.WithDeclarations(lst)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.testSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTestDeclaration(tree lexers.NodeTree) (interface{}, error) {
	builder := app.testDeclarationBuilder.Create()
	nameCode := tree.CodeFromName("TEST_NAME_PATTERN")
	if nameCode != "" {
		builder.WithName(nameCode)
	}

	lst := []TestInstruction{}
	testInsCodes := tree.CodesFromName("testInstruction")
	for _, oneTestInsCode := range testInsCodes {
		if testIns, ok := app.testInstruction[oneTestInsCode]; ok {
			lst = append(lst, testIns)
		}
	}

	if len(lst) > 0 {
		builder.WithInstructions(lst)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.testDeclaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTestInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.testInstructionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"instruction",
		"readFile",
		"assert",
	})

	switch section {
	case "instruction":
		if ins, ok := app.instruction[code]; ok {
			builder.WithInstruction(ins)
		}
		break
	case "readFile":
		if readFile, ok := app.readFile[code]; ok {
			builder.WithReadFile(readFile)
		}
		break
	case "assert":
		if ass, ok := app.assert[code]; ok {
			builder.WithAssert(ass)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.testInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitAssert(tree lexers.NodeTree) (interface{}, error) {
	builder := app.assertBuilder.Create()
	intCode := tree.CodeFromName("INT")
	if intCode != "" {
		intValue, err := strconv.Atoi(intCode)
		if err != nil {
			return nil, err
		}

		builder.WithIndex(intValue)
	}

	condition := tree.CodeFromName("VARIABLE_PATTERN")
	if condition != "" {
		builder.WithCondition(condition)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.assert[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitReadFile(tree lexers.NodeTree) (interface{}, error) {
	builder := app.readFileBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	relPathCode := tree.CodeFromName("relativePath")
	if relPath, ok := app.relativePath[relPathCode]; ok {
		builder.WithPath(relPath)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.readFile[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitHeadSection(tree lexers.NodeTree) (interface{}, error) {
	lst := []HeadValue{}
	codes := tree.CodesFromName("headValue")
	for _, oneCode := range codes {
		if ins, ok := app.headValue[oneCode]; ok {
			lst = append(lst, ins)
		}
	}

	ins, err := app.headSectionBuilder.Create().WithValues(lst).Now()
	if err != nil {
		return nil, err
	}

	app.headSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitHeadValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.headValueBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"NAME_PATTERN",
		"VERSION_PATTERN",
		"IMPORTS",
		"LOADS",
	})

	switch section {
	case "NAME_PATTERN":
		builder.WithName(code)
		break
	case "VERSION_PATTERN":
		builder.WithVersion(code)
		break
	case "IMPORTS":
		imports := []ImportSingle{}
		codes := tree.CodesFromName("importSingle")
		for _, oneCode := range codes {
			if imp, ok := app.importSingle[oneCode]; ok {
				imports = append(imports, imp)
			}
		}

		builder.WithImport(imports)
		break
	case "LOADS":
		loads := []LoadSingle{}
		codes := tree.CodesFromName("loadSingle")
		for _, oneCode := range codes {
			if imp, ok := app.loadSingle[oneCode]; ok {
				loads = append(loads, imp)
			}
		}

		builder.WithLoad(loads)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.headValue[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitImportSingle(tree lexers.NodeTree) (interface{}, error) {
	builder := app.importSingleBuilder.Create()
	name := tree.CodeFromName("NAME_PATTERN")
	if name != "" {
		builder.WithName(name)
	}

	relPathCode := tree.CodeFromName("relativePath")
	if relPath, ok := app.relativePath[relPathCode]; ok {
		builder.WithPath(relPath)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.importSingle[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLoadSingle(tree lexers.NodeTree) (interface{}, error) {
	builder := app.loadSingleBuilder.Create()
	names := tree.CodesFromName("NAME_PATTERN")
	if len(names) != 2 {
		str := fmt.Sprintf("the loadSingle instruction was expecting %d names, %d provided", 2, len(names))
		return nil, errors.New(str)
	}

	ins, err := builder.WithInternal(names[0]).WithExternal(names[1]).Now()
	if err != nil {
		return nil, err
	}

	app.loadSingle[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLabelSection(tree lexers.NodeTree) (interface{}, error) {
	lst := []LabelDeclaration{}
	codes := tree.CodesFromName("labelDeclaration")
	for _, oneCode := range codes {
		if ins, ok := app.labelDeclaration[oneCode]; ok {
			lst = append(lst, ins)
		}
	}

	ins, err := app.labelSectionBuilder.Create().WithDeclarations(lst).Now()
	if err != nil {
		return nil, err
	}

	app.labelSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLabelDeclaration(tree lexers.NodeTree) (interface{}, error) {
	lst := []LabelInstruction{}
	codes := tree.CodesFromName("labelInstruction")
	for _, oneCode := range codes {
		if ins, ok := app.labelInstruction[oneCode]; ok {
			lst = append(lst, ins)
		}
	}

	name := tree.CodeFromName("LABEL_PATTERN")
	ins, err := app.labelDeclarationBuilder.Create().WithInstructions(lst).WithName(name).Now()
	if err != nil {
		return nil, err
	}

	app.labelDeclaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLabelInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.labelInstructionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"instruction",
		"RET",
	})

	switch section {
	case "instruction":
		if ins, ok := app.instruction[code]; ok {
			builder.WithInstruction(ins)
		}
		break
	case "RET":
		builder.IsRet()
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.labelInstruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitMainSection(tree lexers.NodeTree) (interface{}, error) {
	lst := []Instruction{}
	codes := tree.CodesFromName("instruction")
	for _, oneCode := range codes {
		if ins, ok := app.instruction[oneCode]; ok {
			lst = append(lst, ins)
		}
	}

	ins, err := app.mainSectionBuilder.Create().WithInstructions(lst).Now()
	if err != nil {
		return nil, err
	}

	app.mainSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitInstruction(tree lexers.NodeTree) (interface{}, error) {
	builder := app.instructionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"variable",
		"operation",
		"print",
		"stackFrame",
		"jump",
		"exit",
		"call",
		"module",
		"switch",
		"save",
		"registry",
	})

	switch section {
	case "variable":
		if variable, ok := app.variable[code]; ok {
			builder.WithVariable(variable)
		}
		break
	case "operation":
		if op, ok := app.operation[code]; ok {
			builder.WithOperation(op)
		}
		break
	case "print":
		if pr, ok := app.print[code]; ok {
			builder.WithPrint(pr)
		}
		break
	case "stackFrame":
		if stf, ok := app.stackFrame[code]; ok {
			builder.WithStackFrame(stf)
		}
		break
	case "jump":
		if jmp, ok := app.jump[code]; ok {
			builder.WithJump(jmp)
		}

		break
	case "exit":
		if exit, ok := app.exit[code]; ok {
			builder.WithExit(exit)
		}

		break
	case "call":
		if call, ok := app.call[code]; ok {
			builder.WithCall(call)
		}
		break
	case "module":
		if module, ok := app.module[code]; ok {
			builder.WithModule(module)
		}
		break
	case "switch":
		if swtch, ok := app.swtch[code]; ok {
			builder.WithSwitch(swtch)
		}
		break
	case "save":
		if save, ok := app.save[code]; ok {
			builder.WithSave(save)
		}
		break
	case "registry":
		if reg, ok := app.registry[code]; ok {
			builder.WithRegistry(reg)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.instruction[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitRegistry(tree lexers.NodeTree) (interface{}, error) {
	builder := app.registryBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"register",
		"unregister",
		"fetchRegistry",
	})

	switch section {
	case "register":
		if reg, ok := app.register[code]; ok {
			builder.WithRegister(reg)
		}
		break
	case "unregister":
		if unreg, ok := app.unregister[code]; ok {
			builder.WithUnregister(unreg)
		}
		break
	case "fetchRegistry":
		if fetch, ok := app.fetchRegistry[code]; ok {
			builder.WithFetch(fetch)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.registry[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitFetchRegistry(tree lexers.NodeTree) (interface{}, error) {
	builder := app.fetchRegistryBuilder.Create()
	variableNames := tree.CodesFromName("VARIABLE_PATTERN")
	if len(variableNames) != 2 {
		str := fmt.Sprintf("%d variableName was expected, %d returned", 2, len(variableNames))
		return nil, errors.New(str)
	}

	if variableNames[0] != "" {
		builder.To(variableNames[0])
	}

	if variableNames[1] != "" {
		builder.From(variableNames[1])
	}

	intPointerCode := tree.CodeFromName("intPointer")
	if intPointerCode != "" {
		if intPointer, ok := app.intPointer[intPointerCode]; ok {
			builder.WithIndex(intPointer)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.fetchRegistry[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitUnregister(tree lexers.NodeTree) (interface{}, error) {
	builder := app.unregisterBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.unregister[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitRegister(tree lexers.NodeTree) (interface{}, error) {
	builder := app.registerBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	intPointerCode := tree.CodeFromName("intPointer")
	if intPointerCode != "" {
		if intPointer, ok := app.intPointer[intPointerCode]; ok {
			builder.WithIndex(intPointer)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.register[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitToken(tree lexers.NodeTree) (interface{}, error) {
	builder := app.tokenBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"codeMatch",
		"tokenSection",
	})

	switch section {
	case "codeMatch":
		if codeMatch, ok := app.codeMatch[code]; ok {
			builder.WithCodeMatch(codeMatch)
		}
		break
	case "tokenSection":
		if tokenSection, ok := app.tokenSection[code]; ok {
			builder.WithTokenSection(tokenSection)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.token[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitCodeMatch(tree lexers.NodeTree) (interface{}, error) {
	variableNameCodes := tree.CodesFromName("VARIABLE_PATTERN")
	if len(variableNameCodes) != 2 {
		str := fmt.Sprintf("%d variableName was expected, %d returned", 2, len(variableNameCodes))
		return nil, errors.New(str)
	}

	content := tree.CodeFromName("CONTENT")
	fmt.Printf("\n->%s\n", content)

	builder := app.codeMatchBuilder.Create().WithContent(variableNameCodes[0]).WithSection(variableNameCodes[1])
	callPatternsCode := tree.CodeFromName("callPatterns")
	if callPatternsCode != "" {
		if callPatterns, ok := app.callPatterns[callPatternsCode]; ok {
			builder.WithPatternVariables(callPatterns)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.codeMatch[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTokenSection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.tokenSectionBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"specificTokenCodeWithAmount",
		"specificTokenCode",
		"VARIABLE_PATTERN",
	})

	switch section {
	case "specificTokenCodeWithAmount":
		if specificTokenCode, ok := app.specificTokenCode[code]; ok {
			builder.WithSpecific(specificTokenCode)
		}
		break
	case "specificTokenCode":
		if specificTokenCode, ok := app.specificTokenCode[code]; ok {
			builder.WithSpecific(specificTokenCode)
		}
		break
	case "VARIABLE_PATTERN":
		builder.WithVariableName(code)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.tokenSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitSpecificTokenCodeWithAmount(tree lexers.NodeTree) (interface{}, error) {
	variableNames := tree.CodesFromName("VARIABLE_PATTERN")
	if len(variableNames) != 2 {
		str := fmt.Sprintf("%d variableName was expected, %d returned", 2, len(variableNames))
		return nil, errors.New(str)
	}

	builder := app.specificTokenCodeBuilder.Create().WithVariableName(variableNames[0]).WithAmount(variableNames[1])
	callPatternCode := tree.CodeFromName("callPattern")
	if callPatternCode != "" {
		if callPattern, ok := app.callPattern[callPatternCode]; ok {
			builder.WithPatternVariable(callPattern)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.specificTokenCode[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitSpecificTokenCode(tree lexers.NodeTree) (interface{}, error) {
	builder := app.specificTokenCodeBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariableName(variableName)
	}

	callPatternCode := tree.CodeFromName("callPattern")
	if callPatternCode != "" {
		if callPattern, ok := app.callPattern[callPatternCode]; ok {
			builder.WithPatternVariable(callPattern)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.specificTokenCode[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitCallPatterns(tree lexers.NodeTree) (interface{}, error) {
	callPatterns := []string{}
	patternCode := tree.CodeFromName("callPattern")
	if patternCode != "" {
		if pattern, ok := app.callPattern[patternCode]; ok {
			callPatterns = append(callPatterns, pattern)
		}
	}

	patternCodes := tree.CodesFromName("pipeCallPattern")
	if len(patternCodes) > 0 {
		for _, onePatternCode := range patternCodes {
			if pattern, ok := app.callPattern[onePatternCode]; ok {
				callPatterns = append(callPatterns, pattern)
			}
		}
	}

	if len(callPatterns) > 0 {
		app.callPatterns[tree.Code()] = callPatterns
		return callPatterns, nil
	}

	return nil, errors.New("the callPatterns are invalid")
}

func (app *parser) exitPatternOrRule(tree lexers.NodeTree) (interface{}, error) {
	_, code := tree.BestMatchFromNames([]string{
		"RULE_PATTERN",
		"PATTERN_PATTERN",
	})

	app.patternOrRule[tree.Code()] = code
	return code, nil
}

func (app *parser) exitPipeCallPattern(tree lexers.NodeTree) (interface{}, error) {
	pattern := tree.CodeFromName("callPattern")
	if pattern != "" {
		app.callPattern[tree.Code()] = pattern
		return pattern, nil
	}

	return nil, errors.New("the pipeCallPattern is invalid")
}

func (app *parser) exitCallPattern(tree lexers.NodeTree) (interface{}, error) {
	patternOrRuleCode := tree.CodeFromName("patternOrRule")
	if patternOrRuleCode != "" {
		if name, ok := app.patternOrRule[patternOrRuleCode]; ok {
			app.callPattern[tree.Code()] = name
			return name, nil
		}
	}

	return nil, errors.New("the callPattern is invalid")
}

func (app *parser) exitVariable(tree lexers.NodeTree) (interface{}, error) {
	builder := app.variableBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"declaration",
		"assignment",
		"concatenation",
		"VARIABLE_PATTERN",
	})

	switch section {
	case "declaration":
		if decl, ok := app.declaration[code]; ok {
			builder.WithDeclaration(decl)
		}
		break
	case "assignment":
		if ass, ok := app.assignment[code]; ok {
			builder.WithAssigment(ass)
		}
		break
	case "concatenation":
		if concat, ok := app.concatenation[code]; ok {
			builder.WithConcatenation(concat)
		}
		break
	case "VARIABLE_PATTERN":
		builder.WithDelete(code)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.variable[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitConcatenation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.concatenationBuilder.Create()
	standardOperationCode := tree.CodeFromName("standardOperation")
	if standardOperationCode != "" {
		if op, ok := app.standardOperation[standardOperationCode]; ok {
			builder.WithOperation(op)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.concatenation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitDeclaration(tree lexers.NodeTree) (interface{}, error) {
	builder := app.declarationBuilder.Create()
	variableNameCode := tree.CodeFromName("VARIABLE_PATTERN")
	if variableNameCode != "" {
		builder.WithVariable(variableNameCode)
	}

	typeCode := tree.CodeFromName("type")
	if typeCode != "" {
		if typ, ok := app.typ[typeCode]; ok {
			builder.WithType(typ)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.declaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitAssignment(tree lexers.NodeTree) (interface{}, error) {
	builder := app.assignmentBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	valueCode := tree.CodeFromName("valueRepresentation")
	if valueCode != "" {
		if val, ok := app.valueRepresentation[valueCode]; ok {
			builder.WithValue(val)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.assignment[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitValueRepresentation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.valueRepresentationBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"VARIABLE_PATTERN",
		"value",
	})

	switch section {
	case "VARIABLE_PATTERN":
		builder.WithVariable(code)
		break
	case "value":
		if val, ok := app.value[code]; ok {
			builder.WithValue(val)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.valueRepresentation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.valueBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"numericValue",
		"boolValue",
		"stringValue",
		"structValue",
		"NIL",
	})

	switch section {
	case "numericValue":
		if val, ok := app.numericValue[code]; ok {
			builder.WithNumeric(val)
		}
		break
	case "boolValue":
		if bl, ok := app.boolValue[code]; ok {
			builder.WithBool(bl)
		}
		break
	case "stringValue":
		if str, ok := app.stringValue[code]; ok {
			builder.WithString(str)
		}
		break
	case "NIL":
		builder.IsNil()
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.value[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitNumericValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.numericValueBuilder.Create()
	minusCode := tree.CodeFromName("MINUS")
	if minusCode != "" {
		builder.IsNegative()
	}

	section, code := tree.BestMatchFromNames([]string{
		"INT",
		"floatValue",
	})

	switch section {
	case "INT":
		intValue, err := strconv.Atoi(code)
		if err != nil {
			return nil, err
		}

		builder.WithInt(intValue)
		break
	case "floatValue":
		if fl, ok := app.floatValue[code]; ok {
			builder.WithFloat(fl)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.numericValue[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitBool(tree lexers.NodeTree) (interface{}, error) {
	section, _ := tree.BestMatchFromNames([]string{
		"TRUE",
		"FALSE",
	})

	value := false
	switch section {
	case "TRUE":
		value = true
		break
	case "FALSE":
		value = false
		break
	}

	app.boolValue[tree.Code()] = value
	return value, nil
}

func (app *parser) exitFloatValue(tree lexers.NodeTree) (interface{}, error) {
	values := tree.CodesFromName("INT")
	if len(values) != 2 {
		str := fmt.Sprintf("two (2) INT were expected, %d provided", len(values))
		return nil, errors.New(str)
	}

	valueAsString := fmt.Sprintf("%s.%s", values[0], values[1])
	fl, err := strconv.ParseFloat(valueAsString, 64)
	if err != nil {
		return nil, err
	}

	app.floatValue[tree.Code()] = fl
	return fl, nil
}

func (app *parser) exitStringValue(tree lexers.NodeTree) (interface{}, error) {
	value := tree.CodeFromName("EVERYTHING_EXCEPT_QUOTATION")
	app.stringValue[tree.Code()] = value
	return value, nil
}

func (app *parser) exitType(tree lexers.NodeTree) (interface{}, error) {
	builder := app.typeBuilder.Create()
	section, _ := tree.BestMatchFromNames([]string{
		"BOOL",
		"INT_HEIGHT",
		"INT_SIXTEEN",
		"INT_THIRTY_TWO",
		"INT_SIXTY_FOUR",
		"FLOAT_THIRTY_TWO",
		"FLOAT_SIXTY_FOUR",
		"UINT_HEIGHT",
		"UINT_SIXTEEN",
		"UINT_THIRTY_TWO",
		"UINT_SIXTY_FOUR",
		"STRING",
		"STACKFRAME",
	})

	switch section {
	case "BOOL":
		builder.IsBool()
		break
	case "INT_HEIGHT":
		builder.IsInt8()
		break
	case "INT_SIXTEEN":
		builder.IsInt16()
		break
	case "INT_THIRTY_TWO":
		builder.IsInt32()
		break
	case "INT_SIXTY_FOUR":
		builder.IsInt64()
		break
	case "FLOAT_THIRTY_TWO":
		builder.IsFloat32()
		break
	case "FLOAT_SIXTY_FOUR":
		builder.IsFloat64()
		break
	case "UINT_HEIGHT":
		builder.IsUint8()
		break
	case "UINT_SIXTEEN":
		builder.IsUint16()
		break
	case "UINT_THIRTY_TWO":
		builder.IsUint32()
		break
	case "UINT_SIXTY_FOUR":
		builder.IsUint64()
		break
	case "STRING":
		builder.IsString()
		break
	case "STACKFRAME":
		builder.IsStackFrame()
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.typ[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitOperation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.operationBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"arythmetic",
		"relational",
		"logical",
	})

	switch section {
	case "arythmetic":
		if ary, ok := app.arythmetic[code]; ok {
			builder.WithArythmetic(ary)
		}
		break
	case "relational":
		if rel, ok := app.relational[code]; ok {
			builder.WithRelational(rel)
		}
		break
	case "logical":
		if log, ok := app.logical[code]; ok {
			builder.WithLogical(log)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.operation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitArythmetic(tree lexers.NodeTree) (interface{}, error) {
	builder := app.arythmeticBuilder.Create()
	section, _ := tree.BestMatchFromNames([]string{
		"ADD",
		"SUB",
		"MUL",
		"DIV",
	})

	var standard StandardOperation
	standardCode := tree.CodeFromName("standardOperation")
	if standardCode != "" {
		if st, ok := app.standardOperation[standardCode]; ok {
			standard = st
		}
	}

	var remaining RemainingOperation
	remainingCode := tree.CodeFromName("remainingOperation")
	if remainingCode != "" {
		if rem, ok := app.remainingOperation[remainingCode]; ok {
			remaining = rem
		}
	}

	switch section {
	case "ADD":
		builder.WithAddition(standard)
		break
	case "SUB":
		builder.WithSubstraction(standard)
		break
	case "MUL":
		builder.WithMultiplication(standard)
		break
	case "DIV":
		builder.WithDivision(remaining)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.arythmetic[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitRelational(tree lexers.NodeTree) (interface{}, error) {
	builder := app.relationalBuilder.Create()
	section, _ := tree.BestMatchFromNames([]string{
		"LESS_THAN",
		"EQUAL_INS",
		"NOT_EQUAL_INS",
	})

	var standard StandardOperation
	standardCode := tree.CodeFromName("standardOperation")
	if standardCode != "" {
		if st, ok := app.standardOperation[standardCode]; ok {
			standard = st
		}
	}

	switch section {
	case "LESS_THAN":
		builder.WithLessThan(standard)
		break
	case "EQUAL_INS":
		builder.WithEqual(standard)
		break
	case "NOT_EQUAL_INS":
		builder.WithNotEqual(standard)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.relational[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLogical(tree lexers.NodeTree) (interface{}, error) {
	builder := app.logicalBuilder.Create()
	section, _ := tree.BestMatchFromNames([]string{
		"AND",
		"OR",
	})

	var standard StandardOperation
	standardCode := tree.CodeFromName("standardOperation")
	if standardCode != "" {
		if st, ok := app.standardOperation[standardCode]; ok {
			standard = st
		}
	}

	switch section {
	case "AND":
		builder.WithAnd(standard)
		break
	case "OR":
		builder.WithOr(standard)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.logical[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitStandardOperation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.standardOperationBuilder.Create()
	variableNames := tree.CodesFromName("VARIABLE_PATTERN")
	if len(variableNames) != 3 {
		str := fmt.Sprintf("three (3) variableNames were expected, %d provided", len(variableNames))
		return nil, errors.New(str)
	}

	ins, err := builder.WithResult(variableNames[0]).WithFirst(variableNames[1]).WithSecond(variableNames[2]).Now()
	if err != nil {
		return nil, err
	}

	app.standardOperation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitRemainingOperation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.remainingOperationBuilder.Create()
	variableNames := tree.CodesFromName("VARIABLE_PATTERN")
	if len(variableNames) != 4 {
		str := fmt.Sprintf("four (4) variableNames were expected, %d provided", len(variableNames))
		return nil, errors.New(str)
	}

	ins, err := builder.WithFirst(variableNames[2]).WithSecond(variableNames[3]).WithResult(variableNames[0]).WithRemaining(variableNames[1]).Now()
	if err != nil {
		return nil, err
	}

	app.remainingOperation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitPrint(tree lexers.NodeTree) (interface{}, error) {
	builder := app.printBuilder.Create()
	valueCode := tree.CodeFromName("valueRepresentation")
	if valueCode != "" {
		if val, ok := app.valueRepresentation[valueCode]; ok {
			builder.WithValue(val)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.print[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitJump(tree lexers.NodeTree) (interface{}, error) {
	builder := app.jumpBuilder.Create()
	condition := tree.CodeFromName("VARIABLE_PATTERN")

	if condition != "" {
		builder.WithCondition(condition)
	}

	labelCode := tree.CodeFromName("LABEL_PATTERN")
	if labelCode != "" {
		builder.WithLabel(labelCode)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.jump[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitMatch(tree lexers.NodeTree) (interface{}, error) {
	builder := app.matchBuilder.Create()
	input := tree.CodeFromName("VARIABLE_PATTERN")
	if input != "" {
		builder.WithInput(input)
	}

	patternCode := tree.CodeFromName("matchPattern")
	if pattern, ok := app.matchPattern[patternCode]; ok {
		builder.WithPattern(pattern)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.match[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitMatchPattern(tree lexers.NodeTree) (interface{}, error) {
	pattern := tree.CodeFromName("PATTERN_PATTERN")
	if pattern != "" {
		app.matchPattern[tree.Code()] = pattern
		return pattern, nil
	}

	return nil, errors.New("the match pattern is invalid")
}

func (app *parser) exitExit(tree lexers.NodeTree) (interface{}, error) {
	builder := app.exitBuilder.Create()
	condition := tree.CodeFromName("VARIABLE_PATTERN")
	if condition != "" {
		builder.WithCondition(condition)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.exit[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitCall(tree lexers.NodeTree) (interface{}, error) {
	builder := app.callBuilder.Create()
	variables := tree.CodesFromName("VARIABLE_PATTERN")
	amount := len(variables)
	if amount == 1 {
		builder.WithStackFrame(variables[0])
	}

	if amount == 2 {
		builder.WithStackFrame(variables[1]).WithCondition(variables[0])
	}

	if amount > 2 {
		str := fmt.Sprintf("the call was expecting 1 or 2 variables, %d provided", amount)
		return nil, errors.New(str)
	}

	name := tree.CodeFromName("NAME_PATTERN")
	if name != "" {
		builder.WithName(name)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.call[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitModule(tree lexers.NodeTree) (interface{}, error) {
	builder := app.moduleBuilder.Create()
	variable := tree.CodeFromName("VARIABLE_PATTERN")
	if variable != "" {
		builder.WithStackFrame(variable)
	}

	names := tree.CodesFromName("NAME_PATTERN")
	if len(names) != 2 {
		str := fmt.Sprintf("%d names were expected, %d provided", 2, len(names))
		return nil, errors.New(str)
	}

	ins, err := builder.WithName(names[0]).WithSymbol(names[1]).Now()
	if err != nil {
		return nil, err
	}

	app.module[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitSwitch(tree lexers.NodeTree) (interface{}, error) {
	builder := app.switchBuilder.Create()
	variable := tree.CodeFromName("VARIABLE_PATTERN")
	if variable != "" {
		builder.WithVariable(variable)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.swtch[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitSave(tree lexers.NodeTree) (interface{}, error) {
	builder := app.saveBuilder.Create()
	variables := tree.CodesFromName("VARIABLE_PATTERN")
	amount := len(variables)
	if amount > 2 {
		str := fmt.Sprintf("the save instruction was not expecting more than 2 variables, %d provided", amount)
		return nil, errors.New(str)
	}

	if amount <= 0 {
		str := fmt.Sprintf("the save instruction was expecting at least 1 variable, %d provided", amount)
		return nil, errors.New(str)
	}

	builder.To(variables[0])
	if amount == 2 {
		builder.From(variables[1])
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.save[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitStackFrame(tree lexers.NodeTree) (interface{}, error) {
	builder := app.stackFrameBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"PUSH",
		"POP",
		"skip",
		"index",
	})

	switch section {
	case "PUSH":
		builder.IsPush()
		break
	case "POP":
		builder.IsPop()
		break
	case "skip":
		if skip, ok := app.skip[code]; ok {
			builder.WithSkip(skip)
		}
		break
	case "index":
		if index, ok := app.index[code]; ok {
			builder.WithIndex(index)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.stackFrame[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitIndex(tree lexers.NodeTree) (interface{}, error) {
	builder := app.indexBuilder.Create()
	variableName := tree.CodeFromName("VARIABLE_PATTERN")
	if variableName != "" {
		builder.WithVariable(variableName)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.index[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitSkip(tree lexers.NodeTree) (interface{}, error) {
	builder := app.skipBuilder.Create()
	pointerCode := tree.CodeFromName("intPointer")
	if pointerCode != "" {
		if intPointer, ok := app.intPointer[pointerCode]; ok {
			builder.WithPointer(intPointer)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.skip[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitIntPointer(tree lexers.NodeTree) (interface{}, error) {
	builder := app.intPointerBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"INT",
		"VARIABLE_PATTERN",
	})

	switch section {
	case "INT":
		val, err := strconv.Atoi(code)
		if err != nil {
			return nil, err
		}

		builder.WithInt(int64(val))
		break
	case "VARIABLE_PATTERN":
		builder.WithVariable(code)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.intPointer[tree.Code()] = ins
	return ins, nil
}
