package parsers

import (
	"strings"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lparser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
)

const quaotationChar = '"'
const scriptHeadDelimiter = "----"

// NewProgramBuilder creates a new program builder
func NewProgramBuilder() ProgramBuilder {
	return createProgramBuilder()
}

// NewTestableBuilder creates a new testable builder
func NewTestableBuilder() TestableBuilder {
	return createTestableBuilder()
}

// NewExecutableBuilder creates a new executable builder
func NewExecutableBuilder() ExecutableBuilder {
	return createExecutableBuilder()
}

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() ApplicationBuilder {
	return createApplicationBuilder()
}

// NewScopesBuilder creates a new scopes builder
func NewScopesBuilder() ScopesBuilder {
	return createScopesBuilder()
}

// NewScopeBuilder creates a new scope builder
func NewScopeBuilder() ScopeBuilder {
	return createScopeBuilder()
}

// NewCommandBuilder creates a new command builder
func NewCommandBuilder() CommandBuilder {
	return createCommandBuilder()
}

// NewLanguageCommandBuilder creates a new language command builder
func NewLanguageCommandBuilder() LanguageCommandBuilder {
	return createLanguageCommandBuilder()
}

// NewScriptCommandBuilder creates a new script command builder
func NewScriptCommandBuilder() ScriptCommandBuilder {
	return createScriptCommandBuilder()
}

// NewHeadCommandBuilder creates a new head command builder
func NewHeadCommandBuilder() HeadCommandBuilder {
	return createHeadCommandBuilder()
}

// NewMainCommandBuilder creates a new main command builder
func NewMainCommandBuilder() MainCommandBuilder {
	return createMainCommandBuilder()
}

// NewMainCommandInstructionBuilder creates a new main command instruction builder
func NewMainCommandInstructionBuilder() MainCommandInstructionBuilder {
	return createMainCommandInstructionBuilder()
}

// NewTestCommandBuilder creates a new test command builder
func NewTestCommandBuilder() TestCommandBuilder {
	return createTestCommandBuilder()
}

// NewTestCommandInstructionBuilder creates a new test command instruction builder
func NewTestCommandInstructionBuilder() TestCommandInstructionBuilder {
	return createTestCommandInstructionBuilder()
}

// NewLabelCommandBuilder creates a new label command builder
func NewLabelCommandBuilder() LabelCommandBuilder {
	return createLabelCommandBuilder()
}

// NewLabelCommandInstructionBuilder creates a new label command instruction builder
func NewLabelCommandInstructionBuilder() LabelCommandInstructionBuilder {
	return createLabelCommandInstructionBuilder()
}

// NewLanguageApplicationBuilder creates a new language application builder
func NewLanguageApplicationBuilder() LanguageApplicationBuilder {
	return createLanguageApplicationBuilder()
}

// NewLanguageMainSectionBuilder creates a new language main section builder
func NewLanguageMainSectionBuilder() LanguageMainSectionBuilder {
	return createLanguageMainSectionBuilder()
}

// NewLanguageTestSectionBuilder creates a new language test section builder
func NewLanguageTestSectionBuilder() LanguageTestSectionBuilder {
	return createLanguageTestSectionBuilder()
}

// NewLanguageTestDeclarationBuilder creates a new language test declaration builder
func NewLanguageTestDeclarationBuilder() LanguageTestDeclarationBuilder {
	return createLanguageTestDeclarationBuilder()
}

// NewLanguageTestInstructionBuilder creates a new language test instruction builder
func NewLanguageTestInstructionBuilder() LanguageTestInstructionBuilder {
	return createLanguageTestInstructionBuilder()
}

// NewLanguageLabelSectionBuilder creates a new language label section builder
func NewLanguageLabelSectionBuilder() LanguageLabelSectionBuilder {
	return createLanguageLabelSectionBuilder()
}

// NewLanguageLabelDeclarationBuilder creates a new language label declaration builder
func NewLanguageLabelDeclarationBuilder() LanguageLabelDeclarationBuilder {
	return createLanguageLabelDeclarationBuilder()
}

// NewLanguageLabelInstructionBuilder creates a new language label instruction builder
func NewLanguageLabelInstructionBuilder() LanguageLabelInstructionBuilder {
	return createLanguageLabelInstructionBuilder()
}

// NewLanguageInstructionBuilder creates a new language instruction builder
func NewLanguageInstructionBuilder() LanguageInstructionBuilder {
	return createLanguageInstructionBuilder()
}

// NewLanguageInstructionCommonBuilder creates a new language instruction common builder
func NewLanguageInstructionCommonBuilder() LanguageInstructionCommonBuilder {
	return createLanguageInstructionCommonBuilder()
}

// NewLanguageDefinitionBuilder creates a new language definition builder
func NewLanguageDefinitionBuilder() LanguageDefinitionBuilder {
	return createLanguageDefinitionBuilder()
}

// NewLanguageValueBuilder creates a new language value builder
func NewLanguageValueBuilder() LanguageValueBuilder {
	return createLanguageValueBuilder()
}

// NewScriptBuilder creates a new script builder
func NewScriptBuilder() ScriptBuilder {
	return createScriptBuilder()
}

// NewScriptValueBuilder creates a new script value builder
func NewScriptValueBuilder() ScriptValueBuilder {
	return createScriptValueBuilder()
}

// NewScriptTestsBuilder creates a new script tests builder
func NewScriptTestsBuilder() ScriptTestsBuilder {
	return createScriptTestsBuilder()
}

// NewScriptTestBuilder creates a new script test builder
func NewScriptTestBuilder() ScriptTestBuilder {
	return createScriptTestBuilder()
}

// NewPatternMatchBuilder creates a new pattern match builder
func NewPatternMatchBuilder() PatternMatchBuilder {
	return createPatternMatchBuilder()
}

// NewPatternLabelsBuilder creeates a new pattern labels builder
func NewPatternLabelsBuilder() PatternLabelsBuilder {
	return createPatternLabelsBuilder()
}

// NewRelativePathsBuilder creates a new relative paths builder
func NewRelativePathsBuilder() RelativePathsBuilder {
	return createRelativePathsBuilder()
}

// NewFolderNameBuilder creates a new folder name builder
func NewFolderNameBuilder() FolderNameBuilder {
	return createFolderNameBuilder()
}

// NewFolderSectionBuilder creates a new folder section builder
func NewFolderSectionBuilder() FolderSectionBuilder {
	return createFolderSectionBuilder()
}

// NewRelativePathBuilder creates a new relative path builder
func NewRelativePathBuilder() RelativePathBuilder {
	folderNameBuilder := NewFolderNameBuilder()
	folderSectionBuilder := NewFolderSectionBuilder()
	return createRelativePathBuilder(folderNameBuilder, folderSectionBuilder)
}

// NewTestSectionBuilder creates a new test section builder
func NewTestSectionBuilder() TestSectionBuilder {
	return createTestSectionBuilder()
}

// NewTestDeclarationBuilder creates a new test declaration builder
func NewTestDeclarationBuilder() TestDeclarationBuilder {
	return createTestDeclarationBuilder()
}

// NewTestInstructionBuilder creates a new test instruction builder
func NewTestInstructionBuilder() TestInstructionBuilder {
	return createTestInstructionBuilder()
}

// NewAssertBuilder creates a new assert builder
func NewAssertBuilder() AssertBuilder {
	return createAssertBuilder()
}

// NewReadFileBuilder creates a new read file builder
func NewReadFileBuilder() ReadFileBuilder {
	return createReadFileBuilder()
}

// NewHeadSectionBuilder creates a new head section builder
func NewHeadSectionBuilder() HeadSectionBuilder {
	return createHeadSectionBuilder()
}

