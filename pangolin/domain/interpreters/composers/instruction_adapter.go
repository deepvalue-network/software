package composers

import (
	"errors"
	"fmt"
	"math"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	command_labels "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/labels"
	command_mains "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/mains"
	command_tests "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/tests"
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/value"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions/instruction"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type instructionAdapter struct {
	testInstructionBuilder     parsers.TestInstructionBuilder
	assertBuilder              parsers.AssertBuilder
	readFileBuilder            parsers.ReadFileBuilder
	relativePathBuilder        parsers.RelativePathBuilder
	labelInstructionBuilder    parsers.LabelInstructionBuilder
	instructionBuilder         parsers.InstructionBuilder
	exitBuilder                parsers.ExitBuilder
	callBuilder                parsers.CallBuilder
	printBuilder               parsers.PrintBuilder
	operationBuilder           parsers.OperationBuilder
	arythmeticBuilder          parsers.ArythmeticBuilder
	relationalBuilder          parsers.RelationalBuilder
	logicalBuilder             parsers.LogicalBuilder
	variableBuilder            parsers.VariableBuilder
	declarationBuilder         parsers.DeclarationBuilder
	typeBuilder                parsers.TypeBuilder
	valueRepresentationBuilder parsers.ValueRepresentationBuilder
	valueBuilder               parsers.ValueBuilder
	numericValueBuilder        parsers.NumericValueBuilder
	assignmentBuilder          parsers.AssignmentBuilder
	concatenationBuilder       parsers.ConcatenationBuilder
	remainingOperationBuilder  parsers.RemainingOperationBuilder
	standardOperationBuilder   parsers.StandardOperationBuilder
	jumpBuilder                parsers.JumpBuilder
	stackFrameBuilder          parsers.StackFrameBuilder
	indexBuilder               parsers.IndexBuilder
	skipBuilder                parsers.SkipBuilder
	intPointBuilder            parsers.IntPointerBuilder
	localStackFrame            stackframes.StackFrame
	stkFrame                   stackframes.StackFrame
}

func createInstructionAdapter(
	testInstructionBuilder parsers.TestInstructionBuilder,
	assertBuilder parsers.AssertBuilder,
	readFileBuilder parsers.ReadFileBuilder,
	relativePathBuilder parsers.RelativePathBuilder,
	labelInstructionBuilder parsers.LabelInstructionBuilder,
	instructionBuilder parsers.InstructionBuilder,
	exitBuilder parsers.ExitBuilder,
	callBuilder parsers.CallBuilder,
	printBuilder parsers.PrintBuilder,
	operationBuilder parsers.OperationBuilder,
	arythmeticBuilder parsers.ArythmeticBuilder,
	relationalBuilder parsers.RelationalBuilder,
	logicalBuilder parsers.LogicalBuilder,
	variableBuilder parsers.VariableBuilder,
	declarationBuilder parsers.DeclarationBuilder,
	typeBuilder parsers.TypeBuilder,
	valueRepresentationBuilder parsers.ValueRepresentationBuilder,
	valueBuilder parsers.ValueBuilder,
	numericValueBuilder parsers.NumericValueBuilder,
	assignmentBuilder parsers.AssignmentBuilder,
	concatenationBuilder parsers.ConcatenationBuilder,
	remainingOperationBuilder parsers.RemainingOperationBuilder,
	standardOperationBuilder parsers.StandardOperationBuilder,
	jumpBuilder parsers.JumpBuilder,
	stackFrameBuilder parsers.StackFrameBuilder,
	indexBuilder parsers.IndexBuilder,
	skipBuilder parsers.SkipBuilder,
	intPointBuilder parsers.IntPointerBuilder,
	localStackFrame stackframes.StackFrame,
	stkFrame stackframes.StackFrame,
) InstructionAdapter {
	out := instructionAdapter{
		testInstructionBuilder:     testInstructionBuilder,
		assertBuilder:              assertBuilder,
		readFileBuilder:            readFileBuilder,
		relativePathBuilder:        relativePathBuilder,
		labelInstructionBuilder:    labelInstructionBuilder,
		instructionBuilder:         instructionBuilder,
		exitBuilder:                exitBuilder,
		callBuilder:                callBuilder,
		printBuilder:               printBuilder,
		operationBuilder:           operationBuilder,
		arythmeticBuilder:          arythmeticBuilder,
		relationalBuilder:          relationalBuilder,
		logicalBuilder:             logicalBuilder,
		variableBuilder:            variableBuilder,
		declarationBuilder:         declarationBuilder,
		typeBuilder:                typeBuilder,
		valueRepresentationBuilder: valueRepresentationBuilder,
		valueBuilder:               valueBuilder,
		numericValueBuilder:        numericValueBuilder,
		assignmentBuilder:          assignmentBuilder,
		concatenationBuilder:       concatenationBuilder,
		remainingOperationBuilder:  remainingOperationBuilder,
		standardOperationBuilder:   standardOperationBuilder,
		jumpBuilder:                jumpBuilder,
		stackFrameBuilder:          stackFrameBuilder,
		indexBuilder:               indexBuilder,
		skipBuilder:                skipBuilder,
		intPointBuilder:            intPointBuilder,
		localStackFrame:            localStackFrame,
		stkFrame:                   stkFrame,
	}

	return &out
}

// Test converts a command test instruction to a parsed test instruction
func (app *instructionAdapter) Test(cmdIns command_tests.Instruction) ([]parsers.TestInstruction, error) {
	ins := cmdIns.Instruction()
	scopes := cmdIns.Scopes()
	parsedIns, parsedVariables, err := app.testInstruction(ins, scopes)
	if err != nil {
		return nil, err
	}

	parsedVariablesIns, err := app.parsedVariablesToTestInstructions(parsedVariables)
	if err != nil {
		return nil, err
	}

	return append(parsedVariablesIns, parsedIns), nil
}

// Label converts a command label instruction to a parsed label instruction
func (app *instructionAdapter) Label(cmdIns command_labels.Instruction) ([]parsers.LabelInstruction, error) {
	ins := cmdIns.Instruction()
	scopes := cmdIns.Scopes()
	parsedIns, parsedVariables, err := app.labelInstruction(ins, scopes)
	if err != nil {
		return nil, err
	}

	parsedVariablesIns, err := app.parsedVariablesToLabelInstructions(parsedVariables)
	if err != nil {
		return nil, err
	}

	return append(parsedVariablesIns, parsedIns), nil
}

// Application converts a command main instruction to a parsed instruction
func (app *instructionAdapter) Application(cmdIns command_mains.Instruction) ([]parsers.Instruction, error) {
	ins := cmdIns.Instruction()
	scopes := cmdIns.Scopes()
	parsedIns, parsedVariables, err := app.mainInstruction(ins, scopes)
	if err != nil {
		return nil, err
	}

	parsedVariablesIns, err := app.parsedVariablesToInstructions(parsedVariables)
	if err != nil {
		return nil, err
	}

	return append(parsedVariablesIns, parsedIns), nil
}

func (app *instructionAdapter) testInstruction(testIns test_instruction.Instruction, scopes []bool) (parsers.TestInstruction, []parsers.Variable, error) {
	builder := app.testInstructionBuilder.Create()
	if testIns.IsAssert() {
		assert := testIns.Assert()
		index := assert.Index()

		assertBuilder := app.assertBuilder.Create().WithIndex(index)
		if assert.HasCondition() {
			condition := assert.Condition()
			assertBuilder.WithCondition(condition)
		}

		parsedAssert, err := assertBuilder.Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithAssert(parsedAssert)
	}

	if testIns.IsReadFile() {
		readFile := testIns.ReadFile()
		variable := readFile.Variable()
		path := readFile.Path()
		relativePath, err := app.relativePathBuilder.Create().WithPath(path).Now()
		if err != nil {
			return nil, nil, err
		}

		parsedReadFile, err := app.readFileBuilder.Create().WithVariable(variable).WithPath(relativePath).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithReadFile(parsedReadFile)
	}

	if testIns.IsInstruction() {
		ins := testIns.Instruction()
		parsedIns, variables, err := app.mainInstruction(ins, scopes)
		if err != nil {
			return nil, nil, err
		}

		parsedTestIns, err := builder.WithInstruction(parsedIns).Now()
		if err != nil {
			return nil, nil, err
		}

		return parsedTestIns, variables, nil
	}

	parsedTestIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return parsedTestIns, nil, nil
}

func (app *instructionAdapter) labelInstruction(lblIns label_instruction.Instruction, scopes []bool) (parsers.LabelInstruction, []parsers.Variable, error) {
	builder := app.labelInstructionBuilder.Create()
	if lblIns.IsRet() {
		builder.IsRet()
	}

	if lblIns.IsInstruction() {
		ins := lblIns.Instruction()
		parsedIns, variables, err := app.mainInstruction(ins, scopes)
		if err != nil {
			return nil, nil, err
		}

		parsedLblIns, err := builder.WithInstruction(parsedIns).Now()
		if err != nil {
			return nil, nil, err
		}

		return parsedLblIns, variables, nil

	}

	parsedLblIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return parsedLblIns, nil, nil
}

func (app *instructionAdapter) parsedVariablesToTestInstructions(parsedVariables []parsers.Variable) ([]parsers.TestInstruction, error) {
	insList, err := app.parsedVariablesToInstructions(parsedVariables)
	if err != nil {
		return nil, err
	}

	out := []parsers.TestInstruction{}
	for _, oneIns := range insList {
		lblIns, err := app.testInstructionBuilder.Create().WithInstruction(oneIns).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, lblIns)
	}

	return out, nil
}

