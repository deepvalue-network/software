package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/match"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/token"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/transform"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/variablename"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	stackFrameBuilder     stackframe.Builder
	skipBuilder           stackframe.SkipBuilder
	transformBuilder      transform.Builder
	variableNameBuilder   variablename.Builder
	conditionBuilder      condition.Builder
	propositionBuilder    condition.PropositionBuilder
	remainingBuilder      remaining.Builder
	standardBuilder       standard.Builder
	matchBuilder          match.Builder
	valueBuilder          value.Builder
	varValueAdapter       var_value.Adapter
	varValueFactory       var_value.Factory
	varVariableBuilder    var_variable.Builder
	tokenCodeBuilder      token.CodeBuilder
	tokenCodeMatchBuilder token.CodeMatchBuilder
	tokenBuilder          token.Builder
	callBuilder           call.Builder
	exitBuilder           exit.Builder
	builder               Builder
}

func createAdapter(
	stackFrameBuilder stackframe.Builder,
	skipBuilder stackframe.SkipBuilder,
	transformBuilder transform.Builder,
	variableNameBuilder variablename.Builder,
	conditionBuilder condition.Builder,
	propositionBuilder condition.PropositionBuilder,
	remainingBuilder remaining.Builder,
	standardBuilder standard.Builder,
	matchBuilder match.Builder,
	valueBuilder value.Builder,
	varValueAdapter var_value.Adapter,
	varValueFactory var_value.Factory,
	varVariableBuilder var_variable.Builder,
	tokenCodeBuilder token.CodeBuilder,
	tokenCodeMatchBuilder token.CodeMatchBuilder,
	tokenBuilder token.Builder,
	callBuilder call.Builder,
	exitBuilder exit.Builder,
	builder Builder,
) Adapter {
	out := adapter{
		stackFrameBuilder:     stackFrameBuilder,
		skipBuilder:           skipBuilder,
		transformBuilder:      transformBuilder,
		variableNameBuilder:   variableNameBuilder,
		conditionBuilder:      conditionBuilder,
		propositionBuilder:    propositionBuilder,
		remainingBuilder:      remainingBuilder,
		standardBuilder:       standardBuilder,
		matchBuilder:          matchBuilder,
		valueBuilder:          valueBuilder,
		varValueAdapter:       varValueAdapter,
		varValueFactory:       varValueFactory,
		varVariableBuilder:    varVariableBuilder,
		tokenCodeBuilder:      tokenCodeBuilder,
		tokenCodeMatchBuilder: tokenCodeMatchBuilder,
		tokenBuilder:          tokenBuilder,
		callBuilder:           callBuilder,
		exitBuilder:           exitBuilder,
		builder:               builder,
	}

	return &out
}