// NewHeadValueBuilder creates a new head value builder
func NewHeadValueBuilder() HeadValueBuilder {
	return createHeadValueBuilder()
}

// NewImportSingleBuilder creates a new import single builder
func NewImportSingleBuilder() ImportSingleBuilder {
	return createImportSingleBuilder()
}

// NewLabelSectionBuilder creates a new label section builder
func NewLabelSectionBuilder() LabelSectionBuilder {
	return createLabelSectionBuilder()
}

// NewLabelDeclarationBuilder creates a new label declaration builder
func NewLabelDeclarationBuilder() LabelDeclarationBuilder {
	return createLabelDeclarationBuilder()
}

// NewLabelInstructonBuilder creates a new label instruction builder
func NewLabelInstructonBuilder() LabelInstructionBuilder {
	return createLabelInstructionBuilder()
}

// NewMainSectionBuilder creates a new main section builder
func NewMainSectionBuilder() MainSectionBuilder {
	return createMainSectionBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewRegistryBuilder creates a new registry builder
func NewRegistryBuilder() RegistryBuilder {
	return createRegistryBuilder()
}

// NewFetchRegistryBuilder creates a new fetch registry builder
func NewFetchRegistryBuilder() FetchRegistryBuilder {
	return createFetchRegistryBuilder()
}

// NewUnregisterBuilder creates a new unregister builder
func NewUnregisterBuilder() UnregisterBuilder {
	return createUnregisterBuilder()
}

// NewRegisterBuilder creates a new register builder
func NewRegisterBuilder() RegisterBuilder {
	return createRegisterBuilder()
}

// NewSpecificTokenCodeBuilder creates a new specific token code builder
func NewSpecificTokenCodeBuilder() SpecificTokenCodeBuilder {
	return createSpecificTokenCodeBuilder()
}

// NewTokenSectionBuilder creates a new token section builder
func NewTokenSectionBuilder() TokenSectionBuilder {
	return createTokenSectionBuilder()
}

// NewCodeMatchBuilder creates a new code match builder
func NewCodeMatchBuilder() CodeMatchBuilder {
	return createCodeMatchBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewVariableBuilder creates a new variable builder
func NewVariableBuilder() VariableBuilder {
	return createVariableBuilder()
}

// NewConcatenationBuilder creates a new concatenation builder
func NewConcatenationBuilder() ConcatenationBuilder {
	return createConcatenationBuilder()
}

// NewDeclarationBuilder creates a new declaration builder
func NewDeclarationBuilder() DeclarationBuilder {
	return createDeclarationBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewValueRepresentationBuilder creates a new value representation builder
func NewValueRepresentationBuilder() ValueRepresentationBuilder {
	return createValueRepresentationBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// NewNumericValueBuilder creates a new numeric value buildeer
func NewNumericValueBuilder() NumericValueBuilder {
	return createNumericValueBuilder()
}

// NewTypeBuilder creates a new type builder
func NewTypeBuilder() TypeBuilder {
	return createTypeBuilder()
}

// NewOperationBuilder creates a new operation builder
func NewOperationBuilder() OperationBuilder {
	return createOperationalBuilder()
}

// NewArythmeticBuilder creates a new arythmetic builder
func NewArythmeticBuilder() ArythmeticBuilder {
	return createArythmeticBuilder()
}

// NewRelationalBuilder creates a new relational builder
func NewRelationalBuilder() RelationalBuilder {
	return createRelationalBuilder()
}

// NewLogicalBuilder creates a new logical builder
func NewLogicalBuilder() LogicalBuilder {
	return createLogicalBuilder()
}

// NewStandardOperationBuilder creates a new standard operation builder
func NewStandardOperationBuilder() StandardOperationBuilder {
	return createStandardOperationBuilder()
}

// NewRemainingOperationBuilder creates a new remaining operation builder
func NewRemainingOperationBuilder() RemainingOperationBuilder {
	return createRemainingOperationBuilder()
}

// NewPrintBuilder creates a new print builder
func NewPrintBuilder() PrintBuilder {
	return createPrintBuilder()
}

// NewJumpBuilder creates a new jump builder
func NewJumpBuilder() JumpBuilder {
	return createJumpBuilder()
}

// NewMatchBuilder creates a new match builder
func NewMatchBuilder() MatchBuilder {
	return createMatchBuilder()
}

// NewExitBuilder creates a new exit builder
func NewExitBuilder() ExitBuilder {
	return createExitBuilder()
}

// NewCallBuilder creates a new call builder
func NewCallBuilder() CallBuilder {
	return createCallBuilder()
}

// NewSwitchBuilder creates a new switch builder
func NewSwitchBuilder() SwitchBuilder {
	return createSwitchBuilder()
}

// NewSaveBuilder creates a new save builder
func NewSaveBuilder() SaveBuilder {
	return createSaveBuilder()
}

// NewStackFrameBuilder creates a new stack frame builder
func NewStackFrameBuilder() StackFrameBuilder {
	return createStackFrameBuilder()
}

// NewIndexBuilder creates a new index builder
func NewIndexBuilder() IndexBuilder {
	return createIndexBuilder()
}

// NewSkipBuilder creates a new skip builder
func NewSkipBuilder() SkipBuilder {
	return createSkipBuilder()
}

// NewIntPointerBuilder creates a new int pointer builder
func NewIntPointerBuilder() IntPointerBuilder {
	return createIntPointerBuilder()
}

// NewParserBuilder creates a new ParserBuilder instance
func NewParserBuilder() ParserBuilder {
	application := lparser.NewApplication()
	parserBuilder := lparser.NewBuilder()
	lexerBuilder := lexers.NewBuilder()
	programBuilder := NewProgramBuilder()
	testableBuilder := NewTestableBuilder()
	executableBuilder := NewExecutableBuilder()
	scopesBuilder := NewScopesBuilder()
	scopeBuilder := NewScopeBuilder()
	commandBuilder := NewCommandBuilder()
	languageCommandBuilder := NewLanguageCommandBuilder()
	scriptCommandBuilder := NewScriptCommandBuilder()
	headCommandBuilder := NewHeadCommandBuilder()
	mainCommandBuilder := NewMainCommandBuilder()
	mainCommandInstructionBuilder := NewMainCommandInstructionBuilder()
	testCommandBuilder := NewTestCommandBuilder()
	testCommandInstructionBuilder := NewTestCommandInstructionBuilder()
	labelCommandBuilder := NewLabelCommandBuilder()
	labelCommandInstructionBuilder := NewLabelCommandInstructionBuilder()
	languageApplicationBuilder := NewLanguageApplicationBuilder()
	languageMainSectionBuilder := NewLanguageMainSectionBuilder()
	languageTestSectionBuilder := NewLanguageTestSectionBuilder()
	languageTestDeclarationBuilder := NewLanguageTestDeclarationBuilder()
	languageTestInstructionBuilder := NewLanguageTestInstructionBuilder()
	languageLabelSectionBuilder := NewLanguageLabelSectionBuilder()
	languageLabelDeclarationBuilder := NewLanguageLabelDeclarationBuilder()
	languageLabelInstructionBuilder := NewLanguageLabelInstructionBuilder()
	languageInstructionBuilder := NewLanguageInstructionBuilder()
	languageInstructionCommonBuilder := NewLanguageInstructionCommonBuilder()
	languageDefinitionBuilder := NewLanguageDefinitionBuilder()
	languageValueBuilder := NewLanguageValueBuilder()
	scriptBuilder := NewScriptBuilder()
	scriptValueBuilder := NewScriptValueBuilder()
	scriptTestsBuilder := NewScriptTestsBuilder()
	scriptTestBuilder := NewScriptTestBuilder()
	patternMatchBuilder := NewPatternMatchBuilder()
	patternLabelsBuilder := NewPatternLabelsBuilder()
	relativePathsBuilder := NewRelativePathsBuilder()
	folderNameBuilder := NewFolderNameBuilder()
	folderSectionBuilder := NewFolderSectionBuilder()
	relativePathBuilder := NewRelativePathBuilder()
	applicationBuilder := NewApplicationBuilder()
	testSectionBuilder := NewTestSectionBuilder()
	testDeclarationBuilder := NewTestDeclarationBuilder()
	testInstructionBuilder := NewTestInstructionBuilder()
	assertBuilder := NewAssertBuilder()
	readFileBuilder := NewReadFileBuilder()
	headSectionBuilder := NewHeadSectionBuilder()
	headValueBuilder := NewHeadValueBuilder()
	importSingleBuilder := NewImportSingleBuilder()
	labelSectionBuilder := NewLabelSectionBuilder()
	labelDeclarationBuilder := NewLabelDeclarationBuilder()
	labelInstructionBuilder := NewLabelInstructonBuilder()
	mainSectionBuilder := NewMainSectionBuilder()
	instructionBuilder := NewInstructionBuilder()
	registryBuilder := NewRegistryBuilder()
	fetchRegistryBuilder := NewFetchRegistryBuilder()
	unregisterBuilder := NewUnregisterBuilder()
	registerBuilder := NewRegisterBuilder()
	specificTokenCodeBuilder := NewSpecificTokenCodeBuilder()
	tokenSectionBuilder := NewTokenSectionBuilder()
	codeMatchBuilder := NewCodeMatchBuilder()
	tokenBuilder := NewTokenBuilder()
	variableBuilder := NewVariableBuilder()
	concatenationBuilder := NewConcatenationBuilder()
	declarationBuilder := NewDeclarationBuilder()
	assignmentBuilder := NewAssignmentBuilder()
	valueRepresentationBuilder := NewValueRepresentationBuilder()
	valueBuilder := NewValueBuilder()
	numericValueBuilder := NewNumericValueBuilder()
	typeBuilder := NewTypeBuilder()
	operationBuilder := NewOperationBuilder()
	arythmeticBuilder := NewArythmeticBuilder()
	relationalBuilder := NewRelationalBuilder()
	logicalBuilder := NewLogicalBuilder()
	standardOperationBuilder := NewStandardOperationBuilder()
	remainingOperationBuilder := NewRemainingOperationBuilder()
	printBuilder := NewPrintBuilder()
	jumpBuilder := NewJumpBuilder()
	matchBuilder := NewMatchBuilder()
	exitBuilder := NewExitBuilder()
	callBuilder := NewCallBuilder()
	switchBuilder := NewSwitchBuilder()
	saveBuilder := NewSaveBuilder()
	stackFrameBuilder := NewStackFrameBuilder()
	indexBuilder := NewIndexBuilder()
	skipBuilder := NewSkipBuilder()
	intPointerBuilder := NewIntPointerBuilder()

	return createParserBuilder(
		application,
		parserBuilder,
		lexerBuilder,
		programBuilder,
		testableBuilder,
		executableBuilder,
		scopesBuilder,
		scopeBuilder,
		commandBuilder,
		languageCommandBuilder,
		scriptCommandBuilder,
		headCommandBuilder,
		mainCommandBuilder,
		mainCommandInstructionBuilder,
		testCommandBuilder,
		testCommandInstructionBuilder,
		labelCommandBuilder,
		labelCommandInstructionBuilder,
		languageApplicationBuilder,
		languageMainSectionBuilder,
		languageTestSectionBuilder,
		languageTestDeclarationBuilder,
		languageTestInstructionBuilder,
		languageLabelSectionBuilder,
		languageLabelDeclarationBuilder,
		languageLabelInstructionBuilder,
		languageInstructionBuilder,
		languageInstructionCommonBuilder,
		languageDefinitionBuilder,
		languageValueBuilder,
		scriptBuilder,
		scriptValueBuilder,
		scriptTestsBuilder,
		scriptTestBuilder,
		patternMatchBuilder,
		patternLabelsBuilder,
		relativePathsBuilder,
		relativePathBuilder,
		folderSectionBuilder,
		folderNameBuilder,
		applicationBuilder,
		testSectionBuilder,
		testDeclarationBuilder,
		testInstructionBuilder,
		assertBuilder,
		readFileBuilder,
		headSectionBuilder,
		headValueBuilder,
		importSingleBuilder,
		labelSectionBuilder,
		labelDeclarationBuilder,
		labelInstructionBuilder,
		mainSectionBuilder,
		instructionBuilder,
		registryBuilder,
		fetchRegistryBuilder,
		unregisterBuilder,
		registerBuilder,
		specificTokenCodeBuilder,
		tokenSectionBuilder,
		codeMatchBuilder,
		tokenBuilder,
		variableBuilder,
		concatenationBuilder,
		declarationBuilder,
		assignmentBuilder,
		valueRepresentationBuilder,
		valueBuilder,
		numericValueBuilder,
		typeBuilder,
		operationBuilder,
		arythmeticBuilder,
		relationalBuilder,
		logicalBuilder,
		standardOperationBuilder,
		remainingOperationBuilder,
		printBuilder,
		jumpBuilder,
		matchBuilder,
		exitBuilder,
		callBuilder,
		switchBuilder,
		saveBuilder,
		stackFrameBuilder,
		indexBuilder,
		skipBuilder,
		intPointerBuilder,
	)
}

// NewWhiteSpaceEvent creates a new whiteSpace event
func NewWhiteSpaceEvent(wsChannelName string) (lexers.Event, error) {
	evtFn := func(from int, to int, script []rune, rule grammar.Rule) []rune {
		if strings.Index(string(script), scriptHeadDelimiter) != -1 {
			subset := strings.TrimSpace(string(script[from:]))
			idx := strings.Index(subset, scriptHeadDelimiter)
			if idx == 0 {
				script = append(script[:from], script[to:]...)
				return []rune(
					strings.TrimSpace(string(script)),
				)
			}

			if idx == -1 {
				return script
			}
		}

		amount := strings.Count(string(script[:from]), string(quaotationChar))
		if amount%2 != 0 {
			return script
		}

		script = append(script[:from], script[to:]...)
		return script
	}

	return lexers.NewEventBuilder().Create().WithToken(wsChannelName).WithFn(evtFn).Now()
}

// ParserBuilder represents a parser builder
type ParserBuilder interface {
	Create() ParserBuilder
	WithLexerAdapter(lexerAdapter lexers.Adapter) ParserBuilder
	Now() (Parser, error)
}

// Parser represents a parser
type Parser interface {
	Execute(lexer lexers.Lexer) (interface{}, error)
	ExecuteFile(filePath string) (interface{}, error)
	ExecuteScript(script string) (interface{}, error)
}

// ProgramBuilder represents the program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithTestable(testable Testable) ProgramBuilder
	WithLanguage(lang LanguageApplication) ProgramBuilder
	Now() (Program, error)
}

// Program represents the program
type Program interface {
	IsTestable() bool
	Testable() Testable
	IsLanguage() bool
	Language() LanguageApplication
}

// TestableBuilder represents a testable builder
type TestableBuilder interface {
	Create() TestableBuilder
	WithExecutable(executable Executable) TestableBuilder
	WithLanguage(language LanguageDefinition) TestableBuilder
	Now() (Testable, error)
}

// Testable represents a testable program
type Testable interface {
	IsExecutable() bool
	Executable() Executable
	IsLanguage() bool
	Language() LanguageDefinition
}

// ExecutableBuilder represents an executable builder
type ExecutableBuilder interface {
	Create() ExecutableBuilder
	WithApplication(application Application) ExecutableBuilder
	WithScript(script Script) ExecutableBuilder
	Now() (Executable, error)
}

// Executable represents an executable program
type Executable interface {
	IsApplication() bool
	Application() Application
	IsScript() bool
	Script() Script
}

// ScriptBuilder represents a script builder
type ScriptBuilder interface {
	Create() ScriptBuilder
	WithValues(values []ScriptValue) ScriptBuilder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Name() string
	Version() string
	Script() RelativePath
	Language() RelativePath
	Output() string
	HasTests() bool
	Tests() ScriptTests
}

// ScriptValueBuilder represents a script value builder
type ScriptValueBuilder interface {
	Create() ScriptValueBuilder
	WithName(name string) ScriptValueBuilder
	WithVersion(version string) ScriptValueBuilder
	WithScriptPath(scriptPath RelativePath) ScriptValueBuilder
	WithLanguagePath(langPath RelativePath) ScriptValueBuilder
	WithOutput(output string) ScriptValueBuilder
	WithScriptTests(scriptTests ScriptTests) ScriptValueBuilder
	Now() (ScriptValue, error)
}

// ScriptValue represents a script value
type ScriptValue interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsScript() bool
	Script() RelativePath
	IsLanguage() bool
	Language() RelativePath
	IsOutput() bool
	Output() string
	IsScriptTests() bool
	ScriptTests() ScriptTests
}

// ScriptTestsBuilder represents a script tests builder
type ScriptTestsBuilder interface {
	Create() ScriptTestsBuilder
	WithTests(tests []ScriptTest) ScriptTestsBuilder
	Now() (ScriptTests, error)
}

// ScriptTests represents script tests
type ScriptTests interface {
	All() []ScriptTest
}

// ScriptTestBuilder represents a script test builder
type ScriptTestBuilder interface {
	Create() ScriptTestBuilder
	WithName(name string) ScriptTestBuilder
	WithPath(path RelativePath) ScriptTestBuilder
	Now() (ScriptTest, error)
}

// ScriptTest represents a script test
type ScriptTest interface {
	Name() string
	Path() RelativePath
}

// LanguageApplicationBuilder represents a language application builder
type LanguageApplicationBuilder interface {
	Create() LanguageApplicationBuilder
	WithHead(head HeadSection) LanguageApplicationBuilder
	WithLabels(labels LanguageLabelSection) LanguageApplicationBuilder
	WithMain(main LanguageMainSection) LanguageApplicationBuilder
	WithTests(tests LanguageTestSection) LanguageApplicationBuilder
	Now() (LanguageApplication, error)
}

// LanguageApplication represents a language application
type LanguageApplication interface {
	Head() HeadSection
	Labels() LanguageLabelSection
	Main() LanguageMainSection
	HasTests() bool
	Tests() LanguageTestSection
}

// LanguageMainSectionBuilder represents a language main section builder
type LanguageMainSectionBuilder interface {
	Create() LanguageMainSectionBuilder
	WithInstructions(instructions []LanguageInstruction) LanguageMainSectionBuilder
	Now() (LanguageMainSection, error)
}

// LanguageMainSection represents a language main section
type LanguageMainSection interface {
	Instructions() []LanguageInstruction
}

// LanguageTestSectionBuilder represents a language test section builder
type LanguageTestSectionBuilder interface {
	Create() LanguageTestSectionBuilder
	WithDeclarations(declarations []LanguageTestDeclaration) LanguageTestSectionBuilder
	Now() (LanguageTestSection, error)
}

// LanguageTestSection represents a language test section
type LanguageTestSection interface {
	Declarations() []LanguageTestDeclaration
}

// LanguageTestDeclarationBuilder represents a language test declaration builder
type LanguageTestDeclarationBuilder interface {
	Create() LanguageTestDeclarationBuilder
	WithName(name string) LanguageTestDeclarationBuilder
	WithInstructions(instructions []LanguageTestInstruction) LanguageTestDeclarationBuilder
	Now() (LanguageTestDeclaration, error)
}

// LanguageTestDeclaration represents a language test declaration
type LanguageTestDeclaration interface {
	Name() string
	Instructions() []LanguageTestInstruction
}

// LanguageTestInstructionBuilder represents a language test instruction builder
type LanguageTestInstructionBuilder interface {
	Create() LanguageTestInstructionBuilder
	WithLanguageInstruction(languageIns LanguageInstructionCommon) LanguageTestInstructionBuilder
	WithTestInstruction(testIns TestInstruction) LanguageTestInstructionBuilder
	IsInterpret() LanguageTestInstructionBuilder
	Now() (LanguageTestInstruction, error)
}

// LanguageTestInstruction represents a language test instruction
type LanguageTestInstruction interface {
	IsLanguageInstruction() bool
	LanguageInstruction() LanguageInstructionCommon
	IsTestInstruction() bool
	TestInstruction() TestInstruction
	IsInterpret() bool
}

// LanguageLabelSectionBuilder represents a language label section builder
type LanguageLabelSectionBuilder interface {
	Create() LanguageLabelSectionBuilder
	WithDeclarations(declarations []LanguageLabelDeclaration) LanguageLabelSectionBuilder
	Now() (LanguageLabelSection, error)
}

// LanguageLabelSection represents a language label section
type LanguageLabelSection interface {
	Declarations() []LanguageLabelDeclaration
}

// LanguageLabelDeclarationBuilder represents a language label declaration builder
type LanguageLabelDeclarationBuilder interface {
	Create() LanguageLabelDeclarationBuilder
	WithName(name string) LanguageLabelDeclarationBuilder
	WithInstructions(instructions []LanguageLabelInstruction) LanguageLabelDeclarationBuilder
	Now() (LanguageLabelDeclaration, error)
}

// LanguageLabelDeclaration represents a language label declaration
type LanguageLabelDeclaration interface {
	Name() string
	Instructions() []LanguageLabelInstruction
}

// LanguageLabelInstructionBuilder represents a language label instruction builder
type LanguageLabelInstructionBuilder interface {
	Create() LanguageLabelInstructionBuilder
	WithLanguageInstruction(languageInstruction LanguageInstruction) LanguageLabelInstructionBuilder
	WithLabelInstruction(labelInstruction LabelInstruction) LanguageLabelInstructionBuilder
	WithToken(token Token) LanguageLabelInstructionBuilder
	Now() (LanguageLabelInstruction, error)
}

// LanguageLabelInstruction represents a language label instruction
type LanguageLabelInstruction interface {
	IsLanguageInstruction() bool
	LanguageInstruction() LanguageInstruction
	IsLabelInstruction() bool
	LabelInstruction() LabelInstruction
	IsToken() bool
	Token() Token
}

// LanguageInstructionBuilder represents a language instruction builder
type LanguageInstructionBuilder interface {
	Create() LanguageInstructionBuilder
	WithInstruction(ins LanguageInstructionCommon) LanguageInstructionBuilder
	WithCommand(command Command) LanguageInstructionBuilder
	Now() (LanguageInstruction, error)
}

// LanguageInstruction represents a language instruction
type LanguageInstruction interface {
	IsInstruction() bool
	Instruction() LanguageInstructionCommon
	IsCommand() bool
	Command() Command
}

// LanguageInstructionCommonBuilder represents a language instruction common builder
type LanguageInstructionCommonBuilder interface {
	Create() LanguageInstructionCommonBuilder
	WithInstruction(ins Instruction) LanguageInstructionCommonBuilder
	WithMatch(match Match) LanguageInstructionCommonBuilder
	Now() (LanguageInstructionCommon, error)
}

// LanguageInstructionCommon represents a language instruction common
type LanguageInstructionCommon interface {
	IsInstruction() bool
	Instruction() Instruction
	IsMatch() bool
	Match() Match
}

// CommandBuilder represents a command builder
type CommandBuilder interface {
	Create() CommandBuilder
	WithLanguage(language LanguageCommand) CommandBuilder
	WithScript(script ScriptCommand) CommandBuilder
	WithHead(head HeadCommand) CommandBuilder
	WithMain(main MainCommand) CommandBuilder
	WithLabel(label LabelCommand) CommandBuilder
	WithTest(test TestCommand) CommandBuilder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	IsLanguage() bool
	Language() LanguageCommand
	IsScript() bool
	Script() ScriptCommand
	IsHead() bool
	Head() HeadCommand
	IsMain() bool
	Main() MainCommand
	IsLabel() bool
	Label() LabelCommand
	IsTest() bool
	Test() TestCommand
}

// LanguageCommandBuilder represents a language command builder
type LanguageCommandBuilder interface {
	Create() LanguageCommandBuilder
	WithVariable(variable string) LanguageCommandBuilder
	WithValues(values []LanguageValue) LanguageCommandBuilder
	Now() (LanguageCommand, error)
}

// LanguageCommand represents a language command
type LanguageCommand interface {
	Variable() string
	Values() []LanguageValue
}

// ScriptCommandBuilder represents a script command builder
type ScriptCommandBuilder interface {
	Create() ScriptCommandBuilder
	WithVariable(variable string) ScriptCommandBuilder
	WithValues(values []ScriptValue) ScriptCommandBuilder
	Now() (ScriptCommand, error)
}

// ScriptCommand represents a script command
type ScriptCommand interface {
	Variable() string
	Values() []ScriptValue
}

// HeadCommandBuilder represents an head command builder
type HeadCommandBuilder interface {
	Create() HeadCommandBuilder
	WithVariable(variable string) HeadCommandBuilder
	WithValues(values []HeadValue) HeadCommandBuilder
	Now() (HeadCommand, error)
}

// HeadCommand represents a head command
type HeadCommand interface {
	Variable() string
	Values() []HeadValue
}

// MainCommandBuilder represents a main command builder
type MainCommandBuilder interface {
	Create() MainCommandBuilder
	WithVariable(variable string) MainCommandBuilder
	WithInstructions(ins []MainCommandInstruction) MainCommandBuilder
	Now() (MainCommand, error)
}

// MainCommand represents a main command
type MainCommand interface {
	Variable() string
	Instructions() []MainCommandInstruction
}

// MainCommandInstructionBuilder represents a main command instruction builder
type MainCommandInstructionBuilder interface {
	Create() MainCommandInstructionBuilder
	WithInstruction(ins Instruction) MainCommandInstructionBuilder
	WithScopes(scopes Scopes) MainCommandInstructionBuilder
	Now() (MainCommandInstruction, error)
}

// MainCommandInstruction represents a main command instruction
type MainCommandInstruction interface {
	Instruction() Instruction
	HasScopes() bool
	Scopes() Scopes
}

// TestCommandBuilder represents a test command builder
type TestCommandBuilder interface {
	Create() TestCommandBuilder
	WithVariable(variable string) TestCommandBuilder
	WithName(name string) TestCommandBuilder
	WithInstructions(ins []TestCommandInstruction) TestCommandBuilder
	Now() (TestCommand, error)
}

// TestCommand represents a test command
type TestCommand interface {
	Variable() string
	Name() string
	Instructions() []TestCommandInstruction
}

// TestCommandInstructionBuilder represents a test command instruction builder
type TestCommandInstructionBuilder interface {
	Create() TestCommandInstructionBuilder
	WithInstruction(ins TestInstruction) TestCommandInstructionBuilder
	WithScopes(scopes Scopes) TestCommandInstructionBuilder
	Now() (TestCommandInstruction, error)
}

// TestCommandInstruction represents a test command instruction
type TestCommandInstruction interface {
	Instruction() TestInstruction
	HasScopes() bool
	Scopes() Scopes
}

// LabelCommandBuilder represents a label command builder
type LabelCommandBuilder interface {
	Create() LabelCommandBuilder
	WithVariable(variable string) LabelCommandBuilder
	WithName(name string) LabelCommandBuilder
	WithInstructions(ins []LabelCommandInstruction) LabelCommandBuilder
	Now() (LabelCommand, error)
}

// LabelCommand represents a label command
type LabelCommand interface {
	Variable() string
	Name() string
	Instructions() []LabelCommandInstruction
}

// LabelCommandInstructionBuilder represents a label command instruction builder
type LabelCommandInstructionBuilder interface {
	Create() LabelCommandInstructionBuilder
	WithInstruction(ins LabelInstruction) LabelCommandInstructionBuilder
	WithScopes(scopes Scopes) LabelCommandInstructionBuilder
	Now() (LabelCommandInstruction, error)
}

// LabelCommandInstruction represents a label command instruction
type LabelCommandInstruction interface {
	Instruction() LabelInstruction
	HasScopes() bool
	Scopes() Scopes
}

// ScopesBuilder represents a scopes builder
type ScopesBuilder interface {
	Create() ScopesBuilder
	WithScopes(scopes []Scope) ScopesBuilder
	Now() (Scopes, error)
}

// Scopes represents scopes
type Scopes interface {
	All() []Scope
}

// ScopeBuilder represents a scope builder
type ScopeBuilder interface {
	Create() ScopeBuilder
	IsInternal() ScopeBuilder
	IsExternal() ScopeBuilder
	Now() (Scope, error)
}

// Scope represents a scope
type Scope interface {
	IsInternal() bool
	IsExternal() bool
}

// LanguageDefinitionBuilder represents the language definition builder
type LanguageDefinitionBuilder interface {
	Create() LanguageDefinitionBuilder
	WithValues(values []LanguageValue) LanguageDefinitionBuilder
	Now() (LanguageDefinition, error)
}

// LanguageDefinition represents a language definition
type LanguageDefinition interface {
	Root() string
	Tokens() RelativePath
	Rules() RelativePath
	Logic() RelativePath
	Input() string
	HasChannels() bool
	Channels() RelativePath
	HasExtends() bool
	Extends() []RelativePath
	HasPatternMatches() bool
	PatternMatches() []PatternMatch
}

// LanguageValueBuilder represents the language value builder
type LanguageValueBuilder interface {
	Create() LanguageValueBuilder
	WithRoot(root string) LanguageValueBuilder
	WithTokens(tokens RelativePath) LanguageValueBuilder
	WithChannels(channels RelativePath) LanguageValueBuilder
	WithRules(rules RelativePath) LanguageValueBuilder
	WithLogic(logic RelativePath) LanguageValueBuilder
	WithInputVariable(inputVar string) LanguageValueBuilder
	WithExtends(extends []RelativePath) LanguageValueBuilder
	WithPatternMatches(matches []PatternMatch) LanguageValueBuilder
	Now() (LanguageValue, error)
}

// LanguageValue represents a language value
type LanguageValue interface {
	IsRoot() bool
	Root() string
	IsTokens() bool
	Tokens() RelativePath
	IsChannels() bool
	Channels() RelativePath
	IsRules() bool
	Rules() RelativePath
	IsLogic() bool
	Logic() RelativePath
	IsInputVariable() bool
	InputVariable() string
	IsExtends() bool
	Extends() []RelativePath
	IsPatternMatches() bool
	PatternMatches() []PatternMatch
}

// PatternMatchBuilder represents a pattern match builderexitHeadValue
type PatternMatchBuilder interface {
	Create() PatternMatchBuilder
	WithPattern(pattern string) PatternMatchBuilder
	WithLabels(labels PatternLabels) PatternMatchBuilder
	Now() (PatternMatch, error)
}

// PatternMatch represents pattern match
type PatternMatch interface {
	Pattern() string
	Labels() PatternLabels
}

// PatternLabelsBuilder represents the pattern labels builder
type PatternLabelsBuilder interface {
	Create() PatternLabelsBuilder
	WithEnterLabel(enter string) PatternLabelsBuilder
	WithExitLabel(exit string) PatternLabelsBuilder
	Now() (PatternLabels, error)
}

// PatternLabels represents the pattern labels
type PatternLabels interface {
	HasEnterLabel() bool
	EnterLabel() string
	HasExitLabel() bool
	ExitLabel() string
}

// RelativePathsBuilder represents a relative paths builder
type RelativePathsBuilder interface {
	Create() RelativePathsBuilder
	WithRelativePaths(relPaths []RelativePath) RelativePathsBuilder
	Now() (RelativePaths, error)
}

// RelativePaths represents relative paths
type RelativePaths interface {
	All() []RelativePath
}

// RelativePathBuilder represents the relativePath builder
type RelativePathBuilder interface {
	Create() RelativePathBuilder
	WithSections(sections []FolderSection) RelativePathBuilder
	WithPath(path string) RelativePathBuilder
	Now() (RelativePath, error)
}

// RelativePath represents a relative path
type RelativePath interface {
	All() []FolderSection
	Head() []FolderSection
	HasTail() bool
	Tail() FolderSection
	String() string
}

// FolderSectionBuilder represents a folder section builder
type FolderSectionBuilder interface {
	Create() FolderSectionBuilder
	IsTail() FolderSectionBuilder
	WithName(name FolderName) FolderSectionBuilder
	Now() (FolderSection, error)
}

// FolderSection represents a folder section
type FolderSection interface {
	IsTail() bool
	Name() FolderName
	String() string
}

// FolderNameBuilder represents a folderName builder
type FolderNameBuilder interface {
	Create() FolderNameBuilder
	IsCurrent() FolderNameBuilder
	IsPrevious() FolderNameBuilder
	WithName(name string) FolderNameBuilder
	Now() (FolderName, error)
}

// FolderName represents a folder name
type FolderName interface {
	IsCurrent() bool
	IsPrevious() bool
	IsName() bool
	Name() string
	String() string
}

// ApplicationBuilder represents the application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithHead(head HeadSection) ApplicationBuilder
	WithLabel(label LabelSection) ApplicationBuilder
	WithMain(main MainSection) ApplicationBuilder
	WithTest(test TestSection) ApplicationBuilder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Head() HeadSection
	Main() MainSection
	HasLabel() bool
	Label() LabelSection
	HasTest() bool
	Test() TestSection
}

// TestSectionBuilder represents a testSection builder
type TestSectionBuilder interface {
	Create() TestSectionBuilder
	WithDeclarations(declarations []TestDeclaration) TestSectionBuilder
	Now() (TestSection, error)
}

// TestSection represents a test section
type TestSection interface {
	Declarations() []TestDeclaration
}

// TestDeclarationBuilder represnets a testDeclaration builder
type TestDeclarationBuilder interface {
	Create() TestDeclarationBuilder
	WithName(name string) TestDeclarationBuilder
	WithInstructions(instructions []TestInstruction) TestDeclarationBuilder
	Now() (TestDeclaration, error)
}

// TestDeclaration represents a test declaration
type TestDeclaration interface {
	Name() string
	Instructions() []TestInstruction
}

// TestInstructionBuilder represents a testInstruction builder
type TestInstructionBuilder interface {
	Create() TestInstructionBuilder
	WithInstruction(ins Instruction) TestInstructionBuilder
	WithReadFile(readFile ReadFile) TestInstructionBuilder
	WithAssert(assert Assert) TestInstructionBuilder
	Now() (TestInstruction, error)
}

// TestInstruction represents a test instruction
type TestInstruction interface {
	IsAssert() bool
	Assert() Assert
	IsReadFile() bool
	ReadFile() ReadFile
	IsInstruction() bool
	Instruction() Instruction
}

// AssertBuilder represents an assert builder
type AssertBuilder interface {
	Create() AssertBuilder
	WithIndex(index int) AssertBuilder
	WithCondition(condition string) AssertBuilder
	Now() (Assert, error)
}

// Assert represents an assert
type Assert interface {
	Index() int
	HasCondition() bool
	Condition() string
}

// ReadFileBuilder represents a readfile builder
type ReadFileBuilder interface {
	Create() ReadFileBuilder
	WithVariable(variable string) ReadFileBuilder
	WithPath(path RelativePath) ReadFileBuilder
	Now() (ReadFile, error)
}

// ReadFile represents a reafile instruction
type ReadFile interface {
	Variable() string
	Path() RelativePath
}

// HeadSectionBuilder represents the headSection builder
type HeadSectionBuilder interface {
	Create() HeadSectionBuilder
	WithValues(values []HeadValue) HeadSectionBuilder
	Now() (HeadSection, error)
}

// HeadSection represents the headSection
type HeadSection interface {
	Name() string
	Version() string
	HasImport() bool
	Import() []ImportSingle
}

// HeadValueBuilder represents the headValue builder
type HeadValueBuilder interface {
	Create() HeadValueBuilder
	WithName(name string) HeadValueBuilder
	WithVersion(version string) HeadValueBuilder
	WithImport(imp []ImportSingle) HeadValueBuilder
	Now() (HeadValue, error)
}

// HeadValue represents the head value
type HeadValue interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsImport() bool
	Import() []ImportSingle
}

