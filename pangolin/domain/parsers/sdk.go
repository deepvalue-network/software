package parsers

import (
	"strings"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lparser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
)

const quaotationChar = '"'
const scriptHeadDelimiter = "----"

// NewParserBuilder creates a new ParserBuilder instance
func NewParserBuilder() ParserBuilder {
	application := lparser.NewApplication()
	parserBuilder := lparser.NewBuilder()
	lexerBuilder := lexers.NewBuilder()
	programBuilder := createProgramBuilder()
	languageBuilder := createLanguageBuilder()
	scopesBuilder := createScopesBuilder()
	scopeBuilder := createScopeBuilder()
	commandBuilder := createCommandBuilder()
	languageCommandBuilder := createLanguageCommandBuilder()
	scriptCommandBuilder := createScriptCommandBuilder()
	headCommandBuilder := createHeadCommandBuilder()
	mainCommandBuilder := createMainCommandBuilder()
	mainCommandInstructionBuilder := createMainCommandInstructionBuilder()
	testCommandBuilder := createTestCommandBuilder()
	testCommandInstructionBuilder := createTestCommandInstructionBuilder()
	labelCommandBuilder := createLabelCommandBuilder()
	labelCommandInstructionBuilder := createLabelCommandInstructionBuilder()
	languageApplicationBuilder := createLanguageApplicationBuilder()
	languageMainSectionBuilder := createLanguageMainSectionBuilder()
	languageTestSectionBuilder := createLanguageTestSectionBuilder()
	languageTestDeclarationBuilder := createLanguageTestDeclarationBuilder()
	languageTestInstructionBuilder := createLanguageTestInstructionBuilder()
	languageLabelSectionBuilder := createLanguageLabelSectionBuilder()
	languageLabelDeclarationBuilder := createLanguageLabelDeclarationBuilder()
	languageLabelInstructionBuilder := createLanguageLabelInstructionBuilder()
	languageInstructionBuilder := createLanguageInstructionBuilder()
	languageDefinitionBuilder := createLanguageDefinitionBuilder()
	languageValueBuilder := createLanguageValueBuilder()
	scriptBuilder := createScriptBuilder()
	scriptValueBuilder := createScriptValueBuilder()
	patternMatchBuilder := createPatternMatchBuilder()
	patternLabelsBuilder := createPatternLabelsBuilder()
	relativePathsBuilder := createRelativePathsBuilder()
	relativePathBuilder := createRelativePathBuilder()
	folderSectionBuilder := createFolderSectionBuilder()
	folderNameBuilder := createFolderNameBuilder()
	applicationBuilder := createApplicationBuilder()
	testSectionBuilder := createTestSectionBuilder()
	testDeclarationBuilder := createTestDeclarationBuilder()
	testInstructionBuilder := createTestInstructionBuilder()
	assertBuilder := createAssertBuilder()
	readFileBuilder := createReadFileBuilder()
	headSectionBuilder := createHeadSectionBuilder()
	headValueBuilder := createHeadValueBuilder()
	importSingleBuilder := createImportSingleBuilder()
	labelSectionBuilder := createLabelSectionBuilder()
	labelDeclarationBuilder := createLabelDeclarationBuilder()
	labelInstructionBuilder := createLabelInstructionBuilder()
	mainSectionBuilder := createMainSectionBuilder()
	instructionBuilder := createInstructionBuilder()
	specificTokenCodeBuilder := createSpecificTokenCodeBuilder()
	tokenSectionBuilder := createTokenSectionBuilder()
	codeMatchBuilder := createCodeMatchBuilder()
	tokenBuilder := createTokenBuilder()
	variableBuilder := createVariableBuilder()
	concatenationBuilder := createConcatenationBuilder()
	declarationBuilder := createDeclarationBuilder()
	assignmentBuilder := createAssignmentBuilder()
	valueRepresentationBuilder := createValueRepresentationBuilder()
	valueBuilder := createValueBuilder()
	numericValueBuilder := createNumericValueBuilder()
	typeBuilder := createTypeBuilder()
	operationBuilder := createOperationalBuilder()
	arythmeticBuilder := createArythmeticBuilder()
	relationalBuilder := createRelationalBuilder()
	logicalBuilder := createLogicalBuilder()
	transformOperationBuilder := createTransformOperationBuilder()
	standardOperationBuilder := createStandardOperationBuilder()
	remainingOperationBuilder := createRemainingOperationBuilder()
	printBuilder := createPrintBuilder()
	jumpBuilder := createJumpBuilder()
	matchBuilder := createMatchBuilder()
	exitBuilder := createExitBuilder()
	callBuilder := createCallBuilder()
	stackFrameBuilder := createStackFrameBuilder()
	indexBuilder := createIndexBuilder()
	skipBuilder := createSkipBuilder()
	intPointerBuilder := createIntPointerBuilder()

	return createParserBuilder(
		application,
		parserBuilder,
		lexerBuilder,
		programBuilder,
		languageBuilder,
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
		languageDefinitionBuilder,
		languageValueBuilder,
		scriptBuilder,
		scriptValueBuilder,
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
		transformOperationBuilder,
		standardOperationBuilder,
		remainingOperationBuilder,
		printBuilder,
		jumpBuilder,
		matchBuilder,
		exitBuilder,
		callBuilder,
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
	WithApplication(app Application) ProgramBuilder
	WithLanguage(lang Language) ProgramBuilder
	WithScript(script Script) ProgramBuilder
	Now() (Program, error)
}

// Program represents the program
type Program interface {
	IsApplication() bool
	Application() Application
	IsLanguage() bool
	Language() Language
	IsScript() bool
	Script() Script
}

// LanguageBuilder represents a language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithApplication(application LanguageApplication) LanguageBuilder
	WithDefinition(definition LanguageDefinition) LanguageBuilder
	Now() (Language, error)
}

