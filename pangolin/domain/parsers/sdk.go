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
	languageValueBuilder := createLanguageValueBuilder()
	targetBuilder := createTargetBuilder()
	targetSingleBuilder := createTargetSingleBuilder()
	eventBuilder := createEventBuilder()
	scriptBuilder := createScriptBuilder()
	scriptValueBuilder := createScriptValueBuilder()
	patternMatchBuilder := createPatternMatchBuilder()
	patternLabelsBuilder := createPatternLabelsBuilder()
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
	triggerBuilder := createTriggerBuilder()
	formatBuilder := createFormatBuilder()
	specificTokenCodeBuilder := createSpecificTokenCodeBuilder()
	tokenSectionBuilder := createTokenSectionBuilder()
	codeMatchBuilder := createCodeMatchBuilder()
	tokenBuilder := createTokenBuilder()
	variableBuilder := createVariableBuilder()
	concatenationBuilder := createConcatenationBuilder()
	declarationBuilder := createDeclarationBuilder()
	assignmentBuilder := createAssignmentBuilder()
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
	pushBuilder := createPushBuilder()
	popBuilder := createPopBuilder()
	identifierBuilder := createIdentifierBuilder()
	variableNameBuilder := createVariableNameBuilder()

	return createParserBuilder(
		application,
		parserBuilder,
		lexerBuilder,
		programBuilder,
		languageBuilder,
		languageValueBuilder,
		targetBuilder,
		targetSingleBuilder,
		eventBuilder,
		scriptBuilder,
		scriptValueBuilder,
		patternMatchBuilder,
		patternLabelsBuilder,
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
		triggerBuilder,
		formatBuilder,
		specificTokenCodeBuilder,
		tokenSectionBuilder,
		codeMatchBuilder,
		tokenBuilder,
		variableBuilder,
		concatenationBuilder,
		declarationBuilder,
		assignmentBuilder,
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
		pushBuilder,
		popBuilder,
		identifierBuilder,
		variableNameBuilder,
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
}

// ScriptValueBuilder represents a script value builder
type ScriptValueBuilder interface {
	Create() ScriptValueBuilder
	WithName(name string) ScriptValueBuilder
	WithVersion(version string) ScriptValueBuilder
	WithScriptPath(scriptPath RelativePath) ScriptValueBuilder
	WithLanguagePath(langPath RelativePath) ScriptValueBuilder
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
}

// LanguageBuilder represents the language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithValues(values []LanguageValue) LanguageBuilder
	Now() (Language, error)
}

// Language represents a language
type Language interface {
	Root() string
	Tokens() RelativePath
	Rules() RelativePath
	Logic() RelativePath
	Input() string
	Output() string
	Targets() []Target
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
	WithOutputVariable(outputVar string) LanguageValueBuilder
	WithExtends(extends []RelativePath) LanguageValueBuilder
	WithPatternMatches(matches []PatternMatch) LanguageValueBuilder
	WithTargets(targets []Target) LanguageValueBuilder
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
	IsOutputVariable() bool
	OutputVariable() string
	IsExtends() bool
	Extends() []RelativePath
	IsPatternMatches() bool
	PatternMatches() []PatternMatch
	IsTargets() bool
	Targets() []Target
}

// TargetBuilder represents a target builder
type TargetBuilder interface {
	Create() TargetBuilder
	WithName(name string) TargetBuilder
	WithSingles(singles []TargetSingle) TargetBuilder
	Now() (Target, error)
}

// Target represents a target
type Target interface {
	Name() string
	Path() RelativePath
	Events() []Event
}

// TargetSingleBuilder represents a target single builder
type TargetSingleBuilder interface {
	Create() TargetSingleBuilder
	WithEvents(evts []Event) TargetSingleBuilder
	WithPath(path RelativePath) TargetSingleBuilder
	Now() (TargetSingle, error)
}

// TargetSingle represents a target single
type TargetSingle interface {
	IsEvents() bool
	Events() []Event
	IsPath() bool
	Path() RelativePath
}

// EventBuilder represents an event builder
type EventBuilder interface {
	Create() EventBuilder
	WithName(name string) EventBuilder
	WithLabel(label string) EventBuilder
	Now() (Event, error)
}

// Event represents an event
type Event interface {
	Name() string
	Label() string
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
	WithCondition(condition Identifier) AssertBuilder
	Now() (Assert, error)
}

// Assert represents an assert
type Assert interface {
	HasCondition() bool
	Condition() Identifier
}