// LabelSectionBuilder represents the labelSection builder
type LabelSectionBuilder interface {
	Create() LabelSectionBuilder
	WithDeclarations(decl []LabelDeclaration) LabelSectionBuilder
	Now() (LabelSection, error)
}

// LabelSection represents the label section
type LabelSection interface {
	Declarations() []LabelDeclaration
}

// LabelDeclarationBuilder represents the labelDeclaration builder
type LabelDeclarationBuilder interface {
	Create() LabelDeclarationBuilder
	WithName(name string) LabelDeclarationBuilder
	WithInstructions(ins []LabelInstruction) LabelDeclarationBuilder
	Now() (LabelDeclaration, error)
}

// LabelDeclaration represents the label declaration
type LabelDeclaration interface {
	Name() string
	Instructions() []LabelInstruction
}

// LabelInstructionBuilder represents the label instruction builder
type LabelInstructionBuilder interface {
	Create() LabelInstructionBuilder
	IsRet() LabelInstructionBuilder
	WithInstruction(ins Instruction) LabelInstructionBuilder
	Now() (LabelInstruction, error)
}

// LabelInstruction represents the label instruction
type LabelInstruction interface {
	IsRet() bool
	IsInstruction() bool
	Instruction() Instruction
}

// MainSectionBuilder represents the main section builder
type MainSectionBuilder interface {
	Create() MainSectionBuilder
	WithInstructions(ins []Instruction) MainSectionBuilder
	Now() (MainSection, error)
}