// Language represents a language
type Language interface {
	IsApplication() bool
	Application() LanguageApplication
	IsDefinition() bool
	Definition() LanguageDefinition
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
}

// ScriptValueBuilder represents a script value builder
type ScriptValueBuilder interface {
	Create() ScriptValueBuilder
	WithName(name string) ScriptValueBuilder
	WithVersion(version string) ScriptValueBuilder
	WithScriptPath(scriptPath RelativePath) ScriptValueBuilder
	WithLanguagePath(langPath RelativePath) ScriptValueBuilder
	WithOutput(output string) ScriptValueBuilder
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
	WithLanguageInstruction(languageIns LanguageInstruction) LanguageTestInstructionBuilder
	WithTestInstruction(testIns TestInstruction) LanguageTestInstructionBuilder
	Now() (LanguageTestInstruction, error)
}

// LanguageTestInstruction represents a language test instruction
type LanguageTestInstruction interface {
	IsLanguageInstruction() bool
	LanguageInstruction() LanguageInstruction
	IsTestInstruction() bool
	TestInstruction() TestInstruction
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
	WithInstruction(ins Instruction) LanguageInstructionBuilder
	WithMatch(match Match) LanguageInstructionBuilder
	WithCommand(command Command) LanguageInstructionBuilder
	Now() (LanguageInstruction, error)
}

// LanguageInstruction represents a language instruction
type LanguageInstruction interface {
	IsInstruction() bool
	Instruction() Instruction
	IsMatch() bool
	Match() Match
	IsCommand() bool
	Command() Command
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

// TransformOperationBuilder represents a transform operation builder
type TransformOperationBuilder interface {
	Create() TransformOperationBuilder
	WithInput(input string) TransformOperationBuilder
	WithResult(result string) TransformOperationBuilder
	Now() (TransformOperation, error)
}

// TransformOperation represents a transform operation
type TransformOperation interface {
	Input() string
	Result() string
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
	WithCondition(condition string) CallBuilder
	Now() (Call, error)
}

// Call represents a call instruction
type Call interface {
	Name() string
	HasCondition() bool
	Condition() string
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