// ReadFileBuilder represents a readfile builder
type ReadFileBuilder interface {
	Create() ReadFileBuilder
	WithVariable(variable VariableName) ReadFileBuilder
	WithPath(path RelativePath) ReadFileBuilder
	Now() (ReadFile, error)
}

// ReadFile represents a reafile instruction
type ReadFile interface {
	Variable() VariableName
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
	WithMatch(match Match) InstructionBuilder
	WithExit(exit Exit) InstructionBuilder
	WithCall(call Call) InstructionBuilder
	WithToken(token Token) InstructionBuilder
	WithTrigger(trigger Trigger) InstructionBuilder
	WithFormat(format Format) InstructionBuilder
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
	IsMatch() bool
	Match() Match
	IsExit() bool
	Exit() Exit
	IsCall() bool
	Call() Call
	IsToken() bool
	Token() Token
	IsTrigger() bool
	Trigger() Trigger
	IsFormat() bool
	Format() Format
}

// FormatBuilder represents a format builder
type FormatBuilder interface {
	Create() FormatBuilder
	WithResults(results VariableName) FormatBuilder
	WithPattern(pattern Identifier) FormatBuilder
	WithFirst(first Identifier) FormatBuilder
	WithSecond(second Identifier) FormatBuilder
	Now() (Format, error)
}

// Format represents a format
type Format interface {
	Results() VariableName
	Pattern() Identifier
	First() Identifier
	Second() Identifier
}

// TriggerBuilder represents a trigger builder
type TriggerBuilder interface {
	Create() TriggerBuilder
	WithVariableName(variableName VariableName) TriggerBuilder
	WithEvent(event string) TriggerBuilder
	Now() (Trigger, error)
}

// Trigger represents a trigger
type Trigger interface {
	Variable() VariableName
	Event() string
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
	WithContent(content VariableName) CodeMatchBuilder
	WithSection(section VariableName) CodeMatchBuilder
	WithPatternVariables(patterns []string) CodeMatchBuilder
	Now() (CodeMatch, error)
}

// CodeMatch represents code match
type CodeMatch interface {
	Content() VariableName
	Section() VariableName
	PatternVariables() []string
}

// TokenSectionBuilder represents a tokenSection builder
type TokenSectionBuilder interface {
	Create() TokenSectionBuilder
	WithVariableName(variableName VariableName) TokenSectionBuilder
	WithSpecific(specific SpecificTokenCode) TokenSectionBuilder
	Now() (TokenSection, error)
}

// TokenSection represents a token section
type TokenSection interface {
	IsVariableName() bool
	VariableName() VariableName
	IsSpecific() bool
	Specific() SpecificTokenCode
}

// SpecificTokenCodeBuilder represents a specificTokenCode builder
type SpecificTokenCodeBuilder interface {
	Create() SpecificTokenCodeBuilder
	WithVariableName(variableName VariableName) SpecificTokenCodeBuilder
	WithAmount(amount VariableName) SpecificTokenCodeBuilder
	WithPatternVariable(pattern string) SpecificTokenCodeBuilder
	Now() (SpecificTokenCode, error)
}

// SpecificTokenCode represents a specific token code
type SpecificTokenCode interface {
	VariableName() VariableName
	PatternVariable() string
	HasAmount() bool
	Amount() VariableName
}

// MatchBuilder represents a match builder
type MatchBuilder interface {
	Create() MatchBuilder
	WithInput(input Identifier) MatchBuilder
	WithPattern(pattern string) MatchBuilder
	Now() (Match, error)
}