// MainSection represents the main section
type MainSection interface {
	Instructions() []Instruction
}

// ImportSingleBuilder represents an import single builder
type ImportSingleBuilder interface {
	Create() ImportSingleBuilder
	WithName(name string) ImportSingleBuilder
	WithPath(path RelativePath) ImportSingleBuilder
	Now() (ImportSingle, error)
}

// ImportSingle represents an import single
type ImportSingle interface {
	Name() string
	Path() RelativePath
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithVariable(variable Variable) InstructionBuilder
	WithOperation(operation Operation) InstructionBuilder
	WithPrint(print Print) InstructionBuilder
	WithStackFrame(stackFrame StackFrame) InstructionBuilder
	WithJump(jmp Jump) InstructionBuilder
	WithExit(exit Exit) InstructionBuilder
	WithCall(call Call) InstructionBuilder
	WithRegistry(registry Registry) InstructionBuilder
	WithSwitch(swtch Switch) InstructionBuilder
	WithSave(save Save) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsVariable() bool
	Variable() Variable
	IsOperation() bool
	Operation() Operation
	IsPrint() bool
	Print() Print
	IsStackFrame() bool
	StackFrame() StackFrame
	IsJump() bool
	Jump() Jump
	IsExit() bool
	Exit() Exit
	IsCall() bool
	Call() Call
	IsRegistry() bool
	Registry() Registry
	IsSwitch() bool
	Switch() Switch
	IsSave() bool
	Save() Save
}