// ToInstruction converts a parsed instruction to an optimized instruction
func (app *adapter) ToInstruction(parsed parsers.Instruction) (Instruction, error) {
	builder := app.builder.Create()

	if parsed.IsPrint() {
		parsedVal := parsed.Print().Value()
		varValue, err := app.varValueAdapter.ToValue(parsedVal)
		if err != nil {
			return nil, err
		}

		val, err := app.valueBuilder.Create().WithValue(varValue).IsPrint().Now()
		if err != nil {
			return nil, err
		}

		builder.WithValue(val)
	}

	if parsed.IsVariable() {
		vr := parsed.Variable()
		if vr.IsAssignment() {
			ass := vr.Assignment()
			parsedVal := ass.Value()
			val, err := app.varValueAdapter.ToValue(parsedVal)
			if err != nil {
				return nil, err
			}

			parsedVariableName := ass.Variable()
			ins, err := app.varVariableBuilder.Create().WithValue(val).WithName(parsedVariableName).Now()
			if err != nil {
				return nil, err
			}

			builder.WithSave(ins)
		}

		if vr.IsDeclaration() {
			decl := vr.Declaration()
			name := decl.Variable()
			typ := decl.Type()
			val, err := app.varValueFactory.Create(typ)
			if err != nil {
				return nil, err
			}

			ins, err := app.varVariableBuilder.Create().WithName(name).WithValue(val).Now()
			if err != nil {
				return nil, err
			}

			builder.WithInsert(ins)
		}

		if vr.IsConcatenation() {
			concat := vr.Concatenation().Operation()
			st, err := app.standard(concat).IsConcatenation().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStandard(st)
		}

		if vr.IsDelete() {
			del := vr.Delete()
			builder.WithDelete(del)
		}
	}

	if parsed.IsOperation() {
		op := parsed.Operation()
		if op.IsArythmetic() {
			ary := op.Arythmetic()
			if ary.IsAdd() {
				add := ary.Add()
				st, err := app.standard(add).IsAdd().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if ary.IsDiv() {
				div := ary.Div()
				rem, err := app.remaining(div).IsDiv().Now()
				if err != nil {
					return nil, err
				}

				builder.WithRemaining(rem)
			}

			if ary.IsMul() {
				mul := ary.Mul()
				st, err := app.standard(mul).IsMul().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if ary.IsSub() {
				sub := ary.Sub()
				st, err := app.standard(sub).IsSub().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}
		}

		if op.IsRelational() {
			rel := op.Relational()
			if rel.IsEqual() {
				eq := rel.Equal()
				st, err := app.standard(eq).IsEqual().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if rel.IsLessThan() {
				lessThan := rel.LessThan()
				st, err := app.standard(lessThan).IsLessThan().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if rel.IsNotEqual() {
				notEqual := rel.NotEqual()
				st, err := app.standard(notEqual).IsNotEqual().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}
		}

		if op.IsLogical() {
			logical := op.Logical()
			if logical.IsAnd() {
				and := logical.And()
				st, err := app.standard(and).IsAnd().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if logical.IsOr() {
				or := logical.Or()
				st, err := app.standard(or).IsOr().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}
		}
	}

	if parsed.IsStackFrame() {
		stackFrame := parsed.StackFrame()
		if stackFrame.IsPush() {
			stackframe, err := app.stackFrameBuilder.Create().IsPush().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}

		if stackFrame.IsPop() {
			stackframe, err := app.stackFrameBuilder.Create().IsPop().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}

		if stackFrame.IsSkip() {
			skipBuilder := app.skipBuilder.Create()
			ptr := stackFrame.Skip().Pointer()
			if ptr.IsInt() {
				intVal := ptr.Int()
				skipBuilder.WithInt(intVal)
			}

			if ptr.IsVariable() {
				variable := ptr.Variable()
				skipBuilder.WithVariable(variable)
			}

			skip, err := skipBuilder.Now()
			if err != nil {
				return nil, err
			}

			stackframe, err := app.stackFrameBuilder.Create().WithSkip(skip).Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}

		if stackFrame.IsIndex() {
			variable := stackFrame.Index().Variable()
			stackframe, err := app.stackFrameBuilder.Create().WithIndex(variable).Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}
	}

	if parsed.IsJump() {
		jmp := parsed.Jump()
		condition, err := app.conditionFromJump(jmp)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if parsed.IsMatch() {
		parsedMatch := parsed.Match()
		match, err := app.match(parsedMatch).Now()
		if err != nil {
			return nil, err
		}

		builder.WithMatch(match)

	}

	if parsed.IsToken() {
		parsedToken := parsed.Token()
		tok, err := app.token(parsedToken)
		if err != nil {
			return nil, err
		}

		builder.WithToken(tok)
	}

	if parsed.IsExit() {
		parsedExit := parsed.Exit()
		exit, err := app.exit(parsedExit)
		if err != nil {
			return nil, err
		}

		builder.WithExit(exit)
	}

	if parsed.IsCall() {
		parsedCall := parsed.Call()
		call, err := app.call(parsedCall)
		if err != nil {
			return nil, err
		}

		builder.WithCall(call)
	}

	return builder.Now()
}

func (app *adapter) call(parsed parsers.Call) (call.Call, error) {
	name := parsed.Name()
	builder := app.callBuilder.Create().WithName(name)
	if parsed.HasCondition() {
		condition := parsed.Condition()
		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *adapter) exit(parsed parsers.Exit) (exit.Exit, error) {
	builder := app.exitBuilder.Create()
	if parsed.HasCondition() {
		condition := parsed.Condition()
		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *adapter) token(parsed parsers.Token) (token.Token, error) {
	builder := app.tokenBuilder.Create()
	if parsed.IsCodeMatch() {
		parsedCodeMatch := parsed.CodeMatch()
		codeMatch, err := app.tokenCodeMatch(parsedCodeMatch)
		if err != nil {
			return nil, err
		}

		builder.WithCodeMatch(codeMatch)
	}

	if parsed.IsTokenSection() {
		parsedTokenSection := parsed.TokenSection()
		code, err := app.tokenSection(parsedTokenSection)
		if err != nil {
			return nil, err
		}

		builder.WithCode(code)
	}

	return builder.Now()
}

func (app *adapter) tokenCodeMatch(parsed parsers.CodeMatch) (token.CodeMatch, error) {
	ret := parsed.Content()

	section := parsed.Section()
	patternVariables := parsed.PatternVariables()

	return app.tokenCodeMatchBuilder.Create().WithReturn(ret).WithSectionName(section).WithPatterns(patternVariables).Now()
}

func (app *adapter) tokenSection(parsed parsers.TokenSection) (token.Code, error) {
	if parsed.IsVariableName() {
		ret := parsed.VariableName()
		return app.tokenCodeBuilder.Create().WithReturn(ret).Now()
	}

	specific := parsed.Specific()
	return app.tokenSpecificCode(specific)
}

func (app *adapter) tokenSpecificCode(parsed parsers.SpecificTokenCode) (token.Code, error) {
	ret := parsed.VariableName()
	patternVariable := parsed.PatternVariable()
	builder := app.tokenCodeBuilder.Create().WithReturn(ret).WithPattern(patternVariable)
	if parsed.HasAmount() {
		amount := parsed.Amount()
		builder.WithAmount(amount)
	}

	return builder.Now()
}

func (app *adapter) transform(parsed parsers.TransformOperation) transform.Builder {
	builder := app.transformBuilder.Create()
	input := parsed.Input()
	builder.WithInput(input)

	result := parsed.Result()
	return builder.WithResult(result)
}

func (app *adapter) conditionFromJump(parsed parsers.Jump) (condition.Condition, error) {
	label := parsed.Label()
	propositionBuilder := app.propositionBuilder.Create().WithName(label)
	if parsed.HasCondition() {
		condition := parsed.Condition()
		propositionBuilder.WithCondition(condition)
	}

	proposition, err := propositionBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().IsJump().WithProposition(proposition).Now()
}

func (app *adapter) match(parsed parsers.Match) match.Builder {
	input := parsed.Input()
	builder := app.matchBuilder.Create().WithInput(input)
	if parsed.HasPattern() {
		pattern := parsed.Pattern()
		builder.WithPattern(pattern)
	}

	return builder
}

func (app *adapter) remaining(parsed parsers.RemainingOperation) remaining.Builder {
	builder := app.remainingBuilder.Create()
	first := parsed.First()
	builder.WithFirst(first)

	second := parsed.Second()
	builder.WithSecond(second)

	result := parsed.Result()
	builder.WithResult(result)

	remaining := parsed.Remaining()
	builder.WithRemaining(remaining)

	return builder
}

func (app *adapter) standard(parsed parsers.StandardOperation) standard.Builder {
	builder := app.standardBuilder.Create()
	first := parsed.First()
	builder.WithFirst(first)

	second := parsed.Second()
	builder.WithSecond(second)

	result := parsed.Result()
	builder.WithResult(result)

	return builder
}