// Match represents a match
type Match interface {
	Input() Identifier
	HasPattern() bool
	Pattern() string
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithDeclaration(declaration Declaration) VariableBuilder
	WithAssigment(assignment Assignment) VariableBuilder
	WithConcatenation(concatenation Concatenation) VariableBuilder
	WithDelete(delete VariableName) VariableBuilder
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
	Delete() VariableName
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
	WithVariable(variable VariableName) AssignmentBuilder
	WithValue(value Value) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents a variable assignment instruction
type Assignment interface {
	Variable() VariableName
	Value() Value
}

// ValueBuilder represents a value
type ValueBuilder interface {
	Create() ValueBuilder
	IsNil() ValueBuilder
	WithVariable(variable VariableName) ValueBuilder
	WithNumeric(numeric NumericValue) ValueBuilder
	WithBool(bl bool) ValueBuilder
	WithString(str string) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsNil() bool
	IsVariable() bool
	Variable() VariableName
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
	IsNil() TypeBuilder
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
	IsFrame() TypeBuilder
	Now() (Type, error)
}

// Type represents a variable type
type Type interface {
	IsNil() bool
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
	IsFrame() bool
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
	WithInput(input Identifier) TransformOperationBuilder
	WithResult(result VariableName) TransformOperationBuilder
	Now() (TransformOperation, error)
}

// TransformOperation represents a transform operation
type TransformOperation interface {
	Input() Identifier
	Result() VariableName
}

// StandardOperationBuilder represents a standard operation builder
type StandardOperationBuilder interface {
	Create() StandardOperationBuilder
	WithFirst(first Identifier) StandardOperationBuilder
	WithSecond(second Identifier) StandardOperationBuilder
	WithResult(result VariableName) StandardOperationBuilder
	Now() (StandardOperation, error)
}

// StandardOperation represents a standard operation
type StandardOperation interface {
	First() Identifier
	Second() Identifier
	Result() VariableName
}

// RemainingOperationBuilder represents a remaining operation builder
type RemainingOperationBuilder interface {
	Create() RemainingOperationBuilder
	WithFirst(first Identifier) RemainingOperationBuilder
	WithSecond(second Identifier) RemainingOperationBuilder
	WithResult(result VariableName) RemainingOperationBuilder
	WithRemaining(remaining VariableName) RemainingOperationBuilder
	Now() (RemainingOperation, error)
}

// RemainingOperation represents a an operation with a remaining value
type RemainingOperation interface {
	First() Identifier
	Second() Identifier
	Result() VariableName
	Remaining() VariableName
}

// PrintBuilder represents a print instruction builder
type PrintBuilder interface {
	Create() PrintBuilder
	WithValue(value Value) PrintBuilder
	Now() (Print, error)
}

// Print represents a print instruction
type Print interface {
	Value() Value
}

// JumpBuilder represents a jump builder
type JumpBuilder interface {
	Create() JumpBuilder
	WithLabel(label string) JumpBuilder
	WithCondition(condition Identifier) JumpBuilder
	Now() (Jump, error)
}

// Jump represents a jump instruction
type Jump interface {
	Label() string
	HasCondition() bool
	Condition() Identifier
}

// ExitBuilder represents an exit builder
type ExitBuilder interface {
	Create() ExitBuilder
	WithCondition(cond Identifier) ExitBuilder
	Now() (Exit, error)
}

// Exit represents an exit instruction
type Exit interface {
	HasCondition() bool
	Condition() Identifier
}

// CallBuilder represents a call builder
type CallBuilder interface {
	Create() CallBuilder
	WithName(name string) CallBuilder
	WithCondition(condition Identifier) CallBuilder
	Now() (Call, error)
}

// Call represents a call instruction
type Call interface {
	Name() string
	HasCondition() bool
	Condition() Identifier
}

// StackFrameBuilder represents a stackFrame builder
type StackFrameBuilder interface {
	Create() StackFrameBuilder
	WithPush(push Push) StackFrameBuilder
	WithPop(pop Pop) StackFrameBuilder
	Now() (StackFrame, error)
}

// StackFrame represents a stackFrame related instruction
type StackFrame interface {
	IsPush() bool
	Push() Push
	IsPop() bool
	Pop() Pop
}

// PushBuilder represents a push builder
type PushBuilder interface {
	Create() PushBuilder
	WithStackframe(stackframe VariableName) PushBuilder
	Now() (Push, error)
}

// Push represents the push instruction
type Push interface {
	HasStackFrame() bool
	StackFrame() VariableName
}

// PopBuilder represents a pop builder
type PopBuilder interface {
	Create() PopBuilder
	WithStackframe(stackframe TransformOperation) PopBuilder
	Now() (Pop, error)
}

// Pop represents the pop instruction
type Pop interface {
	HasStackFrame() bool
	StackFrame() TransformOperation
}

// IdentifierBuilder represents an identifier builder
type IdentifierBuilder interface {
	Create() IdentifierBuilder
	WithVariable(variable VariableName) IdentifierBuilder
	WithConstant(constant string) IdentifierBuilder
	Now() (Identifier, error)
}

// Identifier represents an identifier
type Identifier interface {
	IsVariable() bool
	Variable() VariableName
	IsConstant() bool
	Constant() string
	String() string
}

// VariableNameBuilder represents a variable name builder
type VariableNameBuilder interface {
	Create() VariableNameBuilder
	WithLocal(local string) VariableNameBuilder
	Now() (VariableName, error)
}

// VariableName represents a variable name
type VariableName interface {
	IsLocal() bool
	Local() string
	String() string
}