// RegistryBuilder represents a registry builder
type RegistryBuilder interface {
	Create() RegistryBuilder
	WithFetch(fetch FetchRegistry) RegistryBuilder
	WithRegister(register Register) RegistryBuilder
	WithUnregister(unregister Unregister) RegistryBuilder
	Now() (Registry, error)
}

// Registry represents a registry instruction
type Registry interface {
	IsFetch() bool
	Fetch() FetchRegistry
	IsRegister() bool
	Register() Register
	IsUnregister() bool
	Unregister() Unregister
}

// FetchRegistryBuilder represents a fetch registry builder
type FetchRegistryBuilder interface {
	Create() FetchRegistryBuilder
	From(from string) FetchRegistryBuilder
	To(to string) FetchRegistryBuilder
	WithIndex(index IntPointer) FetchRegistryBuilder
	Now() (FetchRegistry, error)
}

// FetchRegistry represents a fetch registry instruction
type FetchRegistry interface {
	To() string
	From() string
	HasIndex() bool
	Index() IntPointer
}

// UnregisterBuilder represents an unregister builder
type UnregisterBuilder interface {
	Create() UnregisterBuilder
	WithVariable(name string) UnregisterBuilder
	Now() (Unregister, error)
}

// Unregister represents an unregister instruction
type Unregister interface {
	Variable() string
}