func (app *instructionAdapter) parsedVariablesToLabelInstructions(parsedVariables []parsers.Variable) ([]parsers.LabelInstruction, error) {
	insList, err := app.parsedVariablesToInstructions(parsedVariables)
	if err != nil {
		return nil, err
	}

	out := []parsers.LabelInstruction{}
	for _, oneIns := range insList {
		lblIns, err := app.labelInstructionBuilder.Create().WithInstruction(oneIns).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, lblIns)
	}

	return out, nil
}

func (app *instructionAdapter) parsedVariablesToInstructions(parsedVariables []parsers.Variable) ([]parsers.Instruction, error) {
	out := []parsers.Instruction{}
	if parsedVariables != nil {
		for _, oneParsedVariable := range parsedVariables {
			vrIns, err := app.instructionBuilder.Create().WithVariable(oneParsedVariable).Now()
			if err != nil {
				return nil, err
			}

			out = append(out, vrIns)
		}
	}

	return out, nil
}

func (app *instructionAdapter) mainInstruction(appIns application_instruction.Instruction, scopes []bool) (parsers.Instruction, []parsers.Variable, error) {
	variables := []parsers.Variable{}
	builder := app.instructionBuilder.Create()
	if appIns.IsStackframe() {
		stackFrame := appIns.Stackframe()
		parsedStackFranme, err := app.stackFrame(stackFrame)
		if err != nil {
			return nil, nil, err
		}

		builder.WithStackFrame(parsedStackFranme)
	}

	if appIns.IsCondition() {
		condition := appIns.Condition()
		parsedJump, err := app.condition(condition)
		if err != nil {
			return nil, nil, err
		}

		builder.WithJump(parsedJump)
	}

	if appIns.IsStandard() {
		standard := appIns.Standard()
		parsedVariable, parsedOperation, vrList, err := app.standard(standard, scopes)
		if err != nil {
			return nil, nil, err
		}

		if parsedVariable != nil {
			builder.WithVariable(parsedVariable)
		}

		if parsedOperation != nil {
			builder.WithOperation(parsedOperation)
		}

		variables = append(variables, vrList...)
	}

	if appIns.IsRemaining() {
		remaining := appIns.Remaining()
		parsedOperation, vrList, err := app.remaining(remaining, scopes)
		if err != nil {
			return nil, nil, err
		}

		builder.WithOperation(parsedOperation)
		variables = append(variables, vrList...)
	}

	if appIns.IsValue() {
		value := appIns.Value()
		parsedPrint, err := app.value(value)
		if err != nil {
			return nil, nil, err
		}

		builder.WithPrint(parsedPrint)
	}

	if appIns.IsInsert() {
		insert := appIns.Insert()
		parsedVariable, err := app.insert(insert)
		if err != nil {
			return nil, nil, err
		}

		builder.WithVariable(parsedVariable)
	}

	if appIns.IsSave() {
		save := appIns.Save()
		parsedVariable, err := app.save(save)
		if err != nil {
			return nil, nil, err
		}

		builder.WithVariable(parsedVariable)
	}

	if appIns.IsDelete() {
		del := appIns.Delete()
		parsedVariable, err := app.delete(del)
		if err != nil {
			return nil, nil, err
		}

		builder.WithVariable(parsedVariable)
	}

	if appIns.IsCall() {
		call := appIns.Call()
		parsedCall, err := app.call(call)
		if err != nil {
			return nil, nil, err
		}

		builder.WithCall(parsedCall)
	}

	if appIns.IsExit() {
		exit := appIns.Exit()
		parsedExit, err := app.exit(exit)
		if err != nil {
			return nil, nil, err
		}

		builder.WithExit(parsedExit)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, variables, nil
}

func (app *instructionAdapter) stackFrame(stackFrame stackframe.Stackframe) (parsers.StackFrame, error) {
	builder := app.stackFrameBuilder.Create()
	if stackFrame.IsPush() {
		builder.IsPush()
	}

	if stackFrame.IsPop() {
		builder.IsPop()
	}

	if stackFrame.IsIndex() {
		indexVar := stackFrame.Index()
		parsedIndex, err := app.indexBuilder.Create().WithVariable(indexVar).Now()
		if err != nil {
			return nil, err
		}

		builder.WithIndex(parsedIndex)
	}

	if stackFrame.IsSkip() {
		pointerBuilder := app.intPointBuilder.Create()
		skip := stackFrame.Skip()
		if skip.IsInt() {
			intVal := skip.Int()
			pointerBuilder.WithInt(intVal)
		}

		if skip.IsVariable() {
			variable := skip.Variable()
			pointerBuilder.WithVariable(variable)
		}

		pointer, err := pointerBuilder.Now()
		if err != nil {
			return nil, err
		}

		parsedSkip, err := app.skipBuilder.Create().WithPointer(pointer).Now()
		if err != nil {
			return nil, err
		}

		builder.WithSkip(parsedSkip)
	}

	return builder.Now()
}

func (app *instructionAdapter) condition(condition condition.Condition) (parsers.Jump, error) {
	operation := condition.Operation()
	proposition := condition.Proposition()
	name := proposition.Name()
	if operation.IsJump() {
		jumpBuilder := app.jumpBuilder.Create().WithLabel(name)
		if proposition.HasCondition() {
			condition := proposition.Condition()
			jumpBuilder.WithCondition(condition)
		}

		return jumpBuilder.Now()
	}

	return nil, errors.New("the condition is invalid")
}

func (app *instructionAdapter) standard(standard standard.Standard, scopes []bool) (parsers.Variable, parsers.Operation, []parsers.Variable, error) {
	first := standard.First()
	second := standard.Second()
	variables, err := app.standardScopes(first, second, scopes)
	if err != nil {
		return nil, nil, nil, err
	}

	result := standard.Result()
	st, err := app.standardOperationBuilder.Create().WithFirst(first).WithSecond(second).WithResult(result).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	builder := app.operationBuilder.Create()
	operation := standard.Operation()
	if operation.IsArythmetic() {
		arythmeticBuilder := app.arythmeticBuilder.Create()
		arythmetic := operation.Arythmetic()
		if arythmetic.IsAdd() {
			arythmeticBuilder.WithAddition(st)
		}

		if arythmetic.IsSub() {
			arythmeticBuilder.WithSubstraction(st)
		}

		if arythmetic.IsMul() {
			arythmeticBuilder.WithMultiplication(st)
		}

		parsedArythmetic, err := arythmeticBuilder.Now()
		if err != nil {
			return nil, nil, nil, err
		}

		parsedOp, err := builder.WithArythmetic(parsedArythmetic).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, parsedOp, variables, nil
	}

	if operation.IsRelational() {
		relationalBuilder := app.relationalBuilder.Create()
		rel := operation.Relational()
		if rel.IsLessThan() {
			relationalBuilder.WithLessThan(st)
		}

		if rel.IsEqual() {
			relationalBuilder.WithEqual(st)
		}

		if rel.IsNotEqual() {
			relationalBuilder.WithNotEqual(st)
		}

		parsedRel, err := relationalBuilder.Now()
		if err != nil {
			return nil, nil, nil, err
		}

		parsedOp, err := builder.WithRelational(parsedRel).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, parsedOp, variables, nil
	}

	if operation.IsLogical() {
		logicalBuilder := app.logicalBuilder.Create()
		logical := operation.Logical()
		if logical.IsAnd() {
			logicalBuilder.WithAnd(st)
		}

		if logical.IsOr() {
			logicalBuilder.WithOr(st)
		}

		parsedLogical, err := logicalBuilder.Now()
		if err != nil {
			return nil, nil, nil, err
		}

		parsedOp, err := builder.WithLogical(parsedLogical).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, parsedOp, variables, nil
	}

	if !operation.IsMisc() {
		misc := operation.Misc()
		if misc.IsConcatenation() {
			concat, err := app.concatenationBuilder.Create().WithOperation(st).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			vr, err := app.variableBuilder.Create().WithConcatenation(concat).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			return vr, nil, variables, nil
		}

		return nil, nil, nil, errors.New("the standard misc is invalid")
	}

	return nil, nil, nil, errors.New("the standard is invalid")
}

func (app *instructionAdapter) remaining(rem remaining.Remaining, scopes []bool) (parsers.Operation, []parsers.Variable, error) {
	first := rem.First()
	second := rem.Second()
	variables, err := app.standardScopes(first, second, scopes)
	if err != nil {
		return nil, nil, err
	}

	remaining := rem.Remaining()
	result := rem.Result()
	parsedRemainingOp, err := app.remainingOperationBuilder.Create().WithFirst(first).WithSecond(second).WithResult(result).WithRemaining(remaining).Now()
	if err != nil {
		return nil, nil, err
	}

	builder := app.operationBuilder.Create()
	operation := rem.Operation()
	if operation.IsArythmetic() {
		arythmeticBuilder := app.arythmeticBuilder.Create()
		arythmetic := operation.Arythmetic()
		if arythmetic.IsDiv() {
			arythmeticBuilder.WithDivision(parsedRemainingOp)
		}

		parsedArythmetic, err := arythmeticBuilder.Now()
		if err != nil {
			return nil, nil, err
		}

		parsedOp, err := builder.WithArythmetic(parsedArythmetic).Now()
		if err != nil {
			return nil, nil, err
		}

		return parsedOp, variables, nil
	}

	return nil, nil, errors.New("the remaining is invalid")
}

func (app *instructionAdapter) value(value value.Value) (parsers.Print, error) {
	varVal := value.Value()
	operation := value.Operation()
	valueRepresentation, err := app.varValueToValueRepresentation(varVal)
	if err != nil {
		return nil, err
	}

	if operation.IsPrint() {
		return app.printBuilder.Create().WithValue(valueRepresentation).Now()
	}

	return nil, errors.New("the value is invalid")
}

func (app *instructionAdapter) insert(variable var_variable.Variable) (parsers.Variable, error) {
	value := variable.Value()
	typ, err := app.valueToType(value)
	if err != nil {
		return nil, err
	}

	name := variable.Name()
	decl, err := app.declarationBuilder.Create().WithVariable(name).WithType(typ).Now()
	if err != nil {
		return nil, err
	}

	return app.variableBuilder.Create().WithDeclaration(decl).Now()
}

func (app *instructionAdapter) valueToType(value var_value.Value) (parsers.Type, error) {
	if value.IsVariable() {
		return nil, errors.New("the variable's value was NOT expected to contain a variable")
	}

	if value.IsStackFrame() {
		return app.typeBuilder.Create().IsStackFrame().Now()
	}

	computable := value.Computable()
	return app.computableToType(computable)
}

func (app *instructionAdapter) save(variable var_variable.Variable) (parsers.Variable, error) {
	name := variable.Name()
	value := variable.Value()

	valueRepresentation, err := app.varValueToValueRepresentation(value)
	if err != nil {
		return nil, err
	}

	assign, err := app.assignmentBuilder.Create().WithVariable(name).WithValue(valueRepresentation).Now()
	if err != nil {
		return nil, err
	}

	return app.variableBuilder.Create().WithAssigment(assign).Now()
}

func (app *instructionAdapter) delete(variable string) (parsers.Variable, error) {
	return app.variableBuilder.Create().WithDelete(variable).Now()
}

func (app *instructionAdapter) call(call call.Call) (parsers.Call, error) {
	name := call.Name()
	builder := app.callBuilder.Create().WithName(name)
	if call.HasCondition() {
		condition := call.Condition()
		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *instructionAdapter) exit(exit exit.Exit) (parsers.Exit, error) {
	builder := app.exitBuilder.Create()
	if exit.HasCondition() {
		condition := exit.Condition()
		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *instructionAdapter) standardScopes(first string, second string, scopes []bool) ([]parsers.Variable, error) {
	if scopes == nil {
		return []parsers.Variable{}, nil
	}

	if len(scopes) != 2 {
		str := fmt.Sprintf("the scope in standard instruction expected %d scope elements, %d returned", 2, len(scopes))
		return nil, errors.New(str)
	}

	variables, err := app.scopes([]string{
		first, second,
	}, scopes)

	if err != nil {
		return nil, err
	}

	return variables, nil
}

func (app *instructionAdapter) scopes(globalNames []string, isGlobal []bool) ([]parsers.Variable, error) {
	variables := []parsers.Variable{}
	for _, oneGlobalName := range globalNames {
		decl, assign, err := app.declareThenAssign(oneGlobalName)
		if err != nil {
			return nil, err
		}

		variables = append(variables, decl)
		variables = append(variables, assign)
	}

	return variables, nil
}

func (app *instructionAdapter) declareThenAssign(globalName string) (parsers.Variable, parsers.Variable, error) {
	computedVal, err := app.stkFrame.Current().Fetch(globalName)
	if err != nil {
		return nil, nil, err
	}

	typ, err := app.computableToType(computedVal)
	if err != nil {
		return nil, nil, err
	}

	val, err := app.computableToValue(computedVal)
	if err != nil {
		return nil, nil, err
	}

	decl, err := app.declarationBuilder.Create().WithVariable(globalName).WithType(typ).Now()
	if err != nil {
		return nil, nil, err
	}

	declVariable, err := app.variableBuilder.Create().WithDeclaration(decl).Now()
	if err != nil {
		return nil, nil, err
	}

	valueRepresentation, err := app.valueRepresentationBuilder.Create().WithValue(val).Now()
	if err != nil {
		return nil, nil, err
	}

	assign, err := app.assignmentBuilder.Create().WithValue(valueRepresentation).Now()
	if err != nil {
		return nil, nil, err
	}

	assignVariable, err := app.variableBuilder.Create().WithAssigment(assign).Now()
	if err != nil {
		return nil, nil, err
	}

	return declVariable, assignVariable, nil
}

func (app *instructionAdapter) varValueToValueRepresentation(varVal var_value.Value) (parsers.ValueRepresentation, error) {
	builder := app.valueRepresentationBuilder.Create()
	if varVal.IsComputable() {
		computedVal := varVal.Computable()
		val, err := app.computableToValue(computedVal)
		if err != nil {
			return nil, err
		}

		builder.WithValue(val)
	}

	if varVal.IsVariable() {
		vr := varVal.Variable()
		builder.WithVariable(vr)
	}

	return builder.Now()
}

func (app *instructionAdapter) computableToType(computedVal computable.Value) (parsers.Type, error) {
	typeBuilder := app.typeBuilder.Create()
	if computedVal.IsBool() {
		typeBuilder.IsBool()
	}

	if computedVal.IsString() {
		typeBuilder.IsString()
	}

	if computedVal.IsIntHeight() {
		typeBuilder.IsInt8()
	}

	if computedVal.IsIntSixteen() {
		typeBuilder.IsInt16()
	}

	if computedVal.IsIntThirtyTwo() {
		typeBuilder.IsInt32()
	}

	if computedVal.IsIntSixtyFour() {
		typeBuilder.IsInt64()
	}

	if computedVal.IsUintHeight() {
		typeBuilder.IsUint8()
	}

	if computedVal.IsUintSixteen() {
		typeBuilder.IsUint16()
	}

	if computedVal.IsUintThirtyTwo() {
		typeBuilder.IsUint32()
	}

	if computedVal.IsUintSixtyFour() {
		typeBuilder.IsUint64()
	}

	if computedVal.IsFloatThirtyTwo() {
		typeBuilder.IsFloat32()
	}

	if computedVal.IsFloatSixtyFour() {
		typeBuilder.IsFloat64()
	}

	return typeBuilder.Now()
}

func (app *instructionAdapter) computableToValue(computedVal computable.Value) (parsers.Value, error) {
	isNumeric := true
	valueBuilder := app.valueBuilder.Create()
	if computedVal.IsBool() {
		bl := computedVal.Bool()
		valueBuilder.WithBool(*bl)

		isNumeric = false
	}

	if computedVal.IsString() {
		str := computedVal.String()
		valueBuilder.WithString(*str)

		isNumeric = false
	}

	numericValueBuilder := app.numericValueBuilder.Create()
	if computedVal.IsIntHeight() {
		pVal := computedVal.IntHeight()
		numericValueBuilder.WithInt(int(*pVal))
	}

	if computedVal.IsIntSixteen() {
		pVal := computedVal.IntSixteen()
		numericValueBuilder.WithInt(int(*pVal))
	}

	if computedVal.IsIntThirtyTwo() {
		pVal := computedVal.IntThirtyTwo()
		numericValueBuilder.WithInt(int(*pVal))
	}

	if computedVal.IsIntSixtyFour() {
		pVal := computedVal.IntSixtyFour()
		numericValueBuilder.WithInt(int(*pVal))
	}

	if computedVal.IsUintHeight() {
		pVal := computedVal.UintHeight()
		if *pVal < 0 {
			numericValueBuilder.IsNegative()
		}

		numericValueBuilder.WithInt(int(math.Abs(float64(*pVal))))
	}

	if computedVal.IsUintSixteen() {
		pVal := computedVal.UintSixteen()
		if *pVal < 0 {
			numericValueBuilder.IsNegative()
		}

		numericValueBuilder.WithInt(int(math.Abs(float64(*pVal))))
	}

	if computedVal.IsUintThirtyTwo() {
		pVal := computedVal.UintThirtyTwo()
		if *pVal < 0 {
			numericValueBuilder.IsNegative()
		}

		numericValueBuilder.WithInt(int(math.Abs(float64(*pVal))))
	}

	if computedVal.IsUintSixtyFour() {
		pVal := computedVal.UintSixtyFour()
		if *pVal < 0 {
			numericValueBuilder.IsNegative()
		}

		numericValueBuilder.WithInt(int(math.Abs(float64(*pVal))))
	}

	if computedVal.IsFloatThirtyTwo() {
		pVal := computedVal.FloatThirtyTwo()
		if *pVal < 0.0 {
			numericValueBuilder.IsNegative()
		}

		numericValueBuilder.WithFloat(math.Abs(float64(*pVal)))
	}

	if computedVal.IsFloatSixtyFour() {
		pVal := computedVal.FloatSixtyFour()
		if *pVal < 0.0 {
			numericValueBuilder.IsNegative()
		}

		numericValueBuilder.WithFloat(math.Abs(*pVal))
	}

	if computedVal.IsNil() {
		valueBuilder.IsNil()
		isNumeric = false
	}

	if isNumeric {
		numericValue, err := numericValueBuilder.Now()
		if err != nil {
			return nil, err
		}

		valueBuilder.WithNumeric(numericValue)
	}

	return valueBuilder.Now()
}
