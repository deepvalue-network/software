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
	lexerAdapter               lexers.Adapter
	lexerApplication           lparser.Application
	lexerBuilder               lexers.Builder
	parserBuilder              lparser.Builder
	programBuilder             ProgramBuilder
	languageBuilder            LanguageBuilder
	languageValueBuilder       LanguageValueBuilder
	targetBuilder              TargetBuilder
	targetSingleBuilder        TargetSingleBuilder
	eventBuiilder              EventBuilder
	scriptBuilder              ScriptBuilder
	scriptValueBuilder         ScriptValueBuilder
	patternMatchBuilder        PatternMatchBuilder
	patternLabelsBuilder       PatternLabelsBuilder
	relativePathBuilder        RelativePathBuilder
	folderSectionBuilder       FolderSectionBuilder
	folderNameBuilder          FolderNameBuilder
	applicationBuilder         ApplicationBuilder
	testSectionBuilder         TestSectionBuilder
	testDeclarationBuilder     TestDeclarationBuilder
	testInstructionBuilder     TestInstructionBuilder
	assertBuilder              AssertBuilder
	readFileBuilder            ReadFileBuilder
	headSectionBuilder         HeadSectionBuilder
	headValueBuilder           HeadValueBuilder
	importSingleBuilder        ImportSingleBuilder
	constantSectionBuilder     ConstantSectionBuilder
	constantDeclarationBuilder ConstantDeclarationBuilder
	variableSectionBuilder     VariableSectionBuilder
	variableDeclarationBuilder VariableDeclarationBuilder
	variableDirectionBuilder   VariableDirectionBuilder
	variableIncomingBuilder    VariableIncomingBuilder
	definitionSectionBuilder   DefinitionSectionBuilder
	labelSectionBuilder        LabelSectionBuilder
	labelDeclarationBuilder    LabelDeclarationBuilder
	labelInstructionBuilder    LabelInstructionBuilder
	mainSectionBuilder         MainSectionBuilder
	instructionBuilder         InstructionBuilder
	triggerBuilder             TriggerBuilder
	formatBuilder              FormatBuilder
	specificTokenCodeBuilder   SpecificTokenCodeBuilder
	tokenSectionBuilder        TokenSectionBuilder
	codeMatchBuilder           CodeMatchBuilder
	tokenBuilder               TokenBuilder
	variableBuilder            VariableBuilder
	concatenationBuilder       ConcatenationBuilder
	declarationBuilder         DeclarationBuilder
	assignmentBuilder          AssignmentBuilder
	valueBuilder               ValueBuilder
	numericValueBuilder        NumericValueBuilder
	typeBuilder                TypeBuilder
	operationBuilder           OperationBuilder
	arythmeticBuilder          ArythmeticBuilder
	relationalBuilder          RelationalBuilder
	logicalBuilder             LogicalBuilder
	transformOperationBuilder  TransformOperationBuilder
	standardOperationBuilder   StandardOperationBuilder
	remainingOperationBuilder  RemainingOperationBuilder
	printBuilder               PrintBuilder
	jumpBuilder                JumpBuilder
	matchBuilder               MatchBuilder
	exitBuilder                ExitBuilder
	callBuilder                CallBuilder
	stackFrameBuilder          StackFrameBuilder
	pushBuilder                PushBuilder
	popBuilder                 PopBuilder
	frameAssignmentBuilder     FrameAssignmentBuiler
	identifierBuilder          IdentifierBuilder
	variableNameBuilder        VariableNameBuilder
	program                    map[string]Program
	language                   map[string]Language
	languageValue              map[string]LanguageValue
	target                     map[string]Target
	targetSingle               map[string]TargetSingle
	targetEvents               map[string][]Event
	targetPath                 map[string]RelativePath
	event                      map[string]Event
	script                     map[string]Script
	scriptValue                map[string]ScriptValue
	patternMatch               map[string]PatternMatch
	patternLabels              map[string]PatternLabels
	patternLabelEnter          map[string]string
	patternLabelExit           map[string]string
	relativePath               map[string]RelativePath
	folderSection              map[string]FolderSection
	folderName                 map[string]FolderName
	application                map[string]Application
	testSection                map[string]TestSection
	testDeclaration            map[string]TestDeclaration
	testInstruction            map[string]TestInstruction
	assert                     map[string]Assert
	readFile                   map[string]ReadFile
	headSection                map[string]HeadSection
	headValue                  map[string]HeadValue
	importSingle               map[string]ImportSingle
	constantSection            map[string]ConstantSection
	constantDeclaration        map[string]ConstantDeclaration
	variableSection            map[string]VariableSection
	variableDeclaration        map[string]VariableDeclaration
	variableDirection          map[string]VariableDirection
	variableIncoming           map[string]VariableIncoming
	definitionSection          map[string]DefinitionSection
	labelSection               map[string]LabelSection
	labelDeclaration           map[string]LabelDeclaration
	labelInstruction           map[string]LabelInstruction
	mainSection                map[string]MainSection
	instruction                map[string]Instruction
	trigger                    map[string]Trigger
	format                     map[string]Format
	specificTokenCode          map[string]SpecificTokenCode
	tokenSection               map[string]TokenSection
	codeMatch                  map[string]CodeMatch
	token                      map[string]Token
	callPattern                map[string]string
	callPatterns               map[string][]string
	patternOrRule              map[string]string
	importNames                map[string]string
	extendNames                map[string]string
	variable                   map[string]Variable
	concatenation              map[string]Concatenation
	declaration                map[string]Declaration
	assignment                 map[string]Assignment
	value                      map[string]Value
	numericValue               map[string]NumericValue
	boolValue                  map[string]bool
	floatValue                 map[string]float64
	stringValue                map[string]string
	typ                        map[string]Type
	operation                  map[string]Operation
	arythmetic                 map[string]Arythmetic
	relational                 map[string]Relational
	logical                    map[string]Logical
	transformOperation         map[string]TransformOperation
	standardOperation          map[string]StandardOperation
	remainingOperation         map[string]RemainingOperation
	print                      map[string]Print
	jump                       map[string]Jump
	match                      map[string]Match
	matchPattern               map[string]string
	exit                       map[string]Exit
	call                       map[string]Call
	stackFrame                 map[string]StackFrame
	push                       map[string]Push
	pop                        map[string]Pop
	frameAssignment            map[string]FrameAssignment
	identifier                 map[string]Identifier
	variableName               map[string]VariableName
}