// RegisterBuilder represents a register pointer
type RegisterBuilder interface {
	Create() RegisterBuilder
	WithVariable(variable string) RegisterBuilder
	WithIndex(Index IntPointer) RegisterBuilder
	Now() (Register, error)
}

// Register represents a register instruction
type Register interface {
	Variable() string
	HasIndex() bool
	Index() IntPointer
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithCodeMatch(codeMatch CodeMatch) TokenBuilder
	WithTokenSection(tokenSection TokenSection) TokenBuilder
	Now() (Token, error)
}

// Token represents a token instruction
type Token interface {
	IsCodeMatch() bool
	CodeMatch() CodeMatch
	IsTokenSection() bool
	TokenSection() TokenSection
}

// CodeMatchBuilder represents a codeMatch builder
type CodeMatchBuilder interface {
	Create() CodeMatchBuilder
	WithContent(content string) CodeMatchBuilder
	WithSection(section string) CodeMatchBuilder
	WithPatternVariables(patterns []string) CodeMatchBuilder
	Now() (CodeMatch, error)
}

// CodeMatch represents code match
type CodeMatch interface {
	Content() string
	Section() string
	PatternVariables() []string
}

// TokenSectionBuilder represents a tokenSection builder
type TokenSectionBuilder interface {
	Create() TokenSectionBuilder
	WithVariableName(variableName string) TokenSectionBuilder
	WithSpecific(specific SpecificTokenCode) TokenSectionBuilder
	Now() (TokenSection, error)
}

