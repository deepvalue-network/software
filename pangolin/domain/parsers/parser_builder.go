package parsers

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/lexers"
	lexer_parser "github.com/steve-care-software/products/pangolin/domain/lexers/parser"
)

type parserBuilder struct {
	lexerBuilder               lexers.Builder
	lexerParserApplication     lexer_parser.Application
	lexerParserBuilder         lexer_parser.Builder
	programBuilder             ProgramBuilder
	languageBuilder            LanguageBuilder
	languageValueBuilder       LanguageValueBuilder
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
	tokenCodeBuilder           TokenCodeBuilder
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
	lexerAdapter               lexers.Adapter
}

func createParserBuilder(
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerBuilder lexers.Builder,
	programBuilder ProgramBuilder,
	languageBuilder LanguageBuilder,
	languageValueBuilder LanguageValueBuilder,
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
	tokenCodeBuilder TokenCodeBuilder,
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
) ParserBuilder {
	out := parserBuilder{
		lexerParserApplication:     lexerParserApplication,
		lexerParserBuilder:         lexerParserBuilder,
		lexerBuilder:               lexerBuilder,
		programBuilder:             programBuilder,
		languageBuilder:            languageBuilder,
		languageValueBuilder:       languageValueBuilder,
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
		tokenCodeBuilder:           tokenCodeBuilder,
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
		lexerAdapter:               nil,
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
		app.languageValueBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.relativePathBuilder,
		app.folderSectionBuilder,
		app.folderNameBuilder,
		app.applicationBuilder,
		app.testSectionBuilder,
		app.testDeclarationBuilder,
		app.testInstructionBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.importSingleBuilder,
		app.constantSectionBuilder,
		app.constantDeclarationBuilder,
		app.variableSectionBuilder,
		app.variableDeclarationBuilder,
		app.variableDirectionBuilder,
		app.variableIncomingBuilder,
		app.definitionSectionBuilder,
		app.labelSectionBuilder,
		app.labelDeclarationBuilder,
		app.labelInstructionBuilder,
		app.mainSectionBuilder,
		app.instructionBuilder,
		app.tokenCodeBuilder,
		app.specificTokenCodeBuilder,
		app.tokenSectionBuilder,
		app.codeMatchBuilder,
		app.tokenBuilder,
		app.variableBuilder,
		app.concatenationBuilder,
		app.declarationBuilder,
		app.assignmentBuilder,
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
		app.pushBuilder,
		app.popBuilder,
		app.frameAssignmentBuilder,
		app.identifierBuilder,
		app.variableNameBuilder,
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
		app.languageValueBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.relativePathBuilder,
		app.folderSectionBuilder,
		app.folderNameBuilder,
		app.applicationBuilder,
		app.testSectionBuilder,
		app.testDeclarationBuilder,
		app.testInstructionBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.importSingleBuilder,
		app.constantSectionBuilder,
		app.constantDeclarationBuilder,
		app.variableSectionBuilder,
		app.variableDeclarationBuilder,
		app.variableDirectionBuilder,
		app.variableIncomingBuilder,
		app.definitionSectionBuilder,
		app.labelSectionBuilder,
		app.labelDeclarationBuilder,
		app.labelInstructionBuilder,
		app.mainSectionBuilder,
		app.instructionBuilder,
		app.tokenCodeBuilder,
		app.specificTokenCodeBuilder,
		app.tokenSectionBuilder,
		app.codeMatchBuilder,
		app.tokenBuilder,
		app.variableBuilder,
		app.concatenationBuilder,
		app.declarationBuilder,
		app.assignmentBuilder,
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
		app.pushBuilder,
		app.popBuilder,
		app.frameAssignmentBuilder,
		app.identifierBuilder,
		app.variableNameBuilder,
	)
}