func createParser(
	lexerAdapter lexers.Adapter,
	lexerApplication lparser.Application,
	parserBuilder lparser.Builder,
	lexerBuilder lexers.Builder,
	programBuilder ProgramBuilder,
	languageBuilder LanguageBuilder,
	languageValueBuilder LanguageValueBuilder,
	targetBuilder TargetBuilder,
	targetSingleBuilder TargetSingleBuilder,
	eventBuiilder EventBuilder,
	scriptBuilder ScriptBuilder,
	scriptValueBuilder ScriptValueBuilder,
	patternMatchBuilder PatternMatchBuilder,
	patternLabelsBuilder PatternLabelsBuilder,
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
	constantSectionBuilder ConstantSectionBuilder,
	constantDeclarationBuilder ConstantDeclarationBuilder,
	variableSectionBuilder VariableSectionBuilder,
	variableDeclarationBuilder VariableDeclarationBuilder,
	variableDirectionBuilder VariableDirectionBuilder,
	variableIncomingBuilder VariableIncomingBuilder,
	definitionSectionBuilder DefinitionSectionBuilder,
	labelSectionBuilder LabelSectionBuilder,
	labelDeclarationBuilder LabelDeclarationBuilder,
	labelInstructionBuilder LabelInstructionBuilder,
	mainSectionBuilder MainSectionBuilder,
	instructionBuilder InstructionBuilder,
	triggerBuilder TriggerBuilder,
	formatBuilder FormatBuilder,
	specificTokenCodeBuilder SpecificTokenCodeBuilder,
	tokenSectionBuilder TokenSectionBuilder,
	codeMatchBuilder CodeMatchBuilder,
	tokenBuilder TokenBuilder,
	variableBuilder VariableBuilder,
	concatenationBuilder ConcatenationBuilder,
	declarationBuilder DeclarationBuilder,
	assignmentBuilder AssignmentBuilder,
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
	pushBuilder PushBuilder,
	popBuilder PopBuilder,
	frameAssignmentBuilder FrameAssignmentBuiler,
	identifierBuilder IdentifierBuilder,
	variableNameBuilder VariableNameBuilder,
) (*parser, error) {
	out := &parser{
		lexerApplication:           lexerApplication,
		parserBuilder:              parserBuilder,
		lexerBuilder:               lexerBuilder,
		lexerAdapter:               lexerAdapter,
		programBuilder:             programBuilder,
		languageBuilder:            languageBuilder,
		languageValueBuilder:       languageValueBuilder,
		targetBuilder:              targetBuilder,
		targetSingleBuilder:        targetSingleBuilder,
		eventBuiilder:              eventBuiilder,
		scriptBuilder:              scriptBuilder,
		scriptValueBuilder:         scriptValueBuilder,
		patternMatchBuilder:        patternMatchBuilder,
		patternLabelsBuilder:       patternLabelsBuilder,
		relativePathBuilder:        relativePathBuilder,
		folderSectionBuilder:       folderSectionBuilder,
		folderNameBuilder:          folderNameBuilder,
		applicationBuilder:         applicationBuilder,
		testSectionBuilder:         testSectionBuilder,
		testDeclarationBuilder:     testDeclarationBuilder,
		testInstructionBuilder:     testInstructionBuilder,
		assertBuilder:              assertBuilder,
		readFileBuilder:            readFileBuilder,
		headSectionBuilder:         headSectionBuilder,
		headValueBuilder:           headValueBuilder,
		importSingleBuilder:        importSingleBuilder,
		constantSectionBuilder:     constantSectionBuilder,
		constantDeclarationBuilder: constantDeclarationBuilder,
		variableSectionBuilder:     variableSectionBuilder,
		variableDeclarationBuilder: variableDeclarationBuilder,
		variableDirectionBuilder:   variableDirectionBuilder,
		variableIncomingBuilder:    variableIncomingBuilder,
		definitionSectionBuilder:   definitionSectionBuilder,
		labelSectionBuilder:        labelSectionBuilder,
		labelDeclarationBuilder:    labelDeclarationBuilder,
		labelInstructionBuilder:    labelInstructionBuilder,
		mainSectionBuilder:         mainSectionBuilder,
		instructionBuilder:         instructionBuilder,
		triggerBuilder:             triggerBuilder,
		formatBuilder:              formatBuilder,
		specificTokenCodeBuilder:   specificTokenCodeBuilder,
		tokenSectionBuilder:        tokenSectionBuilder,
		codeMatchBuilder:           codeMatchBuilder,
		tokenBuilder:               tokenBuilder,
		variableBuilder:            variableBuilder,
		concatenationBuilder:       concatenationBuilder,
		declarationBuilder:         declarationBuilder,
		assignmentBuilder:          assignmentBuilder,
		valueBuilder:               valueBuilder,
		numericValueBuilder:        numericValueBuilder,
		typeBuilder:                typeBuilder,
		operationBuilder:           operationBuilder,
		arythmeticBuilder:          arythmeticBuilder,
		relationalBuilder:          relationalBuilder,
		logicalBuilder:             logicalBuilder,
		transformOperationBuilder:  transformOperationBuilder,
		standardOperationBuilder:   standardOperationBuilder,
		remainingOperationBuilder:  remainingOperationBuilder,
		printBuilder:               printBuilder,
		jumpBuilder:                jumpBuilder,
		matchBuilder:               matchBuilder,
		exitBuilder:                exitBuilder,
		callBuilder:                callBuilder,
		stackFrameBuilder:          stackFrameBuilder,
		pushBuilder:                pushBuilder,
		popBuilder:                 popBuilder,
		frameAssignmentBuilder:     frameAssignmentBuilder,
		identifierBuilder:          identifierBuilder,
		variableNameBuilder:        variableNameBuilder,
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
			Token:  "language",
			OnExit: app.exitLanguage,
		},
		lparser.ToEventsParams{
			Token:  "languageValue",
			OnExit: app.exitLanguageValue,
		},
		lparser.ToEventsParams{
			Token:  "target",
			OnExit: app.exitTarget,
		},
		lparser.ToEventsParams{
			Token:  "targetSingle",
			OnExit: app.exitTargetSingle,
		},
		lparser.ToEventsParams{
			Token:  "targetEvents",
			OnExit: app.exitTargetEvents,
		},
		lparser.ToEventsParams{
			Token:  "targetPath",
			OnExit: app.exitTargetPath,
		},
		lparser.ToEventsParams{
			Token:  "event",
			OnExit: app.exitEvent,
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
			Token:  "importSingle",
			OnExit: app.exitImportSingle,
		},
		lparser.ToEventsParams{
			Token:  "constantSection",
			OnExit: app.exitConstantSection,
		},
		lparser.ToEventsParams{
			Token:  "constantDeclaration",
			OnExit: app.exitConstantDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "variableSection",
			OnExit: app.exitVariableSection,
		},
		lparser.ToEventsParams{
			Token:  "variableDeclaration",
			OnExit: app.exitVariableDeclaration,
		},
		lparser.ToEventsParams{
			Token:  "variableDirection",
			OnExit: app.exitVariableDirection,
		},
		lparser.ToEventsParams{
			Token:  "variableIncoming",
			OnExit: app.exitVariableIncoming,
		},
		lparser.ToEventsParams{
			Token:  "definitionSection",
			OnExit: app.exitDefinitionSection,
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
			Token:  "trigger",
			OnExit: app.exitTrigger,
		},
		lparser.ToEventsParams{
			Token:  "format",
			OnExit: app.exitFormat,
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
			Token:  "transformOperation",
			OnExit: app.exitTransformOperation,
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
			Token:  "push",
			OnExit: app.exitPush,
		},
		lparser.ToEventsParams{
			Token:  "pop",
			OnExit: app.exitPop,
		},
		lparser.ToEventsParams{
			Token:  "frameAssignment",
			OnExit: app.exitFrameAssignment,
		},
		lparser.ToEventsParams{
			Token:  "identifier",
			OnExit: app.exitIdentifier,
		},
		lparser.ToEventsParams{
			Token:  "variableName",
			OnExit: app.exitVariableName,
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
	lexer, err := app.lexerAdapter.ToLexer(string(script))
	if err != nil {
		return nil, err
	}

	return app.Execute(lexer)
}

func (app *parser) init() {
	app.program = map[string]Program{}
	app.language = map[string]Language{}
	app.languageValue = map[string]LanguageValue{}
	app.target = map[string]Target{}
	app.targetSingle = map[string]TargetSingle{}
	app.targetEvents = map[string][]Event{}
	app.targetPath = map[string]RelativePath{}
	app.event = map[string]Event{}
	app.script = map[string]Script{}
	app.scriptValue = map[string]ScriptValue{}
	app.patternMatch = map[string]PatternMatch{}
	app.patternLabels = map[string]PatternLabels{}
	app.patternLabelEnter = map[string]string{}
	app.patternLabelExit = map[string]string{}
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
	app.importSingle = map[string]ImportSingle{}
	app.constantSection = map[string]ConstantSection{}
	app.constantDeclaration = map[string]ConstantDeclaration{}
	app.variableSection = map[string]VariableSection{}
	app.variableDeclaration = map[string]VariableDeclaration{}
	app.variableDirection = map[string]VariableDirection{}
	app.variableIncoming = map[string]VariableIncoming{}
	app.definitionSection = map[string]DefinitionSection{}
	app.labelSection = map[string]LabelSection{}
	app.labelDeclaration = map[string]LabelDeclaration{}
	app.labelInstruction = map[string]LabelInstruction{}
	app.mainSection = map[string]MainSection{}
	app.instruction = map[string]Instruction{}
	app.trigger = map[string]Trigger{}
	app.format = map[string]Format{}
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
	app.transformOperation = map[string]TransformOperation{}
	app.standardOperation = map[string]StandardOperation{}
	app.remainingOperation = map[string]RemainingOperation{}
	app.print = map[string]Print{}
	app.jump = map[string]Jump{}
	app.match = map[string]Match{}
	app.matchPattern = map[string]string{}
	app.exit = map[string]Exit{}
	app.call = map[string]Call{}
	app.stackFrame = map[string]StackFrame{}
	app.push = map[string]Push{}
	app.pop = map[string]Pop{}
	app.frameAssignment = map[string]FrameAssignment{}
	app.identifier = map[string]Identifier{}
	app.variableName = map[string]VariableName{}
}

func (app *parser) exitProgram(tree lexers.NodeTree) (interface{}, error) {
	builder := app.programBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"application",
		"language",
		"script",
	})

	switch section {
	case "application":
		if app, ok := app.application[code]; ok {
			builder.WithApplication(app)
		}
		break
	case "language":
		if lang, ok := app.language[code]; ok {
			builder.WithLanguage(lang)
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

	app.program[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguage(tree lexers.NodeTree) (interface{}, error) {
	values := []LanguageValue{}
	codes := tree.CodesFromName("languageValue")
	for _, oneCode := range codes {
		if val, ok := app.languageValue[oneCode]; ok {
			values = append(values, val)
		}
	}

	ins, err := app.languageBuilder.Create().WithValues(values).Now()
	if err != nil {
		return nil, err
	}

	app.language[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitLanguageValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.languageValueBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"LANG_ROOT",
		"singleExtend",
		"relativePath",
		"GLOBAL_VARIABLE_PATTERN",
		"LANG_PATTERN_MATCHES",
		"LANG_TARGETS",
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

	case "GLOBAL_VARIABLE_PATTERN":
		if tree.CodeFromName("IN") != "" {
			builder.WithInputVariable(code)
			break
		}

		if tree.CodeFromName("OUT") != "" {
			builder.WithOutputVariable(code)
			break
		}
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
	case "LANG_TARGETS":
		targets := []Target{}
		codes := tree.CodesFromName("target")
		for _, oneCode := range codes {
			if target, ok := app.target[oneCode]; ok {
				targets = append(targets, target)
			}
		}

		builder.WithTargets(targets)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.languageValue[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTarget(tree lexers.NodeTree) (interface{}, error) {
	builder := app.targetBuilder.Create()
	name := tree.CodeFromName("TARGET_NAME_PATTERN")
	if name != "" {
		builder.WithName(name)
	}

	targetSingles := []TargetSingle{}
	codes := tree.CodesFromName("targetSingle")
	for _, oneCode := range codes {
		if targetSingle, ok := app.targetSingle[oneCode]; ok {
			targetSingles = append(targetSingles, targetSingle)
		}
	}

	if len(targetSingles) > 0 {
		builder.WithSingles(targetSingles)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.target[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTargetSingle(tree lexers.NodeTree) (interface{}, error) {
	builder := app.targetSingleBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"targetEvents",
		"targetPath",
	})

	switch section {
	case "targetEvents":
		if targetEvents, ok := app.targetEvents[code]; ok {
			builder.WithEvents(targetEvents)
		}
		break
	case "targetPath":
		if path, ok := app.targetPath[code]; ok {
			builder.WithPath(path)
		}
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.targetSingle[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitTargetEvents(tree lexers.NodeTree) (interface{}, error) {
	events := []Event{}
	codes := tree.CodesFromName("event")
	for _, oneCode := range codes {
		if evt, ok := app.event[oneCode]; ok {
			events = append(events, evt)
		}
	}

	if len(events) > 0 {
		app.targetEvents[tree.Code()] = events
		return events, nil
	}

	return nil, errors.New("the events inside the targetEvents are invalid")
}

func (app *parser) exitTargetPath(tree lexers.NodeTree) (interface{}, error) {
	relPathCode := tree.CodeFromName("relativePath")
	if relPathCode != "" {
		if relPath, ok := app.relativePath[relPathCode]; ok {
			app.targetPath[tree.Code()] = relPath
			return relPath, nil
		}
	}

	return nil, errors.New("the relativePath inside the targetPath is invalid")
}

func (app *parser) exitEvent(tree lexers.NodeTree) (interface{}, error) {
	builder := app.eventBuiilder.Create()
	name := tree.CodeFromName("EVENT_NAME_PATTERN")
	if name != "" {
		builder.WithName(name)
	}

	label := tree.CodeFromName("LABEL_PATTERN")
	if label != "" {
		builder.WithLabel(label)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.event[tree.Code()] = ins
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
		}
	case "SCRIPT_LANGUAGE":
		pathCode := tree.CodeFromName("relativePath")
		if filePath, ok := app.relativePath[pathCode]; ok {
			builder.WithLanguagePath(filePath)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.scriptValue[tree.Code()] = ins
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

	variable := tree.CodeFromName("GLOBAL_VARIABLE_PATTERN")
	if variable != "" {
		builder.WithVariable(variable)
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

	definitionSectionCode := tree.CodeFromName("definitionSection")
	if definitionSectionCode != "" {
		if defSection, ok := app.definitionSection[definitionSectionCode]; ok {
			builder.WithDefinition(defSection)
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
	identifierCode := tree.CodeFromName("identifier")
	if identifierCode != "" {
		if iden, ok := app.identifier[identifierCode]; ok {
			builder.WithCondition(iden)
		}
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
	nameCode := tree.CodeFromName("variableName")
	if nameCode != "" {
		if name, ok := app.variableName[nameCode]; ok {
			builder.WithVariable(name)
		}
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

func (app *parser) exitConstantSection(tree lexers.NodeTree) (interface{}, error) {
	lst := []ConstantDeclaration{}
	codes := tree.CodesFromName("constantDeclaration")
	for _, oneCode := range codes {
		if ins, ok := app.constantDeclaration[oneCode]; ok {
			lst = append(lst, ins)
		}
	}

	ins, err := app.constantSectionBuilder.Create().WithDeclarations(lst).Now()
	if err != nil {
		return nil, err
	}

	app.constantSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitConstantDeclaration(tree lexers.NodeTree) (interface{}, error) {
	builder := app.constantDeclarationBuilder.Create()

	typeCode := tree.CodeFromName("type")
	if typ, ok := app.typ[typeCode]; ok {
		builder.WithType(typ)
	}

	valueCode := tree.CodeFromName("value")
	if val, ok := app.value[valueCode]; ok {
		builder.WithValue(val)
	}

	constant := tree.CodeFromName("CONSTANT_PATTERN")
	ins, err := builder.WithConstant(constant).Now()
	if err != nil {
		return nil, err
	}

	app.constantDeclaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitVariableSection(tree lexers.NodeTree) (interface{}, error) {
	lst := []VariableDeclaration{}
	codes := tree.CodesFromName("variableDeclaration")
	for _, oneCode := range codes {
		if ins, ok := app.variableDeclaration[oneCode]; ok {
			lst = append(lst, ins)
		}
	}

	ins, err := app.variableSectionBuilder.Create().WithDeclarations(lst).Now()
	if err != nil {
		return nil, err
	}

	app.variableSection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitVariableDeclaration(tree lexers.NodeTree) (interface{}, error) {
	builder := app.variableDeclarationBuilder.Create()
	typeCode := tree.CodeFromName("type")
	if typ, ok := app.typ[typeCode]; ok {
		builder.WithType(typ)
	}

	directionCode := tree.CodeFromName("variableDirection")
	if directionCode != "" {
		if dirIns, ok := app.variableDirection[directionCode]; ok {
			builder.WithDirection(dirIns)
		}
	}

	variable := tree.CodeFromName("GLOBAL_VARIABLE_PATTERN")
	ins, err := builder.WithVariable(variable).Now()
	if err != nil {
		return nil, err
	}

	app.variableDeclaration[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitVariableDirection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.variableDirectionBuilder.Create()
	incomingCode := tree.CodeFromName("variableIncoming")
	if incomingCode != "" {
		if incoming, ok := app.variableIncoming[incomingCode]; ok {
			builder.WithIncoming(incoming)
		}
	}

	outCode := tree.CodeFromName("OUT")
	if outCode != "" {
		builder.IsOutgoing()
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.variableDirection[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitVariableIncoming(tree lexers.NodeTree) (interface{}, error) {
	builder := app.variableIncomingBuilder.Create()
	questionMarkCode := tree.CodeFromName("QUESTION_MARK")
	if questionMarkCode == "" {
		builder.IsMandatory()
	}

	valueCode := tree.CodeFromName("value")
	if valueCode != "" {
		if val, ok := app.value[valueCode]; ok {
			builder.WithOptionalDefaultValue(val)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.variableIncoming[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitDefinitionSection(tree lexers.NodeTree) (interface{}, error) {
	builder := app.definitionSectionBuilder.Create()
	constantSectionCode := tree.CodeFromName("constantSection")
	if constantSectionCode != "" {
		if cons, ok := app.constantSection[constantSectionCode]; ok {
			builder.WithConstants(cons)
		}
	}

	variableSectionCode := tree.CodeFromName("variableSection")
	if variableSectionCode != "" {
		if vr, ok := app.variableSection[variableSectionCode]; ok {
			builder.WithVariables(vr)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.definitionSection[tree.Code()] = ins
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
		"match",
		"exit",
		"call",
		"token",
		"trigger",
		"format",
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
	case "match":
		if mtch, ok := app.match[code]; ok {
			builder.WithMatch(mtch)
		}
	case "exit":
		if exit, ok := app.exit[code]; ok {
			builder.WithExit(exit)
		}
	case "call":
		if call, ok := app.call[code]; ok {
			builder.WithCall(call)
		}
		break
	case "token":
		if token, ok := app.token[code]; ok {
			builder.WithToken(token)
		}
		break
	case "trigger":
		if trigger, ok := app.trigger[code]; ok {
			builder.WithTrigger(trigger)
		}
	case "format":
		if format, ok := app.format[code]; ok {
			builder.WithFormat(format)
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

func (app *parser) exitTrigger(tree lexers.NodeTree) (interface{}, error) {
	builder := app.triggerBuilder.Create()
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if vrName, ok := app.variableName[variableNameCode]; ok {
			builder.WithVariableName(vrName)
		}
	}

	eventName := tree.CodeFromName("EVENT_NAME_PATTERN")
	if eventName != "" {
		builder.WithEvent(eventName)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.trigger[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitFormat(tree lexers.NodeTree) (interface{}, error) {
	builder := app.formatBuilder.Create()
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if results, ok := app.variableName[variableNameCode]; ok {
			builder.WithResults(results)
		}
	}

	identifierCodes := tree.CodesFromName("identifier")
	if len(identifierCodes) != 3 {
		str := fmt.Sprintf("%d variableName was expected, %d returned", 3, len(identifierCodes))
		return nil, errors.New(str)
	}

	if pattern, ok := app.identifier[identifierCodes[0]]; ok {
		builder.WithPattern(pattern)
	}

	if first, ok := app.identifier[identifierCodes[1]]; ok {
		builder.WithFirst(first)
	}

	if second, ok := app.identifier[identifierCodes[2]]; ok {
		builder.WithSecond(second)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.format[tree.Code()] = ins
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
	builder := app.codeMatchBuilder.Create()
	variableNameCodes := tree.CodesFromName("variableName")
	if len(variableNameCodes) != 2 {
		str := fmt.Sprintf("%d variableName was expected, %d returned", 2, len(variableNameCodes))
		return nil, errors.New(str)
	}

	if content, ok := app.variableName[variableNameCodes[0]]; ok {
		builder.WithContent(content)
	}

	if section, ok := app.variableName[variableNameCodes[1]]; ok {
		builder.WithSection(section)
	}

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
		"variableName",
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
	case "variableName":
		if variableName, ok := app.variableName[code]; ok {
			builder.WithVariableName(variableName)
		}
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
	builder := app.specificTokenCodeBuilder.Create()
	variableNameCodes := tree.CodesFromName("variableName")
	if len(variableNameCodes) != 2 {
		str := fmt.Sprintf("%d variableName was expected, %d returned", 2, len(variableNameCodes))
		return nil, errors.New(str)
	}

	if content, ok := app.variableName[variableNameCodes[0]]; ok {
		builder.WithVariableName(content)
	}

	if amount, ok := app.variableName[variableNameCodes[1]]; ok {
		builder.WithAmount(amount)
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

func (app *parser) exitSpecificTokenCode(tree lexers.NodeTree) (interface{}, error) {
	builder := app.specificTokenCodeBuilder.Create()
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if variableName, ok := app.variableName[variableNameCode]; ok {
			builder.WithVariableName(variableName)
		}
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
		"variableName",
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
	case "variableName":
		if variableName, ok := app.variableName[code]; ok {
			builder.WithDelete(variableName)
		}
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
	variableNameCode := tree.CodeFromName("LOCAL_VARIABLE_PATTERN")
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
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if varName, ok := app.variableName[variableNameCode]; ok {
			builder.WithVariable(varName)
		}
	}

	valueCode := tree.CodeFromName("value")
	if valueCode != "" {
		if val, ok := app.value[valueCode]; ok {
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

func (app *parser) exitValue(tree lexers.NodeTree) (interface{}, error) {
	builder := app.valueBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"variableName",
		"numericValue",
		"boolValue",
		"stringValue",
		"structValue",
		"NIL",
	})

	switch section {
	case "variableName":
		if variableName, ok := app.variableName[code]; ok {
			builder.WithVariable(variableName)
		}
		break
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
		"NIL",
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
		"FRAME",
		"TOKEN",
	})

	switch section {
	case "NIL":
		builder.IsNil()
		break
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
	case "FRAME":
		builder.IsFrame()
		break
	case "TOKEN":
		builder.IsToken()
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

func (app *parser) exitTransformOperation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.transformOperationBuilder.Create()
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if varName, ok := app.variableName[variableNameCode]; ok {
			builder.WithResult(varName)
		}
	}

	identifierCode := tree.CodeFromName("identifier")
	if identifierCode != "" {
		if identifier, ok := app.identifier[identifierCode]; ok {
			builder.WithInput(identifier)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.transformOperation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitStandardOperation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.standardOperationBuilder.Create()
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if res, ok := app.variableName[variableNameCode]; ok {
			builder.WithResult(res)
		}
	}

	identifiers := tree.CodesFromName("identifier")
	if len(identifiers) != 2 {
		str := fmt.Sprintf("two (2) identifiers were expected, %d provided", len(identifiers))
		return nil, errors.New(str)
	}

	if first, ok := app.identifier[identifiers[0]]; ok {
		builder.WithFirst(first)
	}

	if second, ok := app.identifier[identifiers[1]]; ok {
		builder.WithSecond(second)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.standardOperation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitRemainingOperation(tree lexers.NodeTree) (interface{}, error) {
	builder := app.remainingOperationBuilder.Create()
	variableNames := tree.CodesFromName("variableName")
	if len(variableNames) != 2 {
		str := fmt.Sprintf("two (2) variableNames were expected, %d provided", len(variableNames))
		return nil, errors.New(str)
	}

	identifiers := tree.CodesFromName("identifier")
	if len(identifiers) != 2 {
		str := fmt.Sprintf("two (2) identifiers were expected, %d provided", len(identifiers))
		return nil, errors.New(str)
	}

	if first, ok := app.identifier[identifiers[0]]; ok {
		builder.WithFirst(first)
	}

	if second, ok := app.identifier[identifiers[1]]; ok {
		builder.WithSecond(second)
	}

	if res, ok := app.variableName[variableNames[0]]; ok {
		builder.WithResult(res)
	}

	if rem, ok := app.variableName[variableNames[1]]; ok {
		builder.WithRemaining(rem)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.remainingOperation[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitPrint(tree lexers.NodeTree) (interface{}, error) {
	builder := app.printBuilder.Create()
	valueCode := tree.CodeFromName("value")
	if valueCode != "" {
		if val, ok := app.value[valueCode]; ok {
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
	identifierCode := tree.CodeFromName("identifier")

	if identifierCode != "" {
		if identifier, ok := app.identifier[identifierCode]; ok {
			builder.WithCondition(identifier)
		}
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
	identifierCode := tree.CodeFromName("identifier")
	if identifier, ok := app.identifier[identifierCode]; ok {
		builder.WithInput(identifier)
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
	identifierCode := tree.CodeFromName("identifier")
	if identifierCode != "" {
		if identifier, ok := app.identifier[identifierCode]; ok {
			builder.WithCondition(identifier)
		}
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
	identifierCode := tree.CodeFromName("identifier")
	if identifierCode != "" {
		if identifier, ok := app.identifier[identifierCode]; ok {
			builder.WithCondition(identifier)
		}
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

func (app *parser) exitStackFrame(tree lexers.NodeTree) (interface{}, error) {
	builder := app.stackFrameBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"push",
		"pop",
		"frameAssignment",
	})

	switch section {
	case "push":
		if push, ok := app.push[code]; ok {
			builder.WithPush(push)
		}

		break
	case "pop":
		if pop, ok := app.pop[code]; ok {
			builder.WithPop(pop)
		}

		break
	case "frameAssignment":
		if ass, ok := app.frameAssignment[code]; ok {
			builder.WithAssignment(ass)
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

func (app *parser) exitPush(tree lexers.NodeTree) (interface{}, error) {
	builder := app.pushBuilder.Create()
	variableNameCode := tree.CodeFromName("variableName")
	if variableNameCode != "" {
		if name, ok := app.variableName[variableNameCode]; ok {
			builder.WithStackframe(name)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.push[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitPop(tree lexers.NodeTree) (interface{}, error) {
	builder := app.popBuilder.Create()
	transformOperationCode := tree.CodeFromName("transformOperation")
	if transformOperationCode != "" {
		if op, ok := app.transformOperation[transformOperationCode]; ok {
			builder.WithStackframe(op)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.pop[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitFrameAssignment(tree lexers.NodeTree) (interface{}, error) {
	builder := app.frameAssignmentBuilder.Create()
	standardOperationCode := tree.CodeFromName("standardOperation")
	if standardOperationCode != "" {
		if std, ok := app.standardOperation[standardOperationCode]; ok {
			builder.WithStandard(std)
		}
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.frameAssignment[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitIdentifier(tree lexers.NodeTree) (interface{}, error) {
	builder := app.identifierBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"variableName",
		"CONSTANT_PATTERN",
	})

	switch section {
	case "variableName":
		if name, ok := app.variableName[code]; ok {
			builder.WithVariable(name)
		}
		break
	case "CONSTANT_PATTERN":
		builder.WithConstant(code)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.identifier[tree.Code()] = ins
	return ins, nil
}

func (app *parser) exitVariableName(tree lexers.NodeTree) (interface{}, error) {
	builder := app.variableNameBuilder.Create()
	section, code := tree.BestMatchFromNames([]string{
		"GLOBAL_VARIABLE_PATTERN",
		"LOCAL_VARIABLE_PATTERN",
	})

	switch section {
	case "GLOBAL_VARIABLE_PATTERN":
		builder.WithGlobal(code)
		break
	case "LOCAL_VARIABLE_PATTERN":
		builder.WithLocal(code)
		break
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, err
	}

	app.variableName[tree.Code()] = ins
	return ins, nil
}