// TokenSection represents a token section
type TokenSection interface {
	IsVariableName() bool
	VariableName() string
	IsSpecific() bool
	Specific() SpecificTokenCode
}

// SpecificTokenCodeBuilder represents a specificTokenCode builder
type SpecificTokenCodeBuilder interface {
	Create() SpecificTokenCodeBuilder
	WithVariableName(variableName string) SpecificTokenCodeBuilder
	WithAmount(amount string) SpecificTokenCodeBuilder
	WithPatternVariable(pattern string) SpecificTokenCodeBuilder
	Now() (SpecificTokenCode, error)
}

// SpecificTokenCode represents a specific token code
type SpecificTokenCode interface {
	VariableName() string
	PatternVariable() string
	HasAmount() bool
	Amount() string
}

// MatchBuilder represents a match builder
type MatchBuilder interface {
	Create() MatchBuilder
	WithInput(input string) MatchBuilder
	WithPattern(pattern string) MatchBuilder
	Now() (Match, error)
}

// Match represents a match
type Match interface {
	Input() string
	HasPattern() bool
	Pattern() string
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithDeclaration(declaration Declaration) VariableBuilder
	WithAssigment(assignment Assignment) VariableBuilder
	WithConcatenation(concatenation Concatenation) VariableBuilder
	WithDelete(delete string) VariableBuilder
	Now() (Variable, error)
}

// Variable represents a variable related instruction
type Variable interface {
	IsDeclaration() bool
	Declaration() Declaration
	IsAssignment() bool
	Assignment() Assignment
	IsConcatenation() bool
	Concatenation() Concatenation
	IsDelete() bool
	Delete() string
}

// ConcatenationBuilder represents a concatenation builder
type ConcatenationBuilder interface {
	Create() ConcatenationBuilder
	WithOperation(operation StandardOperation) ConcatenationBuilder
	Now() (Concatenation, error)
}

// Concatenation represents a concatenation
type Concatenation interface {
	Operation() StandardOperation
}

// DeclarationBuilder represents a declaration builder
type DeclarationBuilder interface {
	Create() DeclarationBuilder
	WithVariable(name string) DeclarationBuilder
	WithType(typ Type) DeclarationBuilder
	Now() (Declaration, error)
}

