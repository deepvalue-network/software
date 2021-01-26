package instruction

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/condition"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/match"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/remaining"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/stackframe"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/standard"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/token"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/transform"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/value"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/variablename"
	var_variable "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable"
	var_value "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

type adapter struct {
	stackFrameBuilder     stackframe.Builder
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
	builder               Builder
}

func createAdapter(
	stackFrameBuilder stackframe.Builder,
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
	builder Builder,
) Adapter {
	out := adapter{
		stackFrameBuilder:     stackFrameBuilder,
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
			variableBuilder := app.varVariableBuilder.Create().WithValue(val)
			if parsedVariableName.IsGlobal() {
				global := parsedVariableName.Global()
				variableBuilder.WithName(global).IsGlobal()
			}

			if parsedVariableName.IsLocal() {
				local := parsedVariableName.Local()
				variableBuilder.WithName(local)
			}

			ins, err := variableBuilder.Now()
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
			if del.IsGlobal() {
				global := del.Global()
				builder.WithDelete(global)
			}

			if del.IsLocal() {
				local := del.Local()
				builder.WithDelete(local)
			}
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
			push := stackFrame.Push()
			if push.HasStackFrame() {
				variableName := push.StackFrame()
				varName := app.variableName(variableName)
				vrName, err := app.variableNameBuilder.Create().WithVariableName(varName).IsPush().Now()
				if err != nil {
					return nil, err
				}

				builder.WithVariableName(vrName)
			}

			if !push.HasStackFrame() {
				stackframe, err := app.stackFrameBuilder.Create().IsPush().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStackframe(stackframe)
			}
		}

		if stackFrame.IsPop() {
			pop := stackFrame.Pop()
			if pop.HasStackFrame() {
				transformOperation := pop.StackFrame()
				trsf, err := app.transform(transformOperation).IsPop().Now()
				if err != nil {
					return nil, err
				}

				builder.WithTransform(trsf)
			}

			if !pop.HasStackFrame() {
				stackframe, err := app.stackFrameBuilder.Create().IsPop().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStackframe(stackframe)
			}
		}

		if stackFrame.IsAssignment() {
			standard := stackFrame.Assignment().Standard()
			st, err := app.standard(standard).IsFrameAssignment().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStandard(st)
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
	paersedContent := parsed.Content()
	ret := app.variableName(paersedContent)

	parsedSection := parsed.Section()
	section := app.variableName(parsedSection)

	tokenVariable := parsed.TokenVariable()
	patternVariables := parsed.PatternVariables()

	return app.tokenCodeMatchBuilder.Create().WithReturn(ret).WithSectionName(section).WithToken(tokenVariable).WithPatterns(patternVariables).Now()
}

func (app *adapter) tokenSection(parsed parsers.TokenSection) (token.Code, error) {
	if parsed.IsCode() {
		parsedCode := parsed.Code()
		return app.tokenCode(parsedCode)
	}

	specific := parsed.Specific()
	return app.tokenSpecificCode(specific)
}

func (app *adapter) tokenCode(parsed parsers.TokenCode) (token.Code, error) {
	content := parsed.Content()
	ret := app.variableName(content)

	tokenVariable := parsed.TokenVariable()
	return app.tokenCodeBuilder.Create().WithReturn(ret).WithToken(tokenVariable).Now()
}

func (app *adapter) tokenSpecificCode(parsed parsers.SpecificTokenCode) (token.Code, error) {
	content := parsed.Content()
	ret := app.variableName(content)

	tokenVariable := parsed.TokenVariable()
	patternVariable := parsed.PatternVariable()

	builder := app.tokenCodeBuilder.Create().WithReturn(ret).WithToken(tokenVariable).WithPattern(patternVariable)
	if parsed.HasAmount() {
		parsedAmount := parsed.Amount()
		amount := app.variableName(parsedAmount)
		builder.WithAmount(amount)
	}

	return builder.Now()
}

func (app *adapter) transform(parsed parsers.TransformOperation) transform.Builder {
	builder := app.transformBuilder.Create()
	inputIdentifier := parsed.Input()
	input := app.identifier(inputIdentifier)
	builder.WithInput(input)

	resultVariableName := parsed.Result()
	result := app.variableName(resultVariableName)
	return builder.WithResult(result)
}

func (app *adapter) conditionFromJump(parsed parsers.Jump) (condition.Condition, error) {
	label := parsed.Label()
	propositionBuilder := app.propositionBuilder.Create().WithName(label)
	if parsed.HasCondition() {
		conditionID := parsed.Condition()
		condition := app.identifier(conditionID)
		propositionBuilder.WithCondition(condition)
	}

	proposition, err := propositionBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().IsJump().WithProposition(proposition).Now()
}

func (app *adapter) match(parsed parsers.Match) match.Builder {
	input := parsed.Input().String()
	builder := app.matchBuilder.Create().WithInput(input)
	if parsed.HasPattern() {
		pattern := parsed.Pattern()
		builder.WithPattern(pattern)
	}

	return builder
}

func (app *adapter) remaining(parsed parsers.RemainingOperation) remaining.Builder {
	builder := app.remainingBuilder.Create()
	firstID := parsed.First()
	first := app.identifier(firstID)
	builder.WithFirst(first)

	secondID := parsed.Second()
	second := app.identifier(secondID)
	builder.WithSecond(second)

	resultVarName := parsed.Result()
	result := app.variableName(resultVarName)
	builder.WithResult(result)

	remainingVarName := parsed.Remaining()
	remaining := app.variableName(remainingVarName)
	builder.WithRemaining(remaining)

	return builder
}

func (app *adapter) standard(parsed parsers.StandardOperation) standard.Builder {
	builder := app.standardBuilder.Create()
	firstID := parsed.First()
	first := app.identifier(firstID)
	builder.WithFirst(first)

	secondID := parsed.Second()
	second := app.identifier(secondID)
	builder.WithSecond(second)

	resultVarName := parsed.Result()
	result := app.variableName(resultVarName)
	builder.WithResult(result)

	return builder
}

func (app *adapter) variableName(parsed parsers.VariableName) string {
	if parsed.IsGlobal() {
		global := parsed.Global()
		return global
	}

	local := parsed.Local()
	return local
}

func (app *adapter) identifier(parsed parsers.Identifier) string {
	if parsed.IsVariable() {
		vr := parsed.Variable()
		if vr.IsGlobal() {
			global := vr.Global()
			return global
		}

		local := vr.Local()
		return local
	}

	constant := parsed.Constant()
	return constant
}