// Declaration represents a variable declaration instruction
type Declaration interface {
	Variable() string
	Type() Type
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithVariable(variable string) AssignmentBuilder
	WithValue(value ValueRepresentation) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents a variable assignment instruction
type Assignment interface {
	Variable() string
	Value() ValueRepresentation
}

// ValueRepresentationBuilder represents a value representation builder
type ValueRepresentationBuilder interface {
	Create() ValueRepresentationBuilder
	WithValue(value Value) ValueRepresentationBuilder
	WithVariable(variable string) ValueRepresentationBuilder
	Now() (ValueRepresentation, error)
}

// ValueRepresentation represents a value representation
type ValueRepresentation interface {
	IsValue() bool
	Value() Value
	IsVariable() bool
	Variable() string
}

// ValueBuilder represents a value
type ValueBuilder interface {
	Create() ValueBuilder
	IsNil() ValueBuilder
	WithNumeric(numeric NumericValue) ValueBuilder
	WithBool(bl bool) ValueBuilder
	WithString(str string) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsNil() bool
	IsNumeric() bool
	Numeric() NumericValue
	IsBool() bool
	Bool() *bool
	IsString() bool
	String() string
}

// NumericValueBuilder represents a numeric value builder
type NumericValueBuilder interface {
	Create() NumericValueBuilder
	IsNegative() NumericValueBuilder
	WithInt(intVal int) NumericValueBuilder
	WithFloat(floatVal float64) NumericValueBuilder
	Now() (NumericValue, error)
}

// NumericValue represents a numeric value
type NumericValue interface {
	IsNegative() bool
	IsInt() bool
	Int() *int
	IsFloat() bool
	Float() *float64
}

// TypeBuilder represents a type builder
type TypeBuilder interface {
	Create() TypeBuilder
	IsBool() TypeBuilder
	IsInt8() TypeBuilder
	IsInt16() TypeBuilder
	IsInt32() TypeBuilder
	IsInt64() TypeBuilder
	IsFloat32() TypeBuilder
	IsFloat64() TypeBuilder
	IsUint8() TypeBuilder
	IsUint16() TypeBuilder
	IsUint32() TypeBuilder
	IsUint64() TypeBuilder
	IsString() TypeBuilder
	IsStackFrame() TypeBuilder
	Now() (Type, error)
}

// Type represents a variable type
type Type interface {
	IsBool() bool
	IsInt8() bool
	IsInt16() bool
	IsInt32() bool
	IsInt64() bool
	IsFloat32() bool
	IsFloat64() bool
	IsUint8() bool
	IsUint16() bool
	IsUint32() bool
	IsUint64() bool
	IsString() bool
	IsStackFrame() bool
	String() string
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithArythmetic(arythmetic Arythmetic) OperationBuilder
	WithRelational(relational Relational) OperationBuilder
	WithLogical(logical Logical) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation instruction
type Operation interface {
	IsArythmetic() bool
	Arythmetic() Arythmetic
	IsRelational() bool
	Relational() Relational
	IsLogical() bool
	Logical() Logical
}

// ArythmeticBuilder represents an arythmetic builder
type ArythmeticBuilder interface {
	Create() ArythmeticBuilder
	WithAddition(add StandardOperation) ArythmeticBuilder
	WithSubstraction(sub StandardOperation) ArythmeticBuilder
	WithMultiplication(mul StandardOperation) ArythmeticBuilder
	WithDivision(div RemainingOperation) ArythmeticBuilder
	Now() (Arythmetic, error)
}

// Arythmetic represents an arythmetic operation instruction
type Arythmetic interface {
	IsAdd() bool
	Add() StandardOperation
	IsSub() bool
	Sub() StandardOperation
	IsMul() bool
	Mul() StandardOperation
	IsDiv() bool
	Div() RemainingOperation
}

// RelationalBuilder represents a relational builder
type RelationalBuilder interface {
	Create() RelationalBuilder
	WithLessThan(lessThan StandardOperation) RelationalBuilder
	WithEqual(equal StandardOperation) RelationalBuilder
	WithNotEqual(notEqual StandardOperation) RelationalBuilder
	Now() (Relational, error)
}

// Relational represents a relational operation instruction
type Relational interface {
	IsLessThan() bool
	LessThan() StandardOperation
	IsEqual() bool
	Equal() StandardOperation
	IsNotEqual() bool
	NotEqual() StandardOperation
}

// LogicalBuilder represents a logical builder
type LogicalBuilder interface {
	Create() LogicalBuilder
	WithAnd(and StandardOperation) LogicalBuilder
	WithOr(or StandardOperation) LogicalBuilder
	Now() (Logical, error)
}

// Logical represents a logical operation instruction
type Logical interface {
	IsAnd() bool
	And() StandardOperation
	IsOr() bool
	Or() StandardOperation
}

// StandardOperationBuilder represents a standard operation builder
type StandardOperationBuilder interface {
	Create() StandardOperationBuilder
	WithFirst(first string) StandardOperationBuilder
	WithSecond(second string) StandardOperationBuilder
	WithResult(result string) StandardOperationBuilder
	Now() (StandardOperation, error)
}

// StandardOperation represents a standard operation
type StandardOperation interface {
	First() string
	Second() string
	Result() string
}

// RemainingOperationBuilder represents a remaining operation builder
type RemainingOperationBuilder interface {
	Create() RemainingOperationBuilder
	WithFirst(first string) RemainingOperationBuilder
	WithSecond(second string) RemainingOperationBuilder
	WithResult(result string) RemainingOperationBuilder
	WithRemaining(remaining string) RemainingOperationBuilder
	Now() (RemainingOperation, error)
}

// RemainingOperation represents a an operation with a remaining value
type RemainingOperation interface {
	First() string
	Second() string
	Result() string
	Remaining() string
}

// PrintBuilder represents a print instruction builder
type PrintBuilder interface {
	Create() PrintBuilder
	WithValue(value ValueRepresentation) PrintBuilder
	Now() (Print, error)
}

// Print represents a print instruction
type Print interface {
	Value() ValueRepresentation
}

// JumpBuilder represents a jump builder
type JumpBuilder interface {
	Create() JumpBuilder
	WithLabel(label string) JumpBuilder
	WithCondition(condition string) JumpBuilder
	Now() (Jump, error)
}

// Jump represents a jump instruction
type Jump interface {
	Label() string
	HasCondition() bool
	Condition() string
}

// ExitBuilder represents an exit builder
type ExitBuilder interface {
	Create() ExitBuilder
	WithCondition(cond string) ExitBuilder
	Now() (Exit, error)
}

// Exit represents an exit instruction
type Exit interface {
	HasCondition() bool
	Condition() string
}

// CallBuilder represents a call builder
type CallBuilder interface {
	Create() CallBuilder
	WithName(name string) CallBuilder
	WithStackFrame(stackFrame string) CallBuilder
	WithCondition(condition string) CallBuilder
	Now() (Call, error)
}

// Call represents a call instruction
type Call interface {
	Name() string
	StackFrame() string
	HasCondition() bool
	Condition() string
}

// SaveBuilder represents a save builder
type SaveBuilder interface {
	Create() SaveBuilder
	From(from string) SaveBuilder
	To(to string) SaveBuilder
	Now() (Save, error)
}

// Save represents a save instruction
type Save interface {
	To() string
	HasFrom() bool
	From() string
}

// SwitchBuilder represents a switch builder
type SwitchBuilder interface {
	Create() SwitchBuilder
	WithVariable(variable string) SwitchBuilder
	Now() (Switch, error)
}

// Switch represents a switch instruction
type Switch interface {
	Variable() string
}

// StackFrameBuilder represents a stackFrame builder
type StackFrameBuilder interface {
	Create() StackFrameBuilder
	IsPush() StackFrameBuilder
	IsPop() StackFrameBuilder
	WithIndex(index Index) StackFrameBuilder
	WithSkip(skip Skip) StackFrameBuilder
	Now() (StackFrame, error)
}

// StackFrame represents a stackFrame related instruction
type StackFrame interface {
	IsPush() bool
	IsPop() bool
	IsIndex() bool
	Index() Index
	IsSkip() bool
	Skip() Skip
}

// IndexBuilder represents an index builder
type IndexBuilder interface {
	Create() IndexBuilder
	WithVariable(variable string) IndexBuilder
	Now() (Index, error)
}

// Index represents an index
type Index interface {
	Variable() string
}

// SkipBuilder represents a skip builder
type SkipBuilder interface {
	Create() SkipBuilder
	WithPointer(pointer IntPointer) SkipBuilder
	Now() (Skip, error)
}

// Skip represents a skip
type Skip interface {
	Pointer() IntPointer
}

// IntPointerBuilder represents an int pointer builder
type IntPointerBuilder interface {
	Create() IntPointerBuilder
	WithInt(intVal int64) IntPointerBuilder
	WithVariable(variable string) IntPointerBuilder
	Now() (IntPointer, error)
}

// IntPointer represents an int pointer
type IntPointer interface {
	IsInt() bool
	Int() int64
	IsVariable() bool
	Variable() string
}
